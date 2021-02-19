// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

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
