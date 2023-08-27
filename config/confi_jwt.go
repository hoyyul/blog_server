package config

type Jwt struct {
	SecretKey string `json:"secret" yaml:"secret_key"`
	Expires   int    `json:"expires" yaml:"expires"`
	Issuer    string `json:"issuer" yaml:"issuer"`
}
