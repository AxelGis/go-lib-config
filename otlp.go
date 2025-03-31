package config

type Otlp struct {
	Host                       string `json:"host" env:"OTLP_HOST"`
	IsEnabled                  bool   `json:"is_enabled" env:"OTLP_ENABLED"`
	IncludeVariables           bool   `json:"include_variables" env:"OTLP_INCLUDE_VARIABLES" default:"false"`
	ShouldCreateSpanFromFields bool   `json:"should_create_span_from_fields" env:"OTLP_SHOULD_CREATE_SPAN_FROM_FIELDS" default:"false"`
}
