package cron

import (
	"testing"
	"time"

	"emperror.dev/errors"
	"gotest.tools/assert"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestCron_calculateRequeueAfter(t *testing.T) {
	type args struct {
		result                    reconcile.Result
		periodicReconcileInterval time.Duration
	}

	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "result is zero",
			args: args{
				result:                    reconcile.Result{},
				periodicReconcileInterval: 5 * time.Second,
			},
			want: 5 * time.Second,
		},
		{
			name: "requeue after is smaller than periodicReconcileInterval",
			args: args{
				result: reconcile.Result{
					RequeueAfter: 3 * time.Second,
				},
				periodicReconcileInterval: 5 * time.Second,
			},
			want: 3 * time.Second,
		},
		{
			name: "requeue after is greater than periodicReconcileInterval",
			args: args{
				result: reconcile.Result{
					RequeueAfter: 7 * time.Second,
				},
				periodicReconcileInterval: 5 * time.Second,
			},
			want: 5 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := calculateRequeueAfter(tt.args.result, tt.args.periodicReconcileInterval)

			assert.Equal(t, tt.want, found)
		})
	}
}

func TestCron_EnsurePeriodicReconcile(t *testing.T) {
	type args struct {
		interval time.Duration
		result   reconcile.Result
		err      error
	}

	tests := []struct {
		name        string
		args        args
		want        reconcile.Result
		expectedErr error
	}{
		{
			name: "basic functionality test",
			args: args{
				interval: 5 * time.Second,
				result:   reconcile.Result{},
				err:      nil,
			},
			want: reconcile.Result{
				Requeue:      true,
				RequeueAfter: 5 * time.Second,
			},
			expectedErr: nil,
		},
		{
			name: "error is not nil",
			args: args{
				interval: 5 * time.Second,
				result:   reconcile.Result{},
				err:      errors.New("test"),
			},
			want: reconcile.Result{
				Requeue:      false,
				RequeueAfter: 0 * time.Second,
			},
			expectedErr: errors.New("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := EnsurePeriodicReconcile(tt.args.interval, tt.args.result, tt.args.err)

			assert.Equal(t, tt.want, found)
			if tt.expectedErr != nil {
				assert.Error(t, err, tt.expectedErr.Error())
			}
		})
	}
}
