package pullsecrets

import (
	"context"
	"testing"

	"gotest.tools/assert"

	"github.com/banzaicloud/imps/api/common"
)

func TestStaticLoginCredentialProvider_NewStaticLoginCredentialProvider(t *testing.T) {
	t.Parallel()
	type args struct {
		parsedDockerConfig common.DockerRegistryConfig
	}

	tests := []struct {
		name   string
		args   args
		wanted StaticLoginCredentialProvider
	}{
		{
			name: "basic functionality test",
			args: args{
				parsedDockerConfig: common.DockerRegistryConfig{
					Auths: map[string]common.LoginCredentials{
						"testCreds": {
							Username: "testUser",
						},
					},
				},
			},
			wanted: StaticLoginCredentialProvider{
				[]LoginCredentialsWithDetails{
					{
						LoginCredentials: common.LoginCredentials{
							Username: "testUser",
						},
						URL: "testCreds",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := NewStaticLoginCredentialProvider(tt.args.parsedDockerConfig)

			assert.DeepEqual(t, tt.wanted, found)
		})
	}
}

func TestStaticLoginCredentialProvider_LoginCredentials(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                          string
		staticLoginCredentialProvider StaticLoginCredentialProvider
		wanted                        []LoginCredentialsWithDetails
	}{
		{
			name: "basic functionality test",
			staticLoginCredentialProvider: StaticLoginCredentialProvider{
				[]LoginCredentialsWithDetails{
					{
						LoginCredentials: common.LoginCredentials{
							Username: "testUser",
						},
						URL: "testCreds",
					},
				},
			},
			wanted: []LoginCredentialsWithDetails{
				{
					LoginCredentials: common.LoginCredentials{
						Username: "testUser",
					},
					URL: "testCreds",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found, err := tt.staticLoginCredentialProvider.LoginCredentials(context.Background())

			assert.DeepEqual(t, tt.wanted, found)
			assert.NilError(t, err)
		})
	}
}
