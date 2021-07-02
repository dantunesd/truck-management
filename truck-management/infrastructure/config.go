package infrastructure

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName string `envconfig:"APP_NAME" default:"Truck-Management"`
	DBConn  string `envconfig:"DB_CONN" default:"root:root@tcp(localhost:3306)/truck_management"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	return cfg, envconfig.Process("", cfg)
}
