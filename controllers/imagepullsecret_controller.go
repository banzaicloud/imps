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
	"time"

	"emperror.dev/emperror"
	"github.com/banzaicloud/imps/api/v1alpha1"
	"github.com/banzaicloud/imps/internal/cron"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlBuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// AlertingPolicyReconciler reconciles a AlertingPolicy object
type ImagePullSecretReconciler struct {
	client.Client
	Log          logur.Logger
	ErrorHandler emperror.ErrorHandler
	Recorder     record.EventRecorder

	ResourceReconciler        reconciler.ResourceReconciler
	PeriodicReconcileInterval time.Duration
}

// +kubebuilder:rbac:groups=images.banzaicloud.io,resources=imagepullsecrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=images.banzaicloud.io,resources=imagepullsecrets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=create;update;patch
func (r *ImagePullSecretReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	result, err := r.reconcile(ctx, req)
	result, err = cron.EnsurePeriodicReconcile(r.PeriodicReconcileInterval, result, err)
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
			handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
				return r.impsMatchingNamespace(object)
			})).
		Watches(
			&source.Kind{Type: &corev1.Pod{}},
			handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
				return r.impsMatchingPod(object)
			})).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
				return r.impsReferencingSecret(object)
			}))

	return builder.Complete(r)
}

func (r *ImagePullSecretReconciler) impsMatchingNamespace(obj client.Object) []ctrl.Request {
	ns, ok := obj.(*corev1.Namespace)
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
			res = append(res, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      imps.GetName(),
					Namespace: imps.GetNamespace(),
				},
			})
		}
	}

	return res
}

func (r *ImagePullSecretReconciler) impsMatchingPod(obj client.Object) []ctrl.Request {
	pod, ok := obj.(*corev1.Pod)
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

func (r *ImagePullSecretReconciler) impsReferencingSecret(obj client.Object) []ctrl.Request {
	secret, ok := obj.(*corev1.Secret)
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
		for _, credential := range imps.Spec.Registry.Credentials {
			if credential.Name == secret.Name && credential.Namespace == secret.Namespace {
				res = append(res, ctrl.Request{
					NamespacedName: types.NamespacedName{
						Name:      imps.GetName(),
						Namespace: imps.GetNamespace(),
					},
				})
			}
		}
	}

	return res
}
