// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package pullsecrets

import (
	"context"
	"encoding/base64"
	"time"
)

type StaticLoginCredentialProvider struct {
	Username string
	Password string
	URL      string
}

func NewStaticLoginCredentialProvider(url, username, password string) StaticLoginCredentialProvider {
	return StaticLoginCredentialProvider{
		Username: username,
		Password: password,
		URL:      url,
	}
}

func (p StaticLoginCredentialProvider) GetURL() string {
	return p.URL
}

func (p StaticLoginCredentialProvider) LoginCredentials(ctx context.Context) (*LoginCredentials, *time.Time, error) {
	return &LoginCredentials{
		Username: p.Username,
		Password: p.Password,
		Auth:     base64.StdEncoding.EncodeToString([]byte(p.Username + ":" + p.Password)),
	}, nil, nil
}
