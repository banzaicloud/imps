// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
