package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	testlogur "logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestRefresherReconciler_isTargetSecret(t *testing.T) {
	type args struct {
		secret types.NamespacedName
	}

	testRefresherReconciler := RefresherReconciler{}
	testRefresherReconciler.TargetSecret = types.NamespacedName{
		Name:      "testSecretName",
		Namespace: "testSecretNamespace",
	}

	tests := []struct {
		name                string
		refresherReconciler RefresherReconciler
		args                args
		want                bool
	}{
		{
			name:                "name and namespace both match",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName",
					Namespace: "testSecretNamespace",
				},
			},
			want: true,
		},
		{
			name:                "name and namespace mismatch",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName2",
					Namespace: "testSecretName2",
				},
			},
			want: false,
		},
		{
			name:                "name mismatch",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName2",
					Namespace: "testSecretName",
				},
			},
			want: false,
		},
		{
			name:                "namespace mismatch",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName",
					Namespace: "testSecretName2",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.refresherReconciler.isTargetSecret(tt.args.secret)

			assert.Equal(t, tt.want, found)
		})
	}
}

func TestRefresherReconciler_isSourceSecret(t *testing.T) {
	type args struct {
		secret types.NamespacedName
	}

	testRefresherReconciler := RefresherReconciler{}
	testRefresherReconciler.SourceSecrets = []types.NamespacedName{
		{
			Name:      "testSecretName",
			Namespace: "testSecretNamespace",
		},
		{
			Name:      "testSecretName2",
			Namespace: "testSecretNamespace",
		},
		{
			Name:      "testSecretName3",
			Namespace: "testSecretNamespace2",
		},
	}

	tests := []struct {
		name                string
		refresherReconciler RefresherReconciler
		args                args
		want                bool
	}{
		{
			name:                "secret is found",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName",
					Namespace: "testSecretNamespace",
				},
			},
			want: true,
		},
		{
			name:                "secret doesn't exist in the namespace",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName2",
					Namespace: "testSecretNamespace2",
				},
			},
			want: false,
		},
		{
			name:                "secret name isn't on the list",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName4",
					Namespace: "testSecretNamespace",
				},
			},
			want: false,
		},
		{
			name:                "secret name and namespace aren't on the list",
			refresherReconciler: testRefresherReconciler,
			args: args{
				secret: types.NamespacedName{
					Name:      "testSecretName4",
					Namespace: "testSecretNamespace3",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.refresherReconciler.isSourceSecret(tt.args.secret)

			assert.Equal(t, tt.want, found)
		})
	}
}

func TestRefresherReconciler_isMatchingSecret(t *testing.T) {
	type args struct {
		obj client.Object
	}

	testRefresherReconciler := RefresherReconciler{}
	testRefresherReconciler.SourceSecrets = []types.NamespacedName{
		{
			Name:      "testSecretName",
			Namespace: "testSecretNamespace",
		},
	}
	testRefresherReconciler.TargetSecret = types.NamespacedName{
		Name:      "testSecretName2",
		Namespace: "testSecretNamespace2",
	}
	testRefresherReconciler.Log = testlogur.NewTestLogger()

	tests := []struct {
		name                string
		refresherReconciler RefresherReconciler
		args                args
		want                []ctrl.Request
	}{
		{
			name:                "object is a secret and is found a source secret",
			refresherReconciler: testRefresherReconciler,
			args: args{
				obj: client.Object(&corev1.Secret{
					TypeMeta: metav1.TypeMeta{
						Kind: "Secret",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testSecretName",
						Namespace: "testSecretNamespace",
					},
				}),
			},
			want: []ctrl.Request{
				{
					NamespacedName: testRefresherReconciler.TargetSecret,
				},
			},
		},
		{
			name:                "object is a secret and is found as a target secret",
			refresherReconciler: testRefresherReconciler,
			args: args{
				obj: client.Object(&corev1.Secret{
					TypeMeta: metav1.TypeMeta{
						Kind: "Secret",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testSecretName2",
						Namespace: "testSecretNamespace2",
					},
				}),
			},
			want: []ctrl.Request{
				{
					NamespacedName: testRefresherReconciler.TargetSecret,
				},
			},
		},
		{
			name:                "object isn't a secret",
			refresherReconciler: testRefresherReconciler,
			args: args{
				obj: client.Object(&corev1.Pod{
					TypeMeta: metav1.TypeMeta{
						Kind: "Pod",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testSecretName",
						Namespace: "testSecretNamespace",
					},
				}),
			},
			want: []ctrl.Request{},
		},
		{
			name:                "object is a secret but isn't found",
			refresherReconciler: testRefresherReconciler,
			args: args{
				obj: client.Object(&corev1.Secret{
					TypeMeta: metav1.TypeMeta{
						Kind: "Secret",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testSecretName3",
						Namespace: "testSecretNamespace",
					},
				}),
			},
			want: []ctrl.Request{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.refresherReconciler.isMatchingSecret(tt.args.obj)

			assert.Equal(t, tt.want, found)
		})
	}
}
