package pullsecrets

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/banzaicloud/imps/api/common"
	impsEcr "github.com/banzaicloud/imps/pkg/ecr"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
	"logur.dev/logur"
)

func TestECRLoginCredentialsProvider_NewECRLoginCredentialsProvider(t *testing.T) {
	type args struct {
		accountID       string
		region          string
		keyID           string
		secretAccessKey string
		roleArn         string
	}

	tests := []struct {
		name   string
		args   args
		wanted ECRLoginCredentialsProvider
	}{
		{
			name: "basic functionality test",
			args: args{
				accountID:       "testAccountID",
				region:          "testRegion",
				keyID:           "testKeyID",
				secretAccessKey: "testSecretAccessKey",
				roleArn:         "testRole",
			},
			wanted: ECRLoginCredentialsProvider{
				Region:    "testRegion",
				AccountID: "testAccountID",
				RoleArn:   "testRole",
				Credentials: aws.Credentials{
					SecretAccessKey: "testSecretAccessKey",
					AccessKeyID:     "testKeyID",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := NewECRLoginCredentialsProvider(tt.args.accountID, tt.args.region, tt.args.keyID, tt.args.secretAccessKey, tt.args.roleArn, nil)

			assert.DeepEqual(t, tt.wanted, found)
		})
	}
}

func TestECRLoginCredentialsProvider_GetURL(t *testing.T) {

	tests := []struct {
		name                        string
		ecrLoginCredentialsProvider ECRLoginCredentialsProvider
		wanted                      string
	}{
		{
			name:                        "empty provider",
			ecrLoginCredentialsProvider: ECRLoginCredentialsProvider{},
			wanted:                      ".dkr.ecr..amazonaws.com",
		},
		{
			name: "non-empty provider",
			ecrLoginCredentialsProvider: ECRLoginCredentialsProvider{
				AccountID: "testAccountID",
				Region:    "testRegion",
			},
			wanted: "testAccountID.dkr.ecr.testRegion.amazonaws.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.ecrLoginCredentialsProvider.GetURL()

			assert.DeepEqual(t, tt.wanted, found)
		})
	}
}

type MockECRClient struct {
	mock.Mock
}

func (m *MockECRClient) GetAuthorizationToken(ctx context.Context, input *ecr.GetAuthorizationTokenInput, _ ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*ecr.GetAuthorizationTokenOutput), args.Error(1)
}

func TestECRLoginCredentialsProvider_LoginCredentials(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	mockClient := &MockECRClient{}
	testTokenName := base64.StdEncoding.EncodeToString([]byte("testUser:testPass"))
	mockTokenOutput := &ecr.GetAuthorizationTokenOutput{
		AuthorizationData: []types.AuthorizationData{
			{
				AuthorizationToken: &testTokenName,
			},
		},
	}
	mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(mockTokenOutput, nil)

	tests := []struct {
		name                        string
		args                        args
		ecrLoginCredentialsProvider ECRLoginCredentialsProvider
		wanted                      []LoginCredentialsWithDetails
	}{
		{
			name: "basic functionality test",
			args: args{
				ctx: context.Background(),
			},
			ecrLoginCredentialsProvider: ECRLoginCredentialsProvider{
				Region:    "testRegion",
				AccountID: "testAccountID",
				RoleArn:   "testRole",
				Credentials: aws.Credentials{
					SecretAccessKey: "testSecretAccessKey",
					AccessKeyID:     "testKeyID",
				},
				Client: mockClient,
			},
			wanted: []LoginCredentialsWithDetails{
				{
					LoginCredentials: common.LoginCredentials{
						Username: "testUser",
						Password: "testPass",
						Auth:     "dGVzdFVzZXI6dGVzdFBhc3M=",
					},
					URL: "testAccountID.dkr.ecr.testRegion.amazonaws.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impsEcr.Initialize(logur.NewTestLogger())
			found, err := tt.ecrLoginCredentialsProvider.LoginCredentials(tt.args.ctx)

			assert.DeepEqual(t, tt.wanted, found)
			assert.NilError(t, err)
		})
	}
}
