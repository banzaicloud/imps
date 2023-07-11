package ecr

import (
	"context"
	"testing"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

type MockECRClient struct {
	mock.Mock
}

func (m *MockECRClient) GetAuthorizationToken(ctx context.Context, input *ecr.GetAuthorizationTokenInput, _ ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {
	args := m.Called(ctx, input)
	// nolint:forcetypeassert
	return args.Get(0).(*ecr.GetAuthorizationTokenOutput), args.Error(1)
}

func TestToken_NewECRToken(t *testing.T) {
	t.Parallel()
	type args struct {
		creds  StringableCredentials
		client ClientInterface
	}

	mockClient := &MockECRClient{}
	testTokenName := "testToken"
	mockTokenOutput := &ecr.GetAuthorizationTokenOutput{
		AuthorizationData: []types.AuthorizationData{
			{
				AuthorizationToken: &testTokenName,
			},
		},
	}
	mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(mockTokenOutput, nil)

	tests := []struct {
		name        string
		args        args
		want        *Token
		expectedErr error
	}{
		{
			name: "basic functionality test",
			args: args{
				creds:  StringableCredentials{},
				client: mockClient,
			},
			want: &Token{
				CurrentToken: &types.AuthorizationData{
					AuthorizationToken: &testTokenName,
				},
			},
			expectedErr: nil,
		},
		{
			name: "no token returned",
			args: args{
				creds:  StringableCredentials{},
				client: nil,
			},
			want:        &Token{},
			expectedErr: errors.New("operation error ECR: GetAuthorizationToken, failed to resolve service endpoint, an AWS region is required, but was not found"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found, err := NewECRToken(context.Background(), tt.args.creds, tt.args.client)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.DeepEqual(t, tt.want.CurrentToken, found.CurrentToken)
				assert.NilError(t, err)
			}
		})
	}
}

func TestToken_Refresh(t *testing.T) {
	t.Parallel()

	testTokenName := "testToken"

	tests := []struct {
		name            string
		mockTokenOutput *ecr.GetAuthorizationTokenOutput
		token           *Token
		expectedErr     error
	}{
		{
			name: "basic functionality test",
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{
					{
						AuthorizationToken: &testTokenName,
					},
				},
			},
			token:       &Token{},
			expectedErr: nil,
		},
		{
			name: "no authorization data",
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: nil,
			},
			token:       &Token{},
			expectedErr: errors.New("no authorization data is returned from ECR"),
		},
		{
			name: "multiple authorization records",
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{
					{
						AuthorizationToken: &testTokenName,
					},
					{
						AuthorizationToken: &testTokenName,
					},
				},
			},
			token:       &Token{},
			expectedErr: errors.New("multiple authorization records are returned for ECR"),
		},
		{
			name: "authorization token is empty",
			mockTokenOutput: &ecr.GetAuthorizationTokenOutput{
				AuthorizationData: []types.AuthorizationData{{}},
			},
			token:       &Token{},
			expectedErr: errors.New("no authorization data is returned from ECR - authorization token is empty"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockClient := &MockECRClient{}
			mockClient.On("GetAuthorizationToken", mock.Anything, mock.Anything).Return(tt.mockTokenOutput, nil)
			tt.token.Client = mockClient

			err := tt.token.Refresh(context.Background())

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
