package ecr

import (
	"context"
	"sync"
	"time"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	ecr_types "github.com/aws/aws-sdk-go-v2/service/ecr/types"
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

var tokenManager *TokenManager = nil

func GetAuthorizationToken(ctx context.Context, region string, credentials aws.Credentials) (ecr_types.AuthorizationData, error) {
	return tokenManager.GetAuthorizationToken(ctx, StringableCredentials{
		Credentials: credentials,
		Region:      region,
	})
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

func (t *TokenManager) GetAuthorizationToken(ctx context.Context, key StringableCredentials) (ecr_types.AuthorizationData, error) {
	t.Lock()
	defer t.Unlock()
	token, found := t.ManagedTokens[key.String()]
	if !found {
		token, err := NewECRToken(ctx, key)
		if err != nil {
			return ecr_types.AuthorizationData{}, err
		}
		t.ManagedTokens[key.String()] = token
		if token.CurrentToken != nil {
			return ecr_types.AuthorizationData{}, errors.New("no token is available")
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
