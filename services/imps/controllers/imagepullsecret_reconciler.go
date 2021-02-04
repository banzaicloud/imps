// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package controllers

import (
	"context"
	"time"

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

	result, err = r.reconcileImagePullSecret(ctx, &imps)

	// TODO: add status handling here
	logger.Info("Reconciling ImagePullSecret finished")
	return result, err
}

func (r *ImagePullSecretReconciler) reconcileImagePullSecret(ctx context.Context, imps *v1alpha1.ImagePullSecret) (ctrl.Result, error) {
	targetNamespaces, err := r.namespacesRequiringSecret(ctx, imps)
	if err != nil {
		r.Log.Warn("cannot get the list of namespaces requiring this secret", map[string]interface{}{
			"error": err,
			"imps":  imps,
		})
		return requeueObject, err
	}

	// Let's continue in case of errors the initial secret creation as in case of ECR the tokens will expire, thus
	// it's better to reconcile what we can decreasing the blast radius of such reconciliation errors
	wasError := false
	// Reconcile secrets in selected namespaces
	for _, namespaceName := range targetNamespaces {
		err = r.reconcileSecretInNamespace(imps, namespaceName)

		if err != nil {
			r.Log.Warn("cannot reconcile secret in namespace, skipping", map[string]interface{}{
				"ns":    namespaceName,
				"error": err,
				"imps":  imps,
			})
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
		return requeueObject, err
	}

	// Purge secrets that should not be there based on the selectors
	for _, existingSecret := range ownedSecrets.Items {
		shouldDelete := false
		if existingSecret.Name != imps.Spec.Target.Secret.Name {
			r.Log.Info("secret name does not match the expected one, removing", map[string]interface{}{
				"secret_name":       existingSecret.Name,
				"secret_name_space": existingSecret.Namespace,
				"imps":              imps.Name,
			})
			shouldDelete = true
		}

		if !targetNamespaces.Has(existingSecret.Namespace) {
			r.Log.Info("found secret in unselected namespace, removing", map[string]interface{}{
				"secret_name":       existingSecret.Name,
				"secret_name_space": existingSecret.Namespace,
				"imps":              imps.Name,
			})
		}

		if shouldDelete {
			_, err := r.ResourceReconciler.ReconcileResource(existingSecret.DeepCopy(), reconciler.StateAbsent)
			if err != nil {
				r.Log.Error("cannot delete secret", map[string]interface{}{
					"secret": existingSecret,
				})
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

func (r *ImagePullSecretReconciler) reconcileSecretInNamespace(imps *v1alpha1.ImagePullSecret, targetNamespace string) error {
	finalLabels := imps.Spec.Target.Secret.Labels.DeepCopy()
	finalLabels[labelImpsOwnerUID] = string(imps.UID)

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:            imps.Spec.Target.Secret.Name,
			Namespace:       targetNamespace,
			Labels:          finalLabels,
			Annotations:     imps.Spec.Target.Secret.Annotations,
			OwnerReferences: []metav1.OwnerReference{imps.GetOwnerReferenceForOwnedObject()},
		},
		Data: map[string][]byte{
			"test": []byte("test"),
		},
	}

	_, err := r.ResourceReconciler.ReconcileResource(secret, reconciler.StatePresent)
	if err != nil {
		return err
	}

	return nil
}
