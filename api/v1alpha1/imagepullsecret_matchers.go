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
	"emperror.dev/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (s ObjectSelectorConfiguration) Matches(meta metav1.ObjectMeta) (bool, error) {
	for _, labelSelector := range s.Labels {
		labelSelector := labelSelector
		matcher, err := metav1.LabelSelectorAsSelector(&labelSelector)
		if err != nil {
			return false, err
		}
		if matcher.Matches(labels.Set(meta.Labels)) {
			return true, nil
		}
	}

	for _, annotationSelector := range s.Annotations {
		labelSelectorFromAnnotations := metav1.LabelSelector{
			MatchExpressions: annotationSelector.MatchExpressions,
			MatchLabels:      annotationSelector.MatchAnnotations,
		}

		matcher, err := metav1.LabelSelectorAsSelector(&labelSelectorFromAnnotations)
		if err != nil {
			return false, err
		}
		if matcher.Matches(labels.Set(meta.Annotations)) {
			return true, nil
		}
	}

	return false, nil
}

func (s ObjectSelectorConfiguration) IsEmpty() bool {
	return len(s.Annotations) == 0 && len(s.Labels) == 0
}

func (i ImagePullSecret) MatchesNamespace(ns *corev1.Namespace) (bool, error) {
	for _, name := range i.Spec.Target.Namespaces.Names {
		if ns.Name == name {
			return true, nil
		}
	}

	match, err := i.Spec.Target.Namespaces.Matches(ns.ObjectMeta)
	if err != nil {
		return false, err
	}

	return match, nil
}

func (i ImagePullSecret) MatchesPod(pod *corev1.Pod) (bool, error) {
	return i.Spec.Target.NamespacesWithPods.Matches(pod.ObjectMeta)
}

func (i ImagePullSecret) SplitNamespacesByMatch(allNs corev1.NamespaceList) ([]corev1.Namespace, []corev1.Namespace, error) {
	match := []corev1.Namespace{}
	nonMatch := []corev1.Namespace{}
	for _, ns := range allNs.Items {
		itemMatches, err := i.MatchesNamespace(ns.DeepCopy())
		if err != nil {
			return nil, nil, errors.WrapWithDetails(err, "cannot filter namespaces", map[string]interface{}{
				"ns":   ns,
				"imps": i,
			})
		}
		if itemMatches {
			match = append(match, ns)
		} else {
			nonMatch = append(nonMatch, ns)
		}
	}
	return match, nonMatch, nil
}
