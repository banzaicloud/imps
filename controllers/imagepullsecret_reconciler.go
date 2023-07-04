// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/banzaicloud/imps/api/v1alpha1"
	"github.com/banzaicloud/imps/pkg/pullsecrets"
	"github.com/cisco-open/operator-tools/pkg/reconciler"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var requeueObject = ctrl.Result{
	Requeue:      true,
	RequeueAfter: 5 * time.Second,
}

func (r *ImagePullSecretReconciler) setFailedStatus(ctx context.Context, imps *v1alpha1.ImagePullSecret, failureError error) {
	imps.Status.Status = v1alpha1.ReconciliationFailed
	imps.Status.Reason = failureError.Error()
	err := r.Status().Update(ctx, imps)
	if err != nil {
		r.Log.Error("cannot update status field", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
	}
}

func (r *ImagePullSecretReconciler) setReadyStatus(ctx context.Context, imps *v1alpha1.ImagePullSecret, targetNamespaces StringSet, pullSecretExpires *time.Time) {
	imps.Status.LastSuccessfulReconciliation = &metav1.Time{Time: time.Now()}
	imps.Status.Status = v1alpha1.ReconciliationReady
	imps.Status.Reason = ""
	imps.Status.ManagedNamespaces = targetNamespaces
	if pullSecretExpires != nil {
		imps.Status.ValiditySeconds = int32(time.Until(*pullSecretExpires) / time.Second)
	} else {
		imps.Status.ValiditySeconds = 0
	}

	err := r.Status().Update(ctx, imps)
	if err != nil {
		r.Log.Error("cannot update status field", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
	}
}

func (r *ImagePullSecretReconciler) reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logur.WithField(r.Log, "imagepullsecret", req.NamespacedName)
	result := ctrl.Result{}

	var imps v1alpha1.ImagePullSecret
	err := r.Get(ctx, req.NamespacedName, &imps)
	if err != nil {
		if !apierrs.IsNotFound(err) {
			return result, errors.WithStack(err)
		}

		// related resources should be deleted automatically because of OwnerRef
		return result, nil
	}

	initialRun := imps.Status.Status == ""

	config := pullsecrets.NewConfigFromSecrets(ctx, r, imps.Spec.Registry.CredentialsAsNamespacedNameList())

	resultingConfig, err := config.ResultingDockerConfig(ctx)
	// Note: this only happens if the json serializer fails
	if err != nil {
		r.setFailedStatus(ctx, &imps, err)
		r.Recorder.Event(&imps, "Warning", "SourceCredentialsError",
			fmt.Sprintf("Cannot marshal resulting docker configuration: %s", err.Error()))
	}

	imps.Status.SourceSecretStatus = resultingConfig.AsStatus()
	pullSecret := resultingConfig.AsSecret("", imps.Spec.Target.Secret.Name)

	result, targetNamespaces, err := r.reconcileImagePullSecret(ctx, &imps, pullSecret)
	if err != nil {
		r.setFailedStatus(ctx, &imps, err)
		r.Recorder.Event(&imps, "Warning", "ReconciliationFailed",
			fmt.Sprintf("Cannot reconcile configuration: %s", err.Error()))
	}

	if err := resultingConfig.AsError(); err != nil {
		r.setFailedStatus(ctx, &imps, err)
		r.Recorder.Event(&imps, "Warning", "SourceCredentialsError",
			fmt.Sprintf("Source cerdentials failed to process: %v", resultingConfig.FailedSecrets()))
	} else {
		if initialRun {
			r.Recorder.Event(&imps, "Normal", "Reconciled", "Successfully reconciled selected secrets")
		}
		r.setReadyStatus(ctx, &imps, targetNamespaces, resultingConfig.Expiration)
	}
	logger.Info("Reconciling ImagePullSecret finished")

	return result, err
}

func (r *ImagePullSecretReconciler) reconcileImagePullSecret(ctx context.Context, imps *v1alpha1.ImagePullSecret, renderedPullSecret *corev1.Secret) (ctrl.Result, StringSet, error) {
	targetNamespaces, err := r.namespacesRequiringSecret(ctx, imps)
	if err != nil {
		r.Log.Warn("cannot get the list of namespaces requiring this secret", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
		r.Recorder.Event(imps, "Warning", "SecretReconciliationError", fmt.Sprintf("Cannot list namespaces requiring the secret: %s", err.Error()))

		return requeueObject, nil, err
	}

	// Let's continue in case of errors the initial secret creation as in case of ECR the tokens will expire, thus
	// it's better to reconcile what we can decreasing the blast radius of such reconciliation errors
	wasError := false
	// Reconcile secrets in selected namespaces
	for _, namespaceName := range targetNamespaces {
		err = r.reconcileSecretInNamespace(imps, namespaceName, renderedPullSecret)

		if err != nil {
			r.Log.Warn("cannot reconcile secret in namespace, skipping", map[string]interface{}{
				"ns":    namespaceName,
				"error": err,
				"imps":  imps,
			})
			r.Recorder.Event(imps, "Warning", "SecretReconciliationError", fmt.Sprintf("Cannot reconcile secret: %s/%s", namespaceName, imps.Spec.Target.Secret.Name))
			wasError = true

			continue
		}
		r.Log.Info("reconciled secret", map[string]interface{}{
			"namespace": namespaceName,
			"name":      imps.Spec.Target.Secret.Name,
		})
	}

	if wasError {
		return requeueObject, nil, errors.New("some secrets failed to reconcile")
	}

	var ownedSecrets corev1.SecretList
	err = r.Client.List(ctx, &ownedSecrets, client.MatchingLabels{
		labelImpsOwnerUID: string(imps.UID),
	})
	if err != nil {
		r.Log.Warn("cannot enumerate secrets owned by imps", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
		r.Recorder.Event(imps, "Warning", "SecretReconciliationError", fmt.Sprintf("Cannot enumerate secrets: %s", err.Error()))

		return requeueObject, nil, err
	}

	// Purge secrets that should not be there based on the selectors
	for _, existingSecret := range ownedSecrets.Items {
		shouldDelete := false
		if existingSecret.Name != imps.Spec.Target.Secret.Name {
			r.Log.Info("secret name does not match the expected one, removing", map[string]interface{}{
				"secret_name":      existingSecret.Name,
				"secret_namespace": existingSecret.Namespace,
				"imps":             imps.Name,
			})
			shouldDelete = true
		}

		if !targetNamespaces.Has(existingSecret.Namespace) {
			r.Log.Info("found secret in unselected namespace, removing", map[string]interface{}{
				"secret_name":      existingSecret.Name,
				"secret_namespace": existingSecret.Namespace,
				"imps":             imps.Name,
			})
			shouldDelete = true
		}

		if shouldDelete {
			_, err := r.ResourceReconciler.ReconcileResource(existingSecret.DeepCopy(), reconciler.StateAbsent)
			if err != nil {
				r.Log.Error("cannot delete secret", map[string]interface{}{
					"secret": existingSecret,
				})
				r.Recorder.Event(imps, "Warning", "SecretDeletionError", fmt.Sprintf("Cannot remove secret %s/%s", existingSecret.Namespace, existingSecret.Name))

				return requeueObject, nil, err
			}
		}
	}

	return ctrl.Result{}, targetNamespaces, nil
}

func (r *ImagePullSecretReconciler) namespacesRequiringSecret(ctx context.Context, imps *v1alpha1.ImagePullSecret) (StringSet, error) {
	var allNamespaces corev1.NamespaceList
	namespacesRequiringSecret := StringSet{}

	err := r.List(ctx, &allNamespaces)
	if err != nil {
		return nil, errors.Wrap(err, "cannot list namespaces")
	}

	matchingNamespaces, nonMatchingNamespaces, err := imps.SplitNamespacesByMatch(allNamespaces)
	if err != nil {
		return nil, err
	}

	for _, ns := range matchingNamespaces {
		namespacesRequiringSecret = append(namespacesRequiringSecret, ns.Name)
	}

	for _, ns := range nonMatchingNamespaces {
		shouldReconcile, err := r.anyPodMatchesSelectorInNS(ctx, imps, ns.DeepCopy())
		if err != nil {
			r.Log.Warn("cannot check for matching pods in namespace, skipping", map[string]interface{}{
				"ns":    ns,
				"error": err,
				"imps":  imps,
			})

			continue
		}
		if shouldReconcile {
			namespacesRequiringSecret = append(namespacesRequiringSecret, ns.Name)
		}
	}

	return namespacesRequiringSecret, nil
}

func (r *ImagePullSecretReconciler) anyPodMatchesSelectorInNS(ctx context.Context, imps *v1alpha1.ImagePullSecret, ns *corev1.Namespace) (bool, error) {
	// Let's prevent pod queries if there are no pod selector rules
	if imps.Spec.Target.NamespacesWithPods.IsEmpty() {
		return false, nil
	}

	var podsInNamespace corev1.PodList
	err := r.List(ctx, &podsInNamespace, client.InNamespace(ns.Name))
	if err != nil {
		return false, err
	}

	for _, pod := range podsInNamespace.Items {
		matches, err := imps.MatchesPod(pod.DeepCopy())
		if err != nil {
			r.Log.Warn("cannot match pod against an imps", map[string]interface{}{
				"error": err,
				"pod":   pod,
				"imps":  imps,
			})

			continue
		}

		if matches {
			return true, nil
		}
	}

	return false, nil
}

func (r *ImagePullSecretReconciler) reconcileSecretInNamespace(imps *v1alpha1.ImagePullSecret, targetNamespace string, referencedSecret *corev1.Secret) error {
	finalLabels := v1alpha1.LabelSet(imps.Spec.Target.Secret.Labels).DeepCopy()
	finalLabels[labelImpsOwnerUID] = string(imps.UID)

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:            imps.Spec.Target.Secret.Name,
			Namespace:       targetNamespace,
			Labels:          finalLabels,
			Annotations:     imps.Spec.Target.Secret.Annotations,
			OwnerReferences: []metav1.OwnerReference{imps.GetOwnerReferenceForOwnedObject()},
		},
		Data:       referencedSecret.Data,
		StringData: referencedSecret.StringData,
		Type:       referencedSecret.Type,
	}

	_, err := r.ResourceReconciler.ReconcileResource(secret, reconciler.StatePresent)
	if err != nil {
		return err
	}

	return nil
}
