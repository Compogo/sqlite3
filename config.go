package sqlite3

import "github.com/Compogo/compogo/configurator"

const (
	DsnFieldName = "db.sqlite.dsn"

	DsnDefault = "file:/tmp/test.db?cache=shared"
)

type Config struct {
	DSN string
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if config.DSN == "" || config.DSN == DsnDefault {
		configurator.SetDefault(DsnFieldName, DsnDefault)
		config.DSN = configurator.GetString(DsnFieldName)
	}

	return config
}
