// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package controllers

import (
	"context"

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

// TODO: For now let's make sure this passes linting
// nolint: unparam
func (r *ImagePullSecretReconciler) reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	logger := logur.WithField(r.Log, "imagepullsecret", req.NamespacedName)
	result := ctrl.Result{}

	logger.Info("Reconciling ImagePullSecret finished")
	return result, nil
}

func (r *ImagePullSecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	builder := ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ImagePullSecret{}, ctrlBuilder.WithPredicates(predicate.GenerationChangedPredicate{}))

	return builder.Complete(r)
}
