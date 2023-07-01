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
	"sync"
	"time"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"logur.dev/logur"
)

const (
	discardUnqueriedTokensAfter = 60 * time.Minute
)

type TokenManager struct {
	sync.Mutex
	ManagedTokens map[string]*Token
	Logger        logur.Logger
}

func Initialize(logger logur.Logger) {
	tokenManager = NewECRTokenManager(logger)
	tokenManager.start()
}

var tokenManager *TokenManager

func GetAuthorizationToken(ctx context.Context, region string, credentials aws.Credentials, roleArn string, client ECRClientInterface) (ecrTypes.AuthorizationData, error) {
	return tokenManager.GetAuthorizationToken(ctx, StringableCredentials{
		Credentials: credentials,
		Region:      region,
		RoleArn:     roleArn,
	}, client)
}

func NewECRTokenManager(logger logur.Logger) *TokenManager {
	return &TokenManager{
		ManagedTokens: map[string]*Token{},
		Logger:        logger,
	}
}

func (t *TokenManager) start() {
	go t.tokenUpdater()
}

func (t *TokenManager) tokenUpdater() {
	c := time.NewTicker(5 * time.Second)
	for range c.C {
		t.updateTokens()
		t.discardOldTokens()
	}
}

func (t *TokenManager) updateTokens() {
	t.Lock()
	defer t.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	for _, token := range t.ManagedTokens {
		if token.CurrentToken == nil || time.Until(*token.CurrentToken.ExpiresAt) < token.TokenValidityDuration/2 {
			err := token.Refresh(ctx)
			if err != nil {
				t.Logger.Warn("error refreshing token", map[string]interface{}{
					"error":             err,
					"aws_access_key_id": token.Creds.AccessKeyID,
					"region":            token.Creds.Region,
				})
			} else {
				t.Logger.Info("token refreshed", map[string]interface{}{
					"aws_access_key_id": token.Creds.AccessKeyID,
					"region":            token.Creds.Region,
				})
			}
		}
	}
}

func (t *TokenManager) discardOldTokens() {
	t.Lock()
	defer t.Unlock()

	for id, token := range t.ManagedTokens {
		if time.Since(token.LastQueriedAt) > discardUnqueriedTokensAfter {
			delete(t.ManagedTokens, id)
		}
	}
}

func (t *TokenManager) GetAuthorizationToken(ctx context.Context, key StringableCredentials, client ECRClientInterface) (ecrTypes.AuthorizationData, error) {
	t.Lock()
	defer t.Unlock()
	token, found := t.ManagedTokens[key.String()]
	if !found {
		token, err := NewECRToken(ctx, key, client)
		if err != nil {
			return ecrTypes.AuthorizationData{}, err
		}
		t.ManagedTokens[key.String()] = token
		if token.CurrentToken == nil {
			return ecrTypes.AuthorizationData{}, errors.New("no token is available")
		}
		t.Logger.Info("token refreshed", map[string]interface{}{
			"aws_access_key_id": token.Creds.AccessKeyID,
			"region":            token.Creds.Region,
		})

		return *token.CurrentToken, nil
	}

	token.LastQueriedAt = time.Now()

	return *token.CurrentToken, nil
}
