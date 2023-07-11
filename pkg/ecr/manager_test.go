package ecr

import (
	"context"
	"testing"
	"time"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
	"logur.dev/logur"
)

func TestTokenManager_GetAuthorizationToken(t *testing.T) {
	t.Parallel()
	type args struct {
		key    StringableCredentials
		client ClientInterface
	}

	testTokenName := "testToken"

	tests := []struct {
		name            string
		args            args
		mockTokenOutput *ecr.GetAuthorizationTokenOutput
		tokenManager    *TokenManager
		wanted          types.AuthorizationData
		expectedErr     error
	}{
		{
			name: "basic functionality test",
			args: args{
				key: StringableCredentials{
					aws.Credentials{
						AccessKeyID: "testAccessKeyID",
					},
					"testRegion",
					"testRole",
				},
				client: &MockECRClient{},
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
				Logger:        &logur.TestLogger{},
			},
			wanted: types.AuthorizationData{
				AuthorizationToken: &testTokenName,
			},
			expectedErr: nil,
		},
		{
			name: "error from NewECRToken",
			args: args{
				key: StringableCredentials{
					aws.Credentials{
						AccessKeyID: "testAccessKeyID",
					},
					"testRegion",
					"testRole",
				},
				client: &MockECRClient{},
			},
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{},
			},
			tokenManager: &TokenManager{
				ManagedTokens: map[string]*Token{},
				Logger:        &logur.TestLogger{},
			},
			wanted:      types.AuthorizationData{},
			expectedErr: errors.New("no authorization data is returned from ECR"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockClient := &MockECRClient{}
			mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(tt.mockTokenOutput, nil)

			found, err := tt.tokenManager.GetAuthorizationToken(context.Background(), tt.args.key, mockClient)

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
	t.Parallel()
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
	t.Parallel()
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
				Logger: &logur.TestLogger{},
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
						CurrentToken: &types.AuthorizationData{
							AuthorizationToken: &testTokenName,
							ExpiresAt:          &expiryTime,
						},
						TokenValidityDuration: 10 * time.Minute,
					},
				},
				Logger: &logur.TestLogger{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockClient := &MockECRClient{}
			mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(tt.mockTokenOutput, nil)
			tt.tokenManager.ManagedTokens["testName"].Client = mockClient

			tt.tokenManager.updateTokens()

			assert.Equal(t, len(tt.tokenManager.ManagedTokens), 1)
			assert.DeepEqual(t, tt.tokenManager.ManagedTokens["testName"].CurrentToken, &types.AuthorizationData{
				AuthorizationToken: &testTokenName,
				ExpiresAt:          &newExpiryTime,
			})
		})
	}
}

func TestTokenManager_NewECRTokenManager(t *testing.T) {
	t.Parallel()
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
				logger: &logur.TestLogger{},
			},
			want: &TokenManager{
				ManagedTokens: map[string]*Token{},
				Logger:        &logur.TestLogger{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := NewECRTokenManager(tt.args.logger)

			assert.Equal(t, len(tt.want.ManagedTokens), len(found.ManagedTokens))
		})
	}
}
