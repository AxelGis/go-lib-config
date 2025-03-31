package config

type Geo struct {
	GoogleMapsApiKey string `json:"google_maps_api_key" env:"GOOGLE_MAPS_API_KEY"`
	ResultType       string `json:"result_type" env:"GOOGLE_MAPS_RESULT_TYPE"`
}
