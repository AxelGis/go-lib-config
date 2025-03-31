package config

// Search server
type Search struct {
	Servers        []string `json:"servers" env:"SEARCH_SERVERS"`
	Index          string   `json:"index" env:"SEARCH_INDEX"`
	Username       string   `json:"username" env:"SEARCH_USERNAME"`
	Password       string   `json:"password" env:"SEARCH_PASSWORD"`
	Insecure       bool     `json:"insecure" env:"SEARCH_INSECURE"`
	TerminateAfter uint64   `json:"terminate_after" env:"SEARCH_TERMINATE_AFTER"`
}
