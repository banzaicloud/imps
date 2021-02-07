// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package controllers

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/types"

	"emperror.dev/errors"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/banzaicloud/backyards/services/imps/api/v1alpha1"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"

	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	requeueObject = ctrl.Result{
		Requeue:      true,
		RequeueAfter: 5 * time.Second,
	}
)

func (r *ImagePullSecretReconciler) setStatus(ctx context.Context, imps *v1alpha1.ImagePullSecret, status v1alpha1.ReconciliationStatus) {
	imps.Status = v1alpha1.ImagePullSecretStatus{Status: status}
	err := r.Status().Update(ctx, imps)
	if err != nil {
		r.Log.Error("cannot update status field", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
	}
}

func (r *ImagePullSecretReconciler) reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
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

	var referencedSecret corev1.Secret
	err = r.Get(ctx, types.NamespacedName{
		Namespace: imps.Spec.Registry.Credentials.Namespace,
		Name:      imps.Spec.Registry.Credentials.Name,
	}, &referencedSecret)

	if err != nil {
		r.setStatus(ctx, &imps, v1alpha1.ReconciliationFailed)
		r.Recorder.Event(&imps, "Warning", "SourceCredentialsAccessError", fmt.Sprintf("Cannot get registry credentials secret: %s/%s", imps.Spec.Registry.Credentials.Namespace, imps.Spec.Registry.Credentials.Name))
		return result, errors.WrapWithDetails(err, "cannot get referenced secret", "imps_name", imps.Name)
	}

	result, err = r.reconcileImagePullSecret(ctx, &imps, &referencedSecret)
	if err != nil {
		r.setStatus(ctx, &imps, v1alpha1.ReconciliationFailed)
	} else {
		if initialRun {
			r.Recorder.Event(&imps, "Normal", "Reconciled", "Successfully reconciled selected secrets")
		}
		r.setStatus(ctx, &imps, v1alpha1.ReconciliationReady)
	}
	logger.Info("Reconciling ImagePullSecret finished")
	return result, err
}

func (r *ImagePullSecretReconciler) reconcileImagePullSecret(ctx context.Context, imps *v1alpha1.ImagePullSecret, referencedSecret *corev1.Secret) (ctrl.Result, error) {
	targetNamespaces, err := r.namespacesRequiringSecret(ctx, imps)
	if err != nil {
		r.Log.Warn("cannot get the list of namespaces requiring this secret", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
		r.Recorder.Event(imps, "Warning", "SecretReconciliationError", fmt.Sprintf("Cannot list namespaces requiring the secret: %s", err.Error()))
		return requeueObject, err
	}

	// Let's continue in case of errors the initial secret creation as in case of ECR the tokens will expire, thus
	// it's better to reconcile what we can decreasing the blast radius of such reconciliation errors
	wasError := false
	// Reconcile secrets in selected namespaces
	for _, namespaceName := range targetNamespaces {
		err = r.reconcileSecretInNamespace(imps, namespaceName, referencedSecret)

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
		return requeueObject, errors.New("some secrets failed to reconcile")
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
		return requeueObject, err
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
				return requeueObject, err
			}
		}
	}

	return ctrl.Result{}, nil
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
	if len(imps.Spec.Target.NamespacesWithPods) == 0 {
		return false, nil
	}

	var podsInNamespace corev1.PodList
	err := r.List(ctx, &podsInNamespace, client.InNamespace(ns.Namespace))
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
		Data: referencedSecret.Data,
	}

	_, err := r.ResourceReconciler.ReconcileResource(secret, reconciler.StatePresent)
	if err != nil {
		return err
	}

	return nil
}
