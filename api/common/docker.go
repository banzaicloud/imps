package common

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Auth     string `json:"auth"` // base64 encoded username:password
}

// DockerRegistryConfig represents a docker compliant image pull secret json file
type DockerRegistryConfig struct {
	Auths map[string]LoginCredentials `json:"auths"`
}
