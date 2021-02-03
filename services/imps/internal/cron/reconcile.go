// Copyright (c) 2019 Banzai Cloud Zrt. All Rights Reserved.

package cron

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/types"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/banzaicloud/backyards/services/imps/api/v1alpha1"
	"github.com/banzaicloud/backyards/services/imps/controllers"
)

// TODO[1.5]: When merging use the central package instead of this one

type PeriodicReconciler struct {
	client.Client
	log           logur.Logger
	impsReconciler *controllers.ImagePullSecretReconciler
}

func NewPeriodicReconciler(log logur.Logger, client client.Client, imps *controllers.ImagePullSecretReconciler) *PeriodicReconciler {
	return &PeriodicReconciler{
		Client:        client,
		log:           log,
		impsReconciler: imps,
	}
}

func (p *PeriodicReconciler) Reconcile(interval int) manager.Runnable {
	// TODO: stop channel?
	return manager.RunnableFunc(func(s <-chan struct{}) error {
		go func() {
			t := time.NewTicker(time.Second * time.Duration(interval))
			for range t.C {
				ctx := context.Background()

				p.log.Info("periodic reconcile")

				var impss v1alpha1.ImagePullSecretList
				err := p.List(ctx, &impss)
				if err != nil {
					p.log.Error("periodic reconcile failed: failed to list imagepullsecrets", map[string]interface{}{"error": err})
				}

				for _, imps := range impss.Items {
					req := types.NamespacedName{
						Name: imps.Name,
					}
					_, err := p.impsReconciler.Reconcile(ctrl.Request{
						NamespacedName: req,
					})
					if err != nil {
						p.log.Error("periodic reconcile failed", map[string]interface{}{"controller": "imagepullsecrets", "resource": req, "error": err})
					}
				}
			}
		}()
		return nil
	})
}
