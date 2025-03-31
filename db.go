package config

// DB ..
type DB struct {
	DSN                string `json:"db" env:"DSN"`
	AutoMigrate        bool   `json:"auto_migrate" env:"DB_AUTO_MIGRATE" default:"false"`
	MaxIdleConnections int    `json:"max_idle_connections" env:"DB_MAX_IDLE_CONNECTIONS" default:"10"`
	MaxOpenConnections int    `json:"max_open_connections" env:"DB_MAX_OPEN_CONNECTIONS" default:"50"`
}
