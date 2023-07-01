package ecr

import (
	"context"
	"testing"
	"time"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
	"logur.dev/logur"
)

func TestTokenManager_GetAuthorizationToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		key    StringableCredentials
		client ECRClientInterface
	}

	testTokenName := "testToken"

	tests := []struct {
		name            string
		args            args
		mockTokenOutput *ecr.GetAuthorizationTokenOutput
		tokenManager    *TokenManager
		wanted          ecrTypes.AuthorizationData
		expectedErr     error
	}{
		{
			name: "basic functionality test",
			args: args{
				ctx: context.Background(),
				key: StringableCredentials{
					aws.Credentials{
						AccessKeyID: "testAccessKeyID",
					},
					"testRegion",
					"testRole",
				},
			},
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{
					{
						AuthorizationToken: &testTokenName,
					},
				},
			},
			tokenManager: &TokenManager{
				ManagedTokens: map[string]*Token{},
				Logger:        logur.NewTestLogger(),
			},
			wanted: ecrTypes.AuthorizationData{
				AuthorizationToken: &testTokenName,
			},
			expectedErr: nil,
		},
		{
			name: "error from NewECRToken",
			args: args{
				ctx: context.Background(),
				key: StringableCredentials{
					aws.Credentials{
						AccessKeyID: "testAccessKeyID",
					},
					"testRegion",
					"testRole",
				},
			},
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{},
			},
			tokenManager: &TokenManager{
				ManagedTokens: map[string]*Token{},
				Logger:        logur.NewTestLogger(),
			},
			wanted:      ecrTypes.AuthorizationData{},
			expectedErr: errors.New("no authorization data is returned from ECR"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockECRClient{}
			mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(tt.mockTokenOutput, nil)

			found, err := tt.tokenManager.GetAuthorizationToken(tt.args.ctx, tt.args.key, mockClient)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.DeepEqual(t, tt.wanted, found)
				assert.NilError(t, err)
			}
		})
	}
}

func TestTokenManager_discardOldTokens(t *testing.T) {

	currentTime := time.Now()

	tests := []struct {
		name         string
		tokenManager *TokenManager
	}{
		{
			name: "basic functionality test",
			tokenManager: &TokenManager{
				ManagedTokens: map[string]*Token{
					"tokenToDiscard": {
						LastQueriedAt: currentTime.Add(-70 * time.Minute),
					},
					"tokenNotToDiscard": {
						LastQueriedAt: currentTime.Add(-40 * time.Minute),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tokenManager.discardOldTokens()

			assert.DeepEqual(t, tt.tokenManager.ManagedTokens, map[string]*Token{
				"tokenNotToDiscard": {
					LastQueriedAt: currentTime.Add(-40 * time.Minute),
				},
			})
		})
	}
}

func TestTokenManager_updateTokens(t *testing.T) {

	testTokenName := "testToken"
	expiryTime := time.Now().Add(5 * time.Minute)
	newExpiryTime := expiryTime.Add(5 * time.Minute)

	tests := []struct {
		name            string
		mockTokenOutput *ecr.GetAuthorizationTokenOutput
		tokenManager    *TokenManager
	}{
		{
			name: "currentToken is nil",
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{
					{
						AuthorizationToken: &testTokenName,
						ExpiresAt:          &newExpiryTime,
					},
				},
			},
			tokenManager: &TokenManager{
				ManagedTokens: map[string]*Token{
					"testName": {
						CurrentToken: nil,
					},
				},
				Logger: logur.NewTestLogger(),
			},
		},
		{
			name: "less than half of the validity duration is left",
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{
					{
						AuthorizationToken: &testTokenName,
						ExpiresAt:          &newExpiryTime,
					},
				},
			},
			tokenManager: &TokenManager{
				ManagedTokens: map[string]*Token{
					"testName": {
						CurrentToken: &ecrTypes.AuthorizationData{
							AuthorizationToken: &testTokenName,
							ExpiresAt:          &expiryTime,
						},
						TokenValidityDuration: 10 * time.Minute,
					},
				},
				Logger: logur.NewTestLogger(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockECRClient{}
			mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(tt.mockTokenOutput, nil)
			tt.tokenManager.ManagedTokens["testName"].Client = mockClient

			tt.tokenManager.updateTokens()

			assert.Equal(t, len(tt.tokenManager.ManagedTokens), 1)
			assert.DeepEqual(t, tt.tokenManager.ManagedTokens["testName"].CurrentToken, &ecrTypes.AuthorizationData{
				AuthorizationToken: &testTokenName,
				ExpiresAt:          &newExpiryTime,
			})
		})
	}
}

func TestTokenManager_NewECRTokenManager(t *testing.T) {
	type args struct {
		logger logur.Logger
	}

	tests := []struct {
		name string
		args args
		want *TokenManager
	}{
		{
			name: "basic functionality test",
			args: args{
				logger: logur.NewTestLogger(),
			},
			want: &TokenManager{
				ManagedTokens: map[string]*Token{},
				Logger:        logur.NewTestLogger(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := NewECRTokenManager(tt.args.logger)

			assert.Equal(t, len(tt.want.ManagedTokens), len(found.ManagedTokens))
		})
	}
}
