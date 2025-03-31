package config

// Server ..
type Server struct {
	GRPCPort  uint32 `json:"grpc_port" env:"GRPC_PORT" default:"8082"`
	HTTPPort  uint32 `json:"http_port" env:"HTTP_PORT" default:"8080"`
	DebugPort uint32 `json:"debug_port" env:"DEBUG_PORT" default:"8084"`
}
