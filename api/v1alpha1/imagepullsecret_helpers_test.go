package v1alpha1

import (
	"testing"

	"gotest.tools/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestImagePullSecret_GetOwnerReferenceForOwnedObject(t *testing.T) {

	testImagePullSecret := ImagePullSecret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "testAPIVersion",
			Kind:       "ImagePullSecret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testImagePullSecret",
			Namespace: "testNamespace",
			UID:       "testUID",
		},
	}

	tests := []struct {
		name            string
		imagePullSecret ImagePullSecret
		want            metav1.OwnerReference
	}{
		{
			name:            "basic functionality check",
			imagePullSecret: testImagePullSecret,
			want: metav1.OwnerReference{
				APIVersion: "testAPIVersion",
				Kind:       "ImagePullSecret",
				Name:       "testImagePullSecret",
				UID:        "testUID",
				Controller: BoolPointer(false),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.imagePullSecret.GetOwnerReferenceForOwnedObject()

			assert.DeepEqual(t, tt.want, found)
		})
	}
}

func TestRegistryConfig_CredentialsAsNamespacedNameList(t *testing.T) {

	testRegistryConfig := RegistryConfig{
		Credentials: []NamespacedName{
			{
				Name:      "testCred",
				Namespace: "testNamespace",
			},
			{
				Name:      "testCred2",
				Namespace: "testNamespace",
			},
			{
				Name:      "testCred3",
				Namespace: "testNamespace2",
			},
		},
	}

	tests := []struct {
		name           string
		registryConfig RegistryConfig
		want           []types.NamespacedName
	}{
		{
			name:           "basic functionality check",
			registryConfig: testRegistryConfig,
			want: []types.NamespacedName{
				{
					Name:      "testCred",
					Namespace: "testNamespace",
				},
				{
					Name:      "testCred2",
					Namespace: "testNamespace",
				},
				{
					Name:      "testCred3",
					Namespace: "testNamespace2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.registryConfig.CredentialsAsNamespacedNameList()

			assert.DeepEqual(t, tt.want, found)
		})
	}
}
