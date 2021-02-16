// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package pullsecrets

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// nolint:gosec
const (
	SecretTypeBasicAuth      = "kubernetes.io/dockerconfigjson"
	SecretTypeECRCredentials = "banzaicloud.io/aws-ecr-login-config"

	SecretKeyDockerConfig = ".dockerconfigjson"

	ECRSecretKeyRegion      = "region"
	ECRSecretKeyAccountID   = "accountID"
	ECRSecretKeyAccessKeyID = "accessKeyID"
	ECRSecretSecretKey      = "secretKey"
)

func NewBasicAuthSecret(secretNamespace, secretName, registry, user, password string) (*corev1.Secret, error) {
	config := NewConfig()
	config.AddRegistryWithUsernamePassword(registry, user, password)

	dockerJSON, err := config.ConfigString()
	if err != nil {
		return nil, err
	}

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Type: SecretTypeBasicAuth,
		StringData: map[string]string{
			SecretKeyDockerConfig: string(dockerJSON),
		},
	}, nil
}

func NewECRLoginCredentialsSecret(secretNamespace, secretName, accountID, region, awsAccessKeyID, awsSecretAccessKey string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Type: SecretTypeECRCredentials,
		StringData: map[string]string{
			ECRSecretKeyRegion:      region,
			ECRSecretKeyAccountID:   accountID,
			ECRSecretKeyAccessKeyID: awsAccessKeyID,
			ECRSecretSecretKey:      awsSecretAccessKey,
		},
	}
}
