package config

// Config common configuration structure
type Config struct {
	AppEnv      string     `json:"app_env" env:"APP_ENV"`
	EnvID       string     `json:"env_id" env:"ENV_ID"`
	LogLevel    string     `json:"log_level" env:"LOG_LEVEL" default:"DEBUG"`
	ServiceName string     `json:"service_name" env:"SERVICE_NAME"`
	Auth        Auth       // auth and roles configs
	BigQuery    BigQuery   // bigquery configs
	Cache       Cache      // redis cache config
	DB          DB         // database config
	Geo         Geo        // google maps configs
	JWT         JWT        // JWT config
	Otlp        Otlp       // otlp configs
	Prometheus  Prometheus // prometheus configs
	Pubsub      Pubsub     // pubsub configs
	Search      Search     // elasticsearch configs
	Server      Server     // server ports
}

type SrvConfig struct {
	Config
}
