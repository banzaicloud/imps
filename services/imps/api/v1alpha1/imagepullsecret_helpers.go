// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

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
