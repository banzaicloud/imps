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

	"emperror.dev/errors"
	ctrlBuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/banzaicloud/imps/pkg/pullsecrets"

	"emperror.dev/emperror"
	"logur.dev/logur"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/banzaicloud/imps/internal/cron"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
)

// RefresherReconciler reconciles a AlertingPolicy object
type RefresherReconciler struct {
	client.Client
	Log          logur.Logger
	ErrorHandler emperror.ErrorHandler
	Scheme       *runtime.Scheme

	ResourceReconciler        reconciler.ResourceReconciler
	PeriodicReconcileInterval time.Duration
	SourceSecrets             []types.NamespacedName
	TargetSecret              types.NamespacedName
}

// +kubebuilder:rbac:groups=images.banzaicloud.io,resources=imagepullsecrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=images.banzaicloud.io,resources=imagepullsecrets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=create;update;patch
func (r *RefresherReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	result, err := r.reconcile(req)
	result, err = cron.EnsurePeriodicReconcile(r.PeriodicReconcileInterval, result, err)
	if err != nil {
		r.ErrorHandler.Handle(err)
	}
	return result, err
}

func (r *RefresherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	builder := ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Secret{}, ctrlBuilder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			&handler.EnqueueRequestsFromMapFunc{
				ToRequests: handler.ToRequestsFunc(r.isMatchingSecret),
			})

	return builder.Complete(r)
}

func (r *RefresherReconciler) isMatchingSecret(obj handler.MapObject) []ctrl.Request {
	secret, ok := obj.Object.(*corev1.Secret)
	if !ok {
		r.Log.Info("object is not a Secret")
		return []ctrl.Request{}
	}

	reconcileTargetRequest := ctrl.Request{
		NamespacedName: r.TargetSecret,
	}

	secretRef := types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}

	if r.isSourceSecret(secretRef) || r.isTargetSecret(secretRef) {
		return []ctrl.Request{reconcileTargetRequest}
	}

	return []ctrl.Request{}
}

func (r *RefresherReconciler) isSourceSecret(secret types.NamespacedName) bool {
	for _, sourceSecret := range r.SourceSecrets {
		if sourceSecret.Namespace == secret.Namespace && sourceSecret.Name == secret.Name {
			return true
		}
	}

	return false
}

func (r *RefresherReconciler) isTargetSecret(secret types.NamespacedName) bool {
	return r.TargetSecret.Namespace == secret.Namespace && r.TargetSecret.Name == secret.Name
}

func (r *RefresherReconciler) reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := logur.WithField(r.Log, "imagepullsecret", req.NamespacedName)
	result := ctrl.Result{}

	if !r.isSourceSecret(req.NamespacedName) && !r.isTargetSecret(req.NamespacedName) {
		return result, nil
	}

	config, err := pullsecrets.NewConfigFromSecrets(ctx, r, r.SourceSecrets)
	if err != nil {
		return result, errors.WithStack(err)
	}

	pullSecret, pullSecretExpires, err := config.Secret(ctx, r.TargetSecret.Namespace, r.TargetSecret.Name)
	if err != nil {
		return result, errors.WrapWithDetails(err, "cannot get referenced secret")
	}

	_, err = r.ResourceReconciler.ReconcileResource(pullSecret, reconciler.StatePresent)
	if err != nil {
		return result, err
	}

	logger.Info("successfully reconciled secret", map[string]interface{}{
		"expiration": pullSecretExpires,
	})

	return result, nil
}
