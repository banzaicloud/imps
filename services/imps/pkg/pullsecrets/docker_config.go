// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
package pullsecrets

import (
	"context"
	"encoding/base64"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"

	"emperror.dev/errors"
)

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Auth     string `json:"auth"` // base64 encoded username:password
}

// DockerRegistryConfig represents a docker compliant image pull secret json file
type DockerRegistryConfig struct {
	Auths map[string]LoginCredentials `json:"auths"`
}

func NewDockerRegistryConfig() DockerRegistryConfig {
	return DockerRegistryConfig{
		Auths: map[string]LoginCredentials{},
	}
}

type LoginCredentialProvider interface {
	LoginCredentials() (LoginCredentials, error)
}

type Config struct {
	Registries map[string]LoginCredentialProvider
}

func NewConfig() *Config {
	return &Config{
		Registries: map[string]LoginCredentialProvider{},
	}
}

func NewConfigFromSecrets(ctx context.Context, c client.Client,refs []types.NamespacedName) (*Config, error) {
	var secret corev1.Secret
	config := NewConfig()

	for _, secretRef := range refs {
		err := c.Get(ctx, client.ObjectKey{
			Namespace: secretRef.Namespace,
			Name: secretRef.Name,
		}, &secret)
		if err != nil {
			return nil, err
		}

		switch secret.Type {
		case SecretTypeBasicAuth:
			dockerConfig, found := secret.Data[SecretKeyDockerConfig]
			if !found {
				return nil, errors.NewWithDetails("no docker configuration found in secret", "secret", secret.ObjectMeta)

			}
			err = config.AddRegistriesFromDockerConfig(dockerConfig)
			if err != nil {
				return nil, err
			}
		case SecretTypeECRCredentials:
			err = config.AddECRFromSecret(secret.Data)
			if err != nil {
				return nil, err
			}
			// TODO
		default:
			return nil, errors.NewWithDetails("unknown secret type", "type", secret.Type, "secret", secret.ObjectMeta)
		}
	}
	// TODO
}

func (c *Config) AddRegistryWithUsernamePassword(url, username, password string) {
	c.Registries[url] = NewStaticLoginCredentialProvider(username, password)
}

func (c *Config) AddRegistriesFromDockerConfig(data []byte) error {
	decodedJson, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}

	var dockerConfig DockerRegistryConfig
	err = json.Unmarshal(decodedJson, &dockerConfig)
	if err != nil {
		return err
	}

	for url, registry := range dockerConfig.Auths {
		decodedAuth, err := base64.StdEncoding.DecodeString(registry.Auth)
		if err != nil {
			return err
		}

		splitAuth := strings.SplitN(string(decodedAuth), ":", 2)

		c.AddRegistryWithUsernamePassword(url, splitAuth[0], splitAuth[1])
	}
	return nil
}

func decodeFieldFromMap(data map[string][]byte, key string) (string, error) {
	value, found := data[key]
	if !found {
		return "", errors.NewWithDetails("no such key", "key", key)
	}

	decodedData, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return "", err
	}

	return string(decodedData), nil
}

func (c *Config) AddECRFromSecret(data map[string][]byte) error {
	accountId, err :=decodeFieldFromMap(data, ECRSecretAccountID)
	if err != nil {
		return err
	}

	region, err := decodeFieldFromMap(data, ECRSecretRegion)
	if err != nil {
		return err
	}

	accKeyID, err := decodeFieldFromMap(data, ECRSecretKeyAccessKeyID)
	if err != nil {
		return err
	}

	secretKey, err := decodeFieldFromMap(data, ECRSecretSecretKey)
	if err != nil {
		return err
	}

	c.AddECR(accountId, region, accKeyID, secretKey)
	return nil
}

func (c Config) AddECR(accountID, region, accessKeyID, secretKey string) {

}

func (c Config) ConfigString() ([]byte, error) {
	finalRegistryConfig := NewDockerRegistryConfig()

	for url, provider := range c.Registries {
		creds, err := provider.LoginCredentials()
		if err != nil {
			return nil, err
		}
		finalRegistryConfig.Auths[url] = creds
	}

	marshaledRegistryConfig, err := json.Marshal(finalRegistryConfig)
	if err != nil {
		return nil, errors.Wrap(err, "cannot encode docker configuration into a JSON")
	}

	return marshaledRegistryConfig, nil
}
