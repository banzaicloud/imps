package controllers

import (
	"context"

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

	logger.Info("Reconciling ImagePullSecret finished")
	return result, nil
}

func (r *ImagePullSecretReconciler) reconcileImagePullSecret(ctx context.Context, imps *v1alpha1.ImagePullSecret) (ctrl.Result, error) {
	var allNamespaces corev1.NamespaceList
	err := r.List(ctx, &allNamespaces)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "cannot list namespaces")
	}

	matchingNamespaces, nonMatchingNamespaces, err := imps.SplitNamespacesByMatch(allNamespaces)
	if err != nil {
		return ctrl.Result{}, err
	}

	for _, ns := range matchingNamespaces {
		err = r.reconcileSecretInNamespace(ctx, imps, ns.Name)
		if err != nil {
			r.Log.Warn("cannot reconcile service in namespace, skipping", map[string]interface{}{
				"ns":    ns,
				"error": err,
				"imps":  imps,
			})
			continue
		}
		r.Log.Info("reconciled service", map[string]interface{}{
			"namespace": ns.Name,
			"name":      imps.Spec.TargetSecret,
		})
	}

	for _, ns := range nonMatchingNamespaces {
		shouldReconcile, err := r.anyPodMatchesSelectorInNS(ctx, imps, &ns)
		if err != nil {
			r.Log.Warn("cannot check for matching pods in namespace, skipping", map[string]interface{}{
				"ns":    ns,
				"error": err,
				"imps":  imps,
			})
			continue
		}
		if shouldReconcile {
			err = r.reconcileSecretInNamespace(ctx, imps, ns.Name)
			if err != nil {
				r.Log.Warn("cannot reconcile service in namespace, skipping", map[string]interface{}{
					"ns":    ns,
					"error": err,
					"imps":  imps,
				})
				continue
			}
			r.Log.Info("reconciled service", map[string]interface{}{
				"namespace": ns.Name,
				"name":      imps.Spec.TargetSecret,
			})
		}
	}

	// TODO: in case of reconciliation errors -> requeue!
	return ctrl.Result{}, nil
}

func (r *ImagePullSecretReconciler) anyPodMatchesSelectorInNS(ctx context.Context, imps *v1alpha1.ImagePullSecret, ns *corev1.Namespace) (bool, error) {
	// Let's prevent pod queries if there are no pod selector rules
	if len(imps.Spec.Pods) == 0 {
		return false, nil
	}

	var podsInNamespace corev1.PodList
	err := r.List(ctx, &podsInNamespace, client.InNamespace(ns.Namespace))
	if err != nil {
		return false, err
	}

	for _, pod := range podsInNamespace.Items {
		matches, err := imps.MatchesPod(&pod)
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

// TODO: remove unmanaged ones

func (r *ImagePullSecretReconciler) reconcileSecretInNamespace(ctx context.Context, imps *v1alpha1.ImagePullSecret, targetNamespace string) error {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        imps.Spec.TargetSecret.Name,
			Namespace:   targetNamespace,
			Labels:      imps.Spec.TargetSecret.Labels,
			Annotations: imps.Spec.TargetSecret.Annotations,
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
