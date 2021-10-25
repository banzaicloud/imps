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
	"encoding/json"
	"fmt"
	"github.com/banzaicloud/imps/api/common"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"emperror.dev/errors"
)

func NewDockerRegistryConfig() common.DockerRegistryConfig {
	return common.DockerRegistryConfig{
		Auths: map[string]common.LoginCredentials{},
	}
}

type LoginCredentialsWithDetails struct {
	common.LoginCredentials
	Expiration *time.Time
	URL        string
}

type LoginCredentialProvider interface {
	LoginCredentials(context.Context) ([]LoginCredentialsWithDetails, error)
}

type Config struct {
	Registries map[string]LoginCredentialProvider
}

func NewConfig() *Config {
	return &Config{
		Registries: map[string]LoginCredentialProvider{},
	}
}

func NewConfigFromSecrets(ctx context.Context, c client.Client, refs []types.NamespacedName) *Config {
	var secret corev1.Secret
	config := NewConfig()

	for _, secretRef := range refs {
		secretName := fmt.Sprintf("%s.%s", secretRef.Namespace, secretRef.Name)
		err := c.Get(ctx, client.ObjectKey{
			Namespace: secretRef.Namespace,
			Name:      secretRef.Name,
		}, &secret)

		if err != nil {
			config.Registries[secretName] = NewErroredCredentialProvider(err)
			continue
		}

		switch secret.Type {
		case common.SecretTypeBasicAuth:
			dockerConfig, found := secret.Data[common.SecretKeyDockerConfig]
			if !found {
				config.Registries[secretName] = NewErroredCredentialProvider(
					errors.NewWithDetails("no docker configuration found in secret", "secret", secret.ObjectMeta))
				continue
			}
			config.Registries[secretName] = config.StaticProviderFromDockerConfig(dockerConfig)
		case common.SecretTypeECRCredentials:
			config.Registries[secretName] = config.ECRProviderFromSecret(secret.Data)
		default:
			config.Registries[secretName] = NewErroredCredentialProvider(
				errors.NewWithDetails("unknown secret type", "type", secret.Type, "secret", secret.ObjectMeta))
		}
	}
	return config
}

func (c *Config) StaticProviderFromDockerConfig(data []byte) LoginCredentialProvider {
	var dockerConfig common.DockerRegistryConfig
	err := json.Unmarshal(data, &dockerConfig)
	if err != nil {
		return NewErroredCredentialProvider(err)
	}

	return NewStaticLoginCredentialProvider(dockerConfig)
}

func getFieldFromMap(data map[string][]byte, key string) (string, error) {
	value, found := data[key]
	if !found {
		return "", fmt.Errorf("no such key: %s", key)
	}

	return string(value), nil
}

func (c *Config) ECRProviderFromSecret(data map[string][]byte) LoginCredentialProvider {
	accountID, err := getFieldFromMap(data, common.ECRSecretAccountID)
	if err != nil {
		return NewErroredCredentialProvider(err)
	}

	region, err := getFieldFromMap(data, common.ECRSecretRegion)
	if err != nil {
		return NewErroredCredentialProvider(err)
	}

	accKeyID, err := getFieldFromMap(data, common.ECRSecretKeyAccessKeyID)
	if err != nil {
		return NewErroredCredentialProvider(err)
	}

	secretKey, err := getFieldFromMap(data, common.ECRSecretSecretKey)
	if err != nil {
		return NewErroredCredentialProvider(err)
	}

	return NewECRLoginCredentialsProvider(accountID, region, accKeyID, secretKey)
}

type ResultingDockerConfig struct {
	ErrorsPerSecret
	ConfigContents []byte
	Expiration     *time.Time
}

func (c Config) ResultingDockerConfig(ctx context.Context) (*ResultingDockerConfig, error) {
	finalRegistryConfig := NewDockerRegistryConfig()
	var minExpiration *time.Time = nil
	secretErrors := NewErrorsPerSecret()

	for secret, provider := range c.Registries {
		secretErrors.AddSecret(secret)
		credentials, err := provider.LoginCredentials(ctx)
		if err != nil {
			secretErrors.SetSecretError(secret, err)
			continue
		}

		for _, credential := range credentials {
			if credential.Expiration != nil {
				if minExpiration == nil || (credential.Expiration).Before(*minExpiration) {
					minExpiration = credential.Expiration
				}
			}
			finalRegistryConfig.Auths[credential.URL] = credential.LoginCredentials
		}
	}

	marshaledRegistryConfig, err := json.Marshal(finalRegistryConfig)
	if err != nil {
		return nil, errors.Wrap(err, "cannot encode docker configuration into a JSON")
	}

	return &ResultingDockerConfig{
		ErrorsPerSecret: secretErrors,
		ConfigContents:  marshaledRegistryConfig,
		Expiration:      minExpiration,
	}, nil
}

func (c ResultingDockerConfig) AsSecret(secretNamespace, secretName string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Type: common.SecretTypeBasicAuth,
		StringData: map[string]string{
			common.SecretKeyDockerConfig: string(c.ConfigContents),
		},
	}
}
