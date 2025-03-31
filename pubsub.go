package config

// Pubsub ..
type Pubsub struct {
	ADC       string   `json:"adc" env:"GOOGLE_ADC" default:".cfg/adc.json"`
	ProjectID string   `json:"project_id" env:"PUBSUB_PROJECT_ID"`
	SvcID     string   `json:"svc_id" env:"PUBSUB_SVC_ID"`
	EnvID     string   `json:"env_id" env:"ENV_ID"`
	TopicIDs  []string `json:"topic_ids" env:"PUBSUB_TOPIC_IDS"`
}
