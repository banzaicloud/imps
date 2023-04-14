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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/banzaicloud/imps/api/common"
	impsEcr "github.com/banzaicloud/imps/pkg/ecr"
)

type ECRLoginCredentialsProvider struct {
	Credentials aws.Credentials
	Region      string
	AccountID   string
	RoleArn     string
}

func NewECRLoginCredentialsProvider(accountID, region, keyID, secretAccessKey string, roleArn string) ECRLoginCredentialsProvider {
	return ECRLoginCredentialsProvider{
		Credentials: aws.Credentials{
			AccessKeyID:     keyID,
			SecretAccessKey: secretAccessKey,
		},
		AccountID: accountID,
		Region:    region,
		RoleArn:   roleArn,
	}
}

func (p ECRLoginCredentialsProvider) GetURL() string {
	return fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", p.AccountID, p.Region)
}

func (p ECRLoginCredentialsProvider) LoginCredentials(ctx context.Context) ([]LoginCredentialsWithDetails, error) {
	token, err := impsEcr.GetAuthorizationToken(ctx, p.Region, p.Credentials, p.RoleArn)
	if err != nil {
		return nil, err
	}
	decodedAuth, err := base64.StdEncoding.DecodeString(*token.AuthorizationToken)
	if err != nil {
		return nil, err
	}

	splitAuth := strings.SplitN(string(decodedAuth), ":", 2)

	return []LoginCredentialsWithDetails{
		{
			LoginCredentials: common.LoginCredentials{
				Username: splitAuth[0],
				Password: splitAuth[1],
				Auth:     *token.AuthorizationToken,
			},
			URL:        p.GetURL(),
			Expiration: token.ExpiresAt,
		},
	}, nil
}
