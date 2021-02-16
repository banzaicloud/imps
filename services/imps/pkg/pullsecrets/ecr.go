// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
package pullsecrets

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ECRLoginCredentialsProvider struct {
	Credentials aws.Credentials
	Region      string // TODO: type
	AccountIDs  []string
}

func (p ECRLoginCredentialsProvider) LoginCredentials(ctx context.Context) (LoginCredentials, error) {
	client := ecr.NewFromConfig(aws.Config{
		Region: p.Region,
		Credentials: aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return p.Credentials, nil
		}),
	})

	// note: RegistryIds is deprecated, any account's registries can be accessed via the returned token
	authToken, err := client.GetAuthorizationToken(ctx, ecr.GetAuthorizationTokenInput{})
}
