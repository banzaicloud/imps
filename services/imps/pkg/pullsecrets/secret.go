package pullsecrets

import corev1 "k8s.io/api/core/v1"

func NewBasicAuthSecret(secretNamespace, secretName, registry, user, password string) corev1.Secret {
	
}
