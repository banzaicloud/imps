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
		r.Log.Info("object is not a Pod")
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
		matches, err := imps.MatchesPod(pod)
		if err != nil {
			r.Log.Info("cannot match imps against a pod", map[string]interface{}{
				"imps":  imps,
				"pod":   pod,
				"error": err,
			})
		}
		if matches {
			// TODO: only add matches if the namespace does not match, to decrease reconciliation
			// steps
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

// TODO: trigger reconciliation if secret changes
