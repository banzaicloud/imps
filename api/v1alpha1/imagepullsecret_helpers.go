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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/banzaicloud/operator-tools/pkg/utils"
)

func (i *ImagePullSecret) GetOwnerReferenceForOwnedObject() metav1.OwnerReference {
	return metav1.OwnerReference{
		APIVersion: i.APIVersion,
		Kind:       i.Kind,
		Name:       i.Name,
		UID:        i.UID,
		Controller: utils.BoolPointer(false),
	}
}

func (r RegistryConfig) CredentialsAsNamespacedNameList() []types.NamespacedName {
	list := make([]types.NamespacedName, len(r.Credentials))
	for idx, cred := range r.Credentials {
		list[idx] = types.NamespacedName{
			Namespace: cred.Namespace,
			Name:      cred.Name,
		}
	}

	return list
}
