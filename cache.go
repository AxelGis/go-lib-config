package config

type Cache struct {
	Addr  string `json:"addr" env:"CACHE_ADDR"`
	Pass  string `json:"pass" env:"CACHE_PASS"`
	DBNum int    `json:"db_num" env:"CACHE_DB_NUM"`
}
