// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package pullsecrets

import "encoding/base64"

type StaticLoginCredentialProvider struct {
	Username string
	Password string
}

func NewStaticLoginCredentialProvider(username, password string) StaticLoginCredentialProvider {
	return StaticLoginCredentialProvider{
		Username: username,
		Password: password,
	}
}

func (p StaticLoginCredentialProvider) LoginCredentials() (LoginCredentials, error) {
	return LoginCredentials{
		Username: p.Username,
		Password: p.Password,
		Auth:     base64.StdEncoding.EncodeToString([]byte(p.Username + ":" + p.Password)),
	}, nil
}
