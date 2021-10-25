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

import "emperror.dev/errors"

const (
	SourceSecretStatus = "Ok"
)

type ErrorsPerSecret map[string]error

func NewErrorsPerSecret() ErrorsPerSecret {
	return ErrorsPerSecret{}
}

func (e ErrorsPerSecret) AddSecret(name string) {
	e[name] = nil
}

func (e ErrorsPerSecret) SetSecretError(name string, err error) {
	e[name] = err
}

func (e ErrorsPerSecret) AsStatus() map[string]string {
	status := map[string]string{}
	for secret, err := range e {
		if err == nil {
			status[secret] = SourceSecretStatus
		} else {
			status[secret] = err.Error()
		}
	}
	return status
}

func (e ErrorsPerSecret) FailedSecrets() []string {
	invalidSecrets := []string{}

	for secret, err := range e {
		if err != nil {
			invalidSecrets = append(invalidSecrets, secret)
		}
	}
	return invalidSecrets
}

func (e *ErrorsPerSecret) AsError() error {
	invalidSecrets := e.FailedSecrets()

	if len(invalidSecrets) == 0 {
		return nil
	}

	return errors.NewWithDetails("some source secrets failed to render", "failed_secrets", invalidSecrets)
}
