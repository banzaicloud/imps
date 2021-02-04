// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/banzaicloud/operator-tools/pkg/reconciler"

	"emperror.dev/emperror"

	"k8s.io/apimachinery/pkg/runtime"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlBuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/banzaicloud/backyards/services/imps/api/v1alpha1"
)

// AlertingPolicyReconciler reconciles a AlertingPolicy object
type ImagePullSecretReconciler struct {
	client.Client
	Log          logur.Logger
	ErrorHandler emperror.ErrorHandler
	Scheme       *runtime.Scheme

	ResourceReconciler reconciler.ResourceReconciler
}

// +kubebuilder:rbac:groups=images.banzaicloud.io,resources=imagepullsecrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=images.banzaicloud.io,resources=imagepullsecrets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

func (r *ImagePullSecretReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	result, err := r.reconcile(req)
	if err != nil {
		r.ErrorHandler.Handle(err)
	}
	return result, err
}

func (r *ImagePullSecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	builder := ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ImagePullSecret{}, ctrlBuilder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Watches(
			&source.Kind{Type: &corev1.Namespace{}},
			&handler.EnqueueRequestsFromMapFunc{
				ToRequests: handler.ToRequestsFunc(r.impsMatchingNamespace),
			}).
		Watches(
			&source.Kind{Type: &corev1.Pod{}},
			&handler.EnqueueRequestsFromMapFunc{
				ToRequests: handler.ToRequestsFunc(r.impsMatchingPod),
			}).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			&handler.EnqueueRequestsFromMapFunc{
				ToRequests: handler.ToRequestsFunc(r.impsReferencingSecret),
			})

	return builder.Complete(r)
}

func (r *ImagePullSecretReconciler) impsMatchingNamespace(obj handler.MapObject) []ctrl.Request {
	ns, ok := obj.Object.(*corev1.Namespace)
	if !ok {
		r.Log.Info("object is not a Namespace")
		return []ctrl.Request{}
	}

	impsList := &v1alpha1.ImagePullSecretList{}

	err := r.Client.List(context.TODO(), impsList)
	if err != nil {
		r.Log.Info(err.Error())
		return []ctrl.Request{}
	}

	var res []ctrl.Request
	for _, imps := range impsList.Items {
		matches, err := imps.MatchesNamespace(ns)
		if err != nil {
			r.Log.Info("cannot match imps against namespace", map[string]interface{}{
				"imps":      imps,
				"namespace": ns,
				"error":     err,
			})
			continue
		}

		if matches {
			res = append(res, ctrl.Request{
				NamespacedName: types.NamespacedName{
					Name:      imps.GetName(),
					Namespace: imps.GetNamespace(),
				},
			})
		}
	}
	return res
}

func (r *ImagePullSecretReconciler) impsMatchingPod(obj handler.MapObject) []ctrl.Request {
	pod, ok := obj.Object.(*corev1.Pod)
	if !ok {
		r.Log.Info("object is not a Pod or Namespace")
		return []ctrl.Request{}
	}

	// If the namespace containing the pod matches, let's not add the pod to the reconciliation queue.
	// This prevents reconciliations to start on each pod startup when the namespace selectors are properly used.
	podsNamespace := &corev1.Namespace{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Name: pod.Namespace,
	}, podsNamespace)

	if err != nil {
		r.Log.Warn("cannot get pod's namespace, please authorize the controller to list namespaces, or too many reconcilations will be executed", map[string]interface{}{
			"error":     err,
			"namespace": pod.Namespace,
		})
		podsNamespace = nil
	}

	impsList := &v1alpha1.ImagePullSecretList{}

	err = r.Client.List(context.TODO(), impsList)
	if err != nil {
		r.Log.Info(err.Error())
		return []ctrl.Request{}
	}
	var res []ctrl.Request
	for _, imps := range impsList.Items {
		matches, err := imps.MatchesPod(pod)
		if err != nil {
			r.Log.Info("cannot match imps against pod", map[string]interface{}{
				"imps":  imps,
				"pod":   pod,
				"error": err,
			})
			continue
		}

		if matches {
			if podsNamespace != nil {
				nsMatches, err := imps.MatchesNamespace(podsNamespace)
				if err != nil {
					r.Log.Info("cannot match imps against namespace", map[string]interface{}{
						"imps":      imps,
						"namespace": podsNamespace,
						"error":     err,
					})
				} else if nsMatches {
					continue
				}
			}
			res = append(res, ctrl.Request{
				NamespacedName: types.NamespacedName{
					Name:      imps.GetName(),
					Namespace: imps.GetNamespace(),
				},
			})
		}
	}
	return res
}

func (r *ImagePullSecretReconciler) impsReferencingSecret(obj handler.MapObject) []ctrl.Request {
	secret, ok := obj.Object.(*corev1.Secret)
	if !ok {
		r.Log.Info("object is not a Secret")
		return []ctrl.Request{}
	}

	impsList := &v1alpha1.ImagePullSecretList{}

	err := r.Client.List(context.TODO(), impsList)
	if err != nil {
		r.Log.Info(err.Error())
		return []ctrl.Request{}
	}

	var res []ctrl.Request
	for _, imps := range impsList.Items {
		if imps.Spec.Registry.Credentials.Name == secret.Name && imps.Spec.Registry.Credentials.Namespace == secret.Namespace {
			res = append(res, ctrl.Request{
				NamespacedName: types.NamespacedName{
					Name:      imps.GetName(),
					Namespace: imps.GetNamespace(),
				},
			})
		}
	}
	return res
}
