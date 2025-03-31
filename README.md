# go-lib-config

Golang configuration library for managing application settings using environment variables.

## Usage

Below is an example of how to use the `go-lib-config` library:

```go
package svc_config

import (
    "log"
    "reflect"
    "sync"
    "time"

    "github.com/joho/godotenv"
    "github.com/AxelGis/go-lib-config"
)

type Service struct {
    EncryptionKey string `json:"encryption_key" env:"DB_ENCRYPTION_KEY"`
}

type RequestsWorker struct {
    TickerInterval time.Duration `json:"ticker_interval" env:"WORKER_TICKER_INTERVAL" default:"30s"`
    BatchSize      int           `json:"batch_size" env:"WORKER_BATCH_SIZE" default:"50"`
    MaxAttempts    int           `json:"max_attempts" env:"WORKER_MAX_ATTEMPTS" default:"100"`
}

type SrvConfig struct {
    config.Config
    Service
    RequestsWorker
}

var (
    conf     *SrvConfig
    confOnce sync.Once
)

func New() *SrvConfig {
    conf := &SrvConfig{}

    if err := godotenv.Load(".cfg/.env"); err != nil {
        log.Print("No .env file found")
    }

    config.FillStructFromEnv(reflect.ValueOf(conf).Elem())
    return conf
}

// GetConfig returns the singleton instance of the configuration
func GetConfig() *SrvConfig {
    confOnce.Do(func() {
        conf = New()
    })
    return conf
}