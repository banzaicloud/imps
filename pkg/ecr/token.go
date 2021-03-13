// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecr

import (
	"context"
	"time"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecr_types "github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const (
	// assumedTokenValidityTime specifies how long to consider the returned token to be valid if not specified in
	// the response
	assumedTokenValidityTime = 20 * time.Minute
)

type Token struct {
	Creds                 StringableCredentials
	CurrentToken          *ecr_types.AuthorizationData
	TokenValidityDuration time.Duration
	LastQueriedAt         time.Time
}

func NewECRToken(ctx context.Context, creds StringableCredentials) (*Token, error) {
	token := &Token{
		Creds: creds,
	}

	err := token.Refresh(ctx)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *Token) Refresh(ctx context.Context) error {
	client := ecr.NewFromConfig(aws.Config{
		Region: t.Creds.Region,
		Credentials: aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     t.Creds.AccessKeyID,
				SecretAccessKey: t.Creds.SecretAccessKey,
			}, nil
		}),
	})

	// note: RegistryIds is deprecated, any account's registries can be accessed via the returned token
	authToken, err := client.GetAuthorizationToken(ctx, &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return err
	}

	if len(authToken.AuthorizationData) == 0 {
		return errors.New("no authorization data is returned from ECR")
	}

	if len(authToken.AuthorizationData) > 1 {
		// This should never happen according to current API specs
		return errors.NewWithDetails("multiple authorization records are returned for ECR", "response", authToken)
	}

	if authToken.AuthorizationData[0].AuthorizationToken == nil {
		return errors.New("no authorization data is returned from ECR - authorization token is empty")
	}

	fetchedToken := authToken.AuthorizationData[0]
	t.CurrentToken = &fetchedToken
	if fetchedToken.ExpiresAt != nil {
		t.TokenValidityDuration = time.Until(*fetchedToken.ExpiresAt)
	} else {
		t.TokenValidityDuration = assumedTokenValidityTime
	}

	return nil
}
