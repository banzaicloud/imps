// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/labels"
)

type LabelSet labels.Set

func (s LabelSet) DeepCopy() LabelSet {
	newSet := LabelSet{}
	for k, v := range s {
		newSet[k] = v
	}
	return newSet
}
