// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

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
