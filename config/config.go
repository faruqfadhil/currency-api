package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	CurrencyDBHost     string `envconfig:"CURRENCY_DB_HOST" default:"localhost"`
	CurrencyDBPort     string `envconfig:"CURRENCY_DB_PORT" default:"3306"`
	CurrencyDBUsername string `envconfig:"CURRENCY_DB_USERNAME" default:"learn"`
	CurrencyDBPassword string `envconfig:"CURRENCY_DB_PASSWORD" default:"ruangguru123"`
	CurrencyDBName     string `envconfig:"CURRENCY_DB_NAME" default:"classicmodels"`
	Port               string `envconfig:"PORT" default:"8081"`
	GinMode            string `envconfig:"GIN_MODE" default:""`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
