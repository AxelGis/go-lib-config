package config

// BigQuery ..
type BigQuery struct {
	DSN     string `json:"dsn" env:"BIGQUERY_DSN"`
	ADCPath string `json:"adc_path" env:"BIGQUERY_ADC_PATH"`
}
