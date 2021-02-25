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

package pullsecrets

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// nolint:gosec
const (
	SecretTypeBasicAuth      = "kubernetes.io/dockerconfigjson"
	SecretTypeECRCredentials = "banzaicloud.io/aws-ecr-login-config"

	SecretKeyDockerConfig = ".dockerconfigjson"

	ECRSecretRegion         = "region"
	ECRSecretAccountID      = "accountID"
	ECRSecretKeyAccessKeyID = "accessKeyID"
	ECRSecretSecretKey      = "secretKey"
)

func NewBasicAuthSecret(secretNamespace, secretName, registry, user, password string) (*corev1.Secret, error) {
	config := NewConfig()
	config.AddRegistryWithUsernamePassword(registry, user, password)

	dockerJSON, _, err := config.ConfigString(context.Background())
	if err != nil {
		return nil, err
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Type: SecretTypeBasicAuth,
		StringData: map[string]string{
			SecretKeyDockerConfig: string(dockerJSON),
		},
	}
	secret.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Secret"))
	return secret, nil
}

func NewECRLoginCredentialsSecret(secretNamespace, secretName, accountID, region, awsAccessKeyID, awsSecretAccessKey string) *corev1.Secret {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Type: SecretTypeECRCredentials,
		StringData: map[string]string{
			ECRSecretRegion:         region,
			ECRSecretAccountID:      accountID,
			ECRSecretKeyAccessKeyID: awsAccessKeyID,
			ECRSecretSecretKey:      awsSecretAccessKey,
		},
	}
	secret.SetGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Secret"))
	return secret
}
