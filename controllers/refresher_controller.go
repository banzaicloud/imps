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
	"emperror.dev/errors"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlBuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/banzaicloud/imps/internal/cron"
	"github.com/banzaicloud/imps/pkg/pullsecrets"
)

// RefresherReconciler reconciles a AlertingPolicy object
type RefresherReconciler struct {
	client.Client
	Log          logur.Logger
	ErrorHandler emperror.ErrorHandler

	ResourceReconciler        reconciler.ResourceReconciler
	PeriodicReconcileInterval time.Duration
	SourceSecrets             []types.NamespacedName
	TargetSecret              types.NamespacedName
}

// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;patch;delete
func (r *RefresherReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	result, err := r.reconcile(ctx, req)
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
			handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
				return r.isMatchingSecret(object)
			}))

	return builder.Complete(r)
}

func (r *RefresherReconciler) isMatchingSecret(obj client.Object) []ctrl.Request {
	secret, ok := obj.(*corev1.Secret)
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

func (r *RefresherReconciler) reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logur.WithField(r.Log, "imagepullsecret", req.NamespacedName)
	result := ctrl.Result{}

	if !r.isSourceSecret(req.NamespacedName) && !r.isTargetSecret(req.NamespacedName) {
		return result, nil
	}

	config := pullsecrets.NewConfigFromSecrets(ctx, r, r.SourceSecrets)
	resultingConfig, err := config.ResultingDockerConfig(ctx)
	if err != nil {
		return result, errors.WithStack(err)
	}

	pullSecret := resultingConfig.AsSecret(r.TargetSecret.Namespace, r.TargetSecret.Name)

	_, err = r.ResourceReconciler.ReconcileResource(pullSecret, reconciler.StatePresent)
	if err != nil {
		return result, err
	}

	if err := resultingConfig.AsError(); err != nil {
		for secret, status := range resultingConfig.AsStatus() {
			if status != pullsecrets.SourceSecretStatus {
				logger.Warn("secret failed to render", map[string]interface{}{
					"source_secret": secret,
					"reason":        status,
				})
			}
		}

		return result, err
	}

	logger.Info("successfully reconciled secret", map[string]interface{}{
		"expiration": resultingConfig.Expiration,
	})

	return result, nil
}
