package pullsecrets

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/banzaicloud/imps/api/common"
)

func TestDockerConfig_NewDockerRegistryConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want common.DockerRegistryConfig
	}{
		{
			name: "basic functionality test",
			want: common.DockerRegistryConfig{
				Auths: make(map[string]common.LoginCredentials),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := NewDockerRegistryConfig()

			assert.DeepEqual(t, tt.want, found)
		})
	}
}

func TestDockerConfig_NewConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "basic functionality test",
			want: &Config{
				Registries: make(map[string]LoginCredentialProvider),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := NewConfig()

			assert.DeepEqual(t, tt.want, found)
		})
	}
}

func TestDockerConfig_StaticProviderFromDockerConfig(t *testing.T) {
	t.Parallel()
	type args struct {
		data []byte
	}

	tests := []struct {
		name   string
		args   args
		config Config
		want   LoginCredentialProvider
	}{
		{
			name: "basic functionality test",
			args: args{
				data: []byte("{}"),
			},
			config: Config{},
			want: StaticLoginCredentialProvider{
				Credentials: []LoginCredentialsWithDetails{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := tt.config.StaticProviderFromDockerConfig(tt.args.data)

			assert.DeepEqual(t, tt.want, found)
		})
	}
}

func TestDockerConfig_getOptionalFieldFromMap(t *testing.T) {
	t.Parallel()
	type args struct {
		data       map[string][]byte
		key        string
		defaultVal string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "value not found in data",
			args: args{
				data:       map[string][]byte{},
				key:        "testKey",
				defaultVal: "testDefault",
			},
			want: "testDefault",
		},
		{
			name: "value is found in data",
			args: args{
				data: map[string][]byte{
					"testKey": []byte("testData"),
				},
				key:        "testKey",
				defaultVal: "testDefault",
			},
			want: "testData",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := getOptionalFieldFromMap(tt.args.data, tt.args.key, tt.args.defaultVal)

			assert.Equal(t, tt.want, found)
		})
	}
}

func TestDockerConfig_getFieldFromMap(t *testing.T) {
	t.Parallel()
	type args struct {
		data map[string][]byte
		key  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "value not found in data",
			args: args{
				data: map[string][]byte{},
				key:  "testKey",
			},
			want: "",
		},
		{
			name: "value is found in data",
			args: args{
				data: map[string][]byte{
					"testKey": []byte("testData"),
				},
				key: "testKey",
			},
			want: "testData",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found, err := getFieldFromMap(tt.args.data, tt.args.key)

			assert.Equal(t, tt.want, found)
			if err != nil {
				assert.ErrorContains(t, err, "no such key")
			}
		})
	}
}

func TestDockerConfig_ECRProviderFromSecret(t *testing.T) {
	t.Parallel()
	type args struct {
		data map[string][]byte
	}

	tests := []struct {
		name   string
		args   args
		config Config
		want   LoginCredentialProvider
	}{
		{
			name: "basic functionality test",
			args: args{
				data: map[string][]byte{
					"accountID":   []byte("testAccountID"),
					"region":      []byte("testRegion"),
					"accessKeyID": []byte("testAccessKeyID"),
					"secretKey":   []byte("testSecretKey"),
					"roleArn":     []byte("testRole"),
				},
			},
			config: Config{},
			want: ECRLoginCredentialsProvider{
				Region:    "testRegion",
				AccountID: "testAccountID",
				RoleArn:   "testRole",
				Credentials: aws.Credentials{
					AccessKeyID:     "testAccessKeyID",
					SecretAccessKey: "testSecretKey",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := tt.config.ECRProviderFromSecret(tt.args.data)

			assert.DeepEqual(t, tt.want, found)
		})
	}
}

func TestDockerConfig_ResultingDockerConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		config Config
		want   *ResultingDockerConfig
	}{
		{
			name:   "empty config",
			config: Config{},
			want: &ResultingDockerConfig{
				ErrorsPerSecret: ErrorsPerSecret{},
				ConfigContents:  []uint8(`{"auths":{}}`),
			},
		},
		{
			name: "non-empty config",
			config: Config{
				Registries: map[string]LoginCredentialProvider{
					"testProvider": StaticLoginCredentialProvider{
						Credentials: []LoginCredentialsWithDetails{
							{
								URL: "test.url",
							},
						},
					},
				},
			},
			want: &ResultingDockerConfig{
				ErrorsPerSecret: ErrorsPerSecret{"testProvider": nil},
				ConfigContents:  []byte(`{"auths":{"test.url":{"username":"","password":"","auth":""}}}`),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found, err := tt.config.ResultingDockerConfig(context.Background())

			assert.DeepEqual(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}

func TestDockerConfig_AsSecret(t *testing.T) {
	t.Parallel()
	type args struct {
		secretNamespace string
		secretName      string
	}

	tests := []struct {
		name                  string
		args                  args
		resultingDockerConfig ResultingDockerConfig
		want                  *corev1.Secret
	}{
		{
			name: "basic functionality test",
			args: args{
				secretName:      "testSecret",
				secretNamespace: "testSecretNamespace",
			},
			want: &corev1.Secret{
				ObjectMeta: v1.ObjectMeta{
					Name:      "testSecret",
					Namespace: "testSecretNamespace",
				},
				Type: common.SecretTypeBasicAuth,
				StringData: map[string]string{
					common.SecretKeyDockerConfig: "",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := tt.resultingDockerConfig.AsSecret(tt.args.secretNamespace, tt.args.secretName)

			assert.DeepEqual(t, tt.want, found)
		})
	}
}
