// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
package pullsecrets

import (
	"encoding/json"

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

func NewConfig() Config {
	return Config{
		Registries: map[string]LoginCredentialProvider{},
	}
}

func (c *Config) AddRegistryWithUsernamePassword(url, username, password string) {
	c.Registries[url] = NewStaticLoginCredentialProvider(username, password)
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
