// Copyright © 2021 Cisco Systems
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

import "context"

/*
ErroredCredentialProvider can be used to store a setup error into the config object

	so that at least the providers that are working correctly gets reconciled.
*/
type ErroredCredentialProvider struct {
	Error error
}

func NewErroredCredentialProvider(err error) ErroredCredentialProvider {
	return ErroredCredentialProvider{
		Error: err,
	}
}

func (p ErroredCredentialProvider) LoginCredentials(_ context.Context) ([]LoginCredentialsWithDetails, error) {
	return nil, p.Error
}
