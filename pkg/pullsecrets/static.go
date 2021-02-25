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
