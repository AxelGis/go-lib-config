package config

type Prometheus struct {
	Port string `json:"prometheus_port" env:"PROMETHEUS_PORT" default:"9090"`
}
