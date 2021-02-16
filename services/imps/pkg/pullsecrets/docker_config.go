// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
package pullsecrets

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	LoginCredentials(context.Context) (*LoginCredentials, *time.Time, error)
	GetURL() string
}

type Config struct {
	Registries map[string]LoginCredentialProvider
}

func NewConfig() *Config {
	return &Config{
		Registries: map[string]LoginCredentialProvider{},
	}
}

func NewConfigFromSecrets(ctx context.Context, c client.Client, refs []types.NamespacedName) (*Config, error) {
	var secret corev1.Secret
	config := NewConfig()

	for _, secretRef := range refs {
		err := c.Get(ctx, client.ObjectKey{
			Namespace: secretRef.Namespace,
			Name:      secretRef.Name,
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
		default:
			return nil, errors.NewWithDetails("unknown secret type", "type", secret.Type, "secret", secret.ObjectMeta)
		}
	}
	return config, nil
}

func (c *Config) AddRegistryWithUsernamePassword(url, username, password string) {
	c.Registries[url] = NewStaticLoginCredentialProvider(url, username, password)
}

func (c *Config) AddRegistriesFromDockerConfig(data []byte) error {
	var dockerConfig DockerRegistryConfig
	err := json.Unmarshal(data, &dockerConfig)
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

func getFieldFromMap(data map[string][]byte, key string) (string, error) {
	value, found := data[key]
	if !found {
		return "", errors.NewWithDetails("no such key", "key", key)
	}

	return string(value), nil
}

func (c *Config) AddECRFromSecret(data map[string][]byte) error {
	accountID, err := getFieldFromMap(data, ECRSecretAccountID)
	if err != nil {
		return err
	}

	region, err := getFieldFromMap(data, ECRSecretRegion)
	if err != nil {
		return err
	}

	accKeyID, err := getFieldFromMap(data, ECRSecretKeyAccessKeyID)
	if err != nil {
		return err
	}

	secretKey, err := getFieldFromMap(data, ECRSecretSecretKey)
	if err != nil {
		return err
	}

	c.AddECR(accountID, region, accKeyID, secretKey)
	return nil
}

func (c Config) AddECR(accountID, region, accessKeyID, secretKey string) {
	provider := NewECRLoginCredentialsProvider(accountID, region, accessKeyID, secretKey)
	c.Registries[provider.GetURL()] = provider
}

func (c Config) ConfigString(ctx context.Context) ([]byte, *time.Time, error) {
	finalRegistryConfig := NewDockerRegistryConfig()
	var minExpiration *time.Time = nil

	for url, provider := range c.Registries {
		creds, expiration, err := provider.LoginCredentials(ctx)
		if err != nil {
			return nil, nil, err
		}

		if expiration != nil {
			if minExpiration == nil || (*expiration).Before(*minExpiration) {
				minExpiration = expiration
			}
		}
		finalRegistryConfig.Auths[url] = *creds
	}

	marshaledRegistryConfig, err := json.Marshal(finalRegistryConfig)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot encode docker configuration into a JSON")
	}

	return marshaledRegistryConfig, minExpiration, nil
}
