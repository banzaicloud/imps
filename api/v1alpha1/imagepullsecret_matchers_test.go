package v1alpha1

import (
	"testing"

	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestImagePullSecret_SplitNamespacesByMatch(t *testing.T) {
	type args struct {
		allNs corev1.NamespaceList
	}

	testImagePullSecret := ImagePullSecret{
		Spec: ImagePullSecretSpec{
			Target: TargetConfig{
				Namespaces: NamespaceSelectorConfiguration{
					Names: []string{"testNamespace", "testNamespace2"},
				},
			},
		},
	}

	tests := []struct {
		name                   string
		imagePullSecret        ImagePullSecret
		args                   args
		expectedMatchingNs     []corev1.Namespace
		expectedNonMatichingNs []corev1.Namespace
	}{
		{
			name:            "basic functionality check",
			imagePullSecret: testImagePullSecret,
			args: args{
				allNs: corev1.NamespaceList{
					Items: []corev1.Namespace{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace",
							},
						},
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace2",
							},
						},
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace3",
							},
						},
					},
				},
			},
			expectedMatchingNs: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace2",
					},
				},
			},
			expectedNonMatichingNs: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace3",
					},
				},
			},
		},
		{
			name:            "no matching namespaces",
			imagePullSecret: testImagePullSecret,
			args: args{
				allNs: corev1.NamespaceList{
					Items: []corev1.Namespace{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace3",
							},
						},
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace4",
							},
						},
					},
				},
			},
			expectedMatchingNs: []corev1.Namespace{},
			expectedNonMatichingNs: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace3",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace4",
					},
				},
			},
		},
		{
			name:            "all namespaces match",
			imagePullSecret: testImagePullSecret,
			args: args{
				allNs: corev1.NamespaceList{
					Items: []corev1.Namespace{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace",
							},
						},
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "testNamespace2",
							},
						},
					},
				},
			},
			expectedMatchingNs: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace2",
					},
				},
			},
			expectedNonMatichingNs: []corev1.Namespace{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchingNs, nonMatichingNs, err := tt.imagePullSecret.SplitNamespacesByMatch(tt.args.allNs)

			assert.DeepEqual(t, tt.expectedMatchingNs, matchingNs)
			assert.DeepEqual(t, tt.expectedNonMatichingNs, nonMatichingNs)
			assert.NilError(t, err)
		})
	}
}

func TestImagePullSecret_MatchesPod(t *testing.T) {
	type args struct {
		pod corev1.Pod
	}

	testImagePullSecret := ImagePullSecret{
		Spec: ImagePullSecretSpec{
			Target: TargetConfig{
				NamespacesWithPods: ObjectSelectorConfiguration{
					Labels: []metav1.LabelSelector{
						{
							MatchLabels: map[string]string{
								"testLabel": "true",
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name            string
		imagePullSecret ImagePullSecret
		args            args
		want            bool
	}{
		{
			name:            "basic functionality check",
			imagePullSecret: testImagePullSecret,
			args: args{
				pod: corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"testLabel":  "true",
							"testLabel2": "true",
						},
					},
				},
			},
			want: true,
		},
		{
			name:            "label value different",
			imagePullSecret: testImagePullSecret,
			args: args{
				pod: corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"testLabel":  "false",
							"testLabel2": "true",
						},
					},
				},
			},
			want: false,
		},
		{
			name:            "different label only",
			imagePullSecret: testImagePullSecret,
			args: args{
				pod: corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"testLabel2": "true",
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := tt.imagePullSecret.MatchesPod(&tt.args.pod)

			assert.Equal(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}

func TestImagePullSecret_MatchesNamespace(t *testing.T) {
	type args struct {
		ns corev1.Namespace
	}

	testImagePullSecret := ImagePullSecret{
		Spec: ImagePullSecretSpec{
			Target: TargetConfig{
				Namespaces: NamespaceSelectorConfiguration{
					Names: []string{"testNamespace"},
				},
			},
		},
	}

	tests := []struct {
		name            string
		imagePullSecret ImagePullSecret
		args            args
		want            bool
	}{
		{
			name:            "namespace matches",
			imagePullSecret: testImagePullSecret,
			args: args{
				ns: corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace",
					},
				},
			},
			want: true,
		},
		{
			name:            "namespace doesn't match",
			imagePullSecret: testImagePullSecret,
			args: args{
				ns: corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testNamespace2",
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := tt.imagePullSecret.MatchesNamespace(&tt.args.ns)

			assert.Equal(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}

func TestObjectSelectorConfiguration_IsEmpty(t *testing.T) {

	testObjectSelectorConfiguration := ObjectSelectorConfiguration{
		Labels: []metav1.LabelSelector{
			{
				MatchLabels: map[string]string{
					"testLabel": "true",
				},
			},
		},
		Annotations: []AnnotationSelector{
			{
				MatchAnnotations: map[string]string{
					"testAnnotation": "true",
				},
			},
		},
	}

	tests := []struct {
		name                        string
		objectSelectorConfiguration ObjectSelectorConfiguration
		want                        bool
	}{
		{
			name:                        "object selector config isn't empty",
			objectSelectorConfiguration: testObjectSelectorConfiguration,
			want:                        false,
		},
		{
			name:                        "object selector config is empty",
			objectSelectorConfiguration: ObjectSelectorConfiguration{},
			want:                        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.objectSelectorConfiguration.IsEmpty()

			assert.Equal(t, tt.want, found)
		})
	}
}

func TestObjectSelectorConfiguration_Matches(t *testing.T) {
	type args struct {
		meta metav1.ObjectMeta
	}

	testObjectSelectorConfiguration := ObjectSelectorConfiguration{
		Labels: []metav1.LabelSelector{
			{
				MatchLabels: map[string]string{
					"testLabel": "true",
				},
			},
		},
		Annotations: []AnnotationSelector{
			{
				MatchAnnotations: map[string]string{
					"testAnnotation": "true",
				},
			},
		},
	}

	tests := []struct {
		name                        string
		objectSelectorConfiguration ObjectSelectorConfiguration
		args                        args
		want                        bool
	}{
		{
			name:                        "label match",
			objectSelectorConfiguration: testObjectSelectorConfiguration,
			args: args{
				meta: metav1.ObjectMeta{
					Labels: map[string]string{
						"testLabel":  "true",
						"testLabel2": "true",
					},
				},
			},
			want: true,
		},
		{
			name:                        "annotation match",
			objectSelectorConfiguration: testObjectSelectorConfiguration,
			args: args{
				meta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"testAnnotation": "true",
					},
				},
			},
			want: true,
		},
		{
			name:                        "no matches",
			objectSelectorConfiguration: testObjectSelectorConfiguration,
			args: args{
				meta: metav1.ObjectMeta{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := tt.objectSelectorConfiguration.Matches(tt.args.meta)

			assert.Equal(t, tt.want, found)
			assert.NilError(t, err)
		})
	}
}
