package ecr

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"gotest.tools/assert"
)

func TestStringableCredentials_GetCreds(t *testing.T) {

	tests := []struct {
		name                  string
		stringableCredentials *StringableCredentials
		want                  aws.Credentials
	}{
		{
			name:                  "empty credentials",
			stringableCredentials: &StringableCredentials{},
			want:                  aws.Credentials{},
		},
		{
			name: "non-empty credentials",
			stringableCredentials: &StringableCredentials{
				aws.Credentials{
					AccessKeyID: "testAccessKeyID",
				},
				"testRegion",
				"testRole",
			},
			want: aws.Credentials{
				AccessKeyID: "testAccessKeyID",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := tt.stringableCredentials.GetCreds(context.Background())

			assert.Equal(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}

func TestStringableCredentials_ToAwsConfig(t *testing.T) {

	tests := []struct {
		name                  string
		stringableCredentials *StringableCredentials
		want                  aws.Config
	}{
		{
			name: "basic functionality test",
			stringableCredentials: &StringableCredentials{
				Region:  "testRegion",
				RoleArn: "testRole",
			},
			want: aws.Config{
				Region:      "testRegion",
				Credentials: aws.NewCredentialsCache(stscreds.NewAssumeRoleProvider(&sts.Client{}, "testRole")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.stringableCredentials.ToAwsConfig()

			assert.Equal(t, tt.want.Region, found.Region)
			wantedCred, _ := tt.want.Credentials.Retrieve(context.Background())
			foundCreds, _ := found.Credentials.Retrieve(context.Background())
			assert.Equal(t, wantedCred, foundCreds)
		})
	}
}

func TestStringableCredentials_Retrieve(t *testing.T) {

	tests := []struct {
		name                  string
		stringableCredentials *StringableCredentials
		want                  aws.Credentials
	}{
		{
			name:                  "empty credentials",
			stringableCredentials: &StringableCredentials{},
			want:                  aws.Credentials{},
		},
		{
			name: "non-empty credentials",
			stringableCredentials: &StringableCredentials{
				aws.Credentials{
					AccessKeyID: "testAccessKeyID",
				},
				"testRegion",
				"testRole",
			},
			want: aws.Credentials{
				AccessKeyID: "testAccessKeyID",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := tt.stringableCredentials.Retrieve(context.Background())

			assert.Equal(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}

func TestStringableCredentials_String(t *testing.T) {

	tests := []struct {
		name                  string
		stringableCredentials *StringableCredentials
		want                  string
	}{
		{
			name:                  "empty credentials",
			stringableCredentials: &StringableCredentials{},
			want:                  "///",
		},
		{
			name: "partially empty credentials",
			stringableCredentials: &StringableCredentials{
				aws.Credentials{
					AccessKeyID: "testAccessKeyID",
				},
				"testRegion",
				"testRole",
			},
			want: "testRegion/testAccessKeyID//",
		},
		{
			name: "non-empty credentials",
			stringableCredentials: &StringableCredentials{
				aws.Credentials{
					AccessKeyID:     "testAccessKeyID",
					SecretAccessKey: "testSecretAccessKey",
					SessionToken:    "testSessionToken",
				},
				"testRegion",
				"testRole",
			},
			want: "testRegion/testAccessKeyID/testSecretAccessKey/testSessionToken",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.stringableCredentials.String()

			assert.Equal(t, tt.want, found)
		})
	}
}
