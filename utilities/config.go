package utilities

import (
	"github.com/joegasewicz/pg-conf"
)

type AppConfig struct {
	Port int
	*pg_conf.PostgresConfig
}

func NewConf() *AppConfig {
	cfg := AppConfig{
		Port: 7001,
		PostgresConfig: &pg_conf.PostgresConfig{
			PGPort:     "5431",
			PGDatabase: "identity_db",
			PGUser:     "admin",
			PGPassword: "admin",
		}}

	pg_conf.Update(cfg.PostgresConfig)
	return &cfg
}

var Config = NewConf()
