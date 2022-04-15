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
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type StringableCredentials struct {
	aws.Credentials
	// Region specifies which region to connect to when using this credential
	Region string
	// Assume a role
	RoleArn string
}

func (c *StringableCredentials) GetCreds(ctx context.Context) (aws.Credentials, error) {
	return c.Credentials, nil
}

func (c *StringableCredentials) ToAwsConfig() aws.Config {
	cfg := aws.Config{
		Region: c.Region,
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return c.Credentials, nil
		}),
	}
	if len(c.RoleArn) != 0 {
		// Create the credentials from AssumeRoleProvider to assume the role
		// referenced by the `RoleArn`.
		stsSvc := sts.NewFromConfig(cfg)
		creds := stscreds.NewAssumeRoleProvider(stsSvc, c.RoleArn)
		cfg.Credentials = aws.NewCredentialsCache(creds)
	}
	return cfg
}

func (c *StringableCredentials) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return c.Credentials, nil
}

func (c StringableCredentials) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", c.Region, c.AccessKeyID, c.SecretAccessKey, c.SessionToken)
}
