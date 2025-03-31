package config

// Auth ..
type Auth struct {
	Audience     string `json:"audience" env:"AUTH_AUDIENCE"`
	CacheEnabled bool   `json:"cache_enabled" env:"AUTH_CACHE_ENABLED" default:"true"`
}
