package common

import (
	"testing"

	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSecret_NewBasicAuthSecret(t *testing.T) {
	type args struct {
		secretNamespace string
		secretName      string
		registry        string
		user            string
		password        string
	}

	tests := []struct {
		name string
		args args
		want *corev1.Secret
	}{
		{
			name: "secret generation works as expected",
			args: args{
				secretNamespace: "testSecretNamespace",
				secretName:      "testSecretName",
				registry:        "test.io",
				user:            "testUser",
				password:        "testPassword",
			},
			want: &corev1.Secret{
				Type: "kubernetes.io/dockerconfigjson",
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testSecretName",
					Namespace: "testSecretNamespace",
				},
				StringData: map[string]string{
					".dockerconfigjson": "{\"auths\":{\"test.io\":{\"username\":\"testUser\",\"password\":\"testPassword\",\"auth\":\"dGVzdFVzZXI6dGVzdFBhc3N3b3Jk\"}}}",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := NewBasicAuthSecret(tt.args.secretNamespace, tt.args.secretName, tt.args.registry, tt.args.user, tt.args.password)

			assert.DeepEqual(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}

func TestSecret_NewECRLoginCredentialsSecret(t *testing.T) {
	type args struct {
		secretNamespace    string
		secretName         string
		accountID          string
		region             string
		awsAccessKeyID     string
		awsSecretAccessKey string
	}

	tests := []struct {
		name string
		args args
		want *corev1.Secret
	}{
		{
			name: "secret generation works as expected",
			args: args{
				secretNamespace:    "testSecretNamespace",
				secretName:         "testSecretName",
				accountID:          "testAccountID",
				region:             "testRegion",
				awsAccessKeyID:     "testKeyID",
				awsSecretAccessKey: "testSecretAccessKey",
			},
			want: &corev1.Secret{
				Type: "banzaicloud.io/aws-ecr-login-config",
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testSecretName",
					Namespace: "testSecretNamespace",
				},
				StringData: map[string]string{
					"region":      "testRegion",
					"accountID":   "testAccountID",
					"accessKeyID": "testKeyID",
					"secretKey":   "testSecretAccessKey",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := NewECRLoginCredentialsSecret(tt.args.secretNamespace, tt.args.secretName, tt.args.accountID, tt.args.region, tt.args.awsAccessKeyID, tt.args.awsSecretAccessKey)

			assert.DeepEqual(t, tt.want, found)
		})
	}
}
