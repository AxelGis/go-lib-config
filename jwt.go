package config

// JWT ..
type JWT struct {
	PublicKey  string `json:"public_key" env:"JWT_PUBLIC_KEY"`
	PrivateKey string `json:"private_key" env:"JWT_PRIVATE_KEY"`
}
