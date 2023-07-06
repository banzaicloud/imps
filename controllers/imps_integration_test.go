package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	apiv1 "github.com/banzaicloud/imps/api/v1alpha1"
)

// +kubebuilder:docs-gen:collapse=Imports

var _ = Describe("IMPS controller", func() {
	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		ImpsName              = "test-imps"
		ImpsNamespace         = "source-ns"
		SourceSecretName      = "source-secret-1"
		TargetSecretName      = "target-secret"
		TargetSecretNamespace = "target-ns"

		timeout  = time.Second * 20
		interval = time.Millisecond * 250
	)

	Describe("When creating an IMPS resource", Ordered, func() {
		It("Secret should be created in the annotated namespace", func() {
			ctx := context.Background()

			sourceNamespace := &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        ImpsNamespace,
					Annotations: map[string]string{"test.io/test-annotation": "test-value"},
				},
			}
			Expect(k8sClient.Create(ctx, sourceNamespace)).Should(Succeed())

			targetNamespace := &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        TargetSecretNamespace,
					Annotations: map[string]string{"test.io/test-annotation": "test-value"},
				},
			}
			Expect(k8sClient.Create(ctx, targetNamespace)).Should(Succeed())

			sourceSecret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      SourceSecretName,
					Namespace: ImpsNamespace,
				},
				StringData: map[string]string{".dockerconfigjson": "{\n  \"auths\": {\n    \"my-registry.example:5000\": {\n      \"username\": \"tiger\",\n      \"password\": \"pass1234\",\n      \"email\": \"tiger@acme.example\",\n      \"auth\": \"dGlnZXI6cGFzczEyMzQ=\"\n    }\n  }\n}"},
				Type:       "kubernetes.io/dockerconfigjson",
			}
			Expect(k8sClient.Create(ctx, sourceSecret)).Should(Succeed())

			pullSecret := &apiv1.ImagePullSecret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      ImpsName,
					Namespace: ImpsNamespace,
				},
				Spec: apiv1.ImagePullSecretSpec{
					Target: apiv1.TargetConfig{
						Secret: apiv1.TargetSecretConfig{
							Name: TargetSecretName,
						},
						Namespaces: apiv1.NamespaceSelectorConfiguration{
							Names: []string{TargetSecretNamespace},
						},
					},
					Registry: apiv1.RegistryConfig{
						Credentials: []apiv1.NamespacedName{
							{
								Name:      SourceSecretName,
								Namespace: ImpsNamespace,
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, pullSecret)).Should(Succeed())

			secretLookupKey := types.NamespacedName{Name: TargetSecretName, Namespace: TargetSecretNamespace}
			createdSecret := &v1.Secret{}

			By("TODO")
			Eventually(func() bool {
				err := k8sClient.Get(ctx, secretLookupKey, createdSecret)
				if err != nil { //nolint:gosimple
					return false
				}

				return true
			}, timeout, interval).Should(BeTrue())
		})
	})
})
