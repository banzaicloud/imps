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

package pullsecrets

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ECRLoginCredentialsProvider struct {
	Credentials aws.Credentials
	Region      string
	AccountID   string
}

func NewECRLoginCredentialsProvider(accountID, region, keyID, secretAccessKey string) ECRLoginCredentialsProvider {
	return ECRLoginCredentialsProvider{
		Credentials: aws.Credentials{
			AccessKeyID:     keyID,
			SecretAccessKey: secretAccessKey,
		},
		AccountID: accountID,
		Region:    region,
	}
}

func (p ECRLoginCredentialsProvider) GetURL() string {
	return fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", p.AccountID, p.Region)
}

func (p ECRLoginCredentialsProvider) LoginCredentials(ctx context.Context) (*LoginCredentials, *time.Time, error) {
	client := ecr.NewFromConfig(aws.Config{
		Region: p.Region,
		Credentials: aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return p.Credentials, nil
		}),
	})

	// note: RegistryIds is deprecated, any account's registries can be accessed via the returned token
	authToken, err := client.GetAuthorizationToken(ctx, &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return nil, nil, err
	}

	if len(authToken.AuthorizationData) == 0 {
		return nil, nil, errors.New("no authorization data is returned from ECR")
	}

	if len(authToken.AuthorizationData) > 1 {
		// This should never happen according to current API specs
		return nil, nil, errors.NewWithDetails("multiple authorization records are returned for ECR", "response", authToken)
	}

	if authToken.AuthorizationData[0].AuthorizationToken == nil {
		return nil, nil, errors.New("no authorization data is returned from ECR - authorization token is empty")
	}

	decodedAuth, err := base64.StdEncoding.DecodeString(*authToken.AuthorizationData[0].AuthorizationToken)
	if err != nil {
		return nil, nil, err
	}

	splitAuth := strings.SplitN(string(decodedAuth), ":", 2)
	return &LoginCredentials{
		Username: splitAuth[0],
		Password: splitAuth[1],
		Auth:     *authToken.AuthorizationData[0].AuthorizationToken,
	}, authToken.AuthorizationData[0].ExpiresAt, nil
}
