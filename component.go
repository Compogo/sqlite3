package sqlite3

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
)

// Component is a ready-to-use Compogo component that provides a SQLite3 database client.
// It automatically:
//   - Registers Config and Client in the DI container
//   - Adds command-line flags for the DSN
//   - Configures the client during Configuration phase
//   - Wraps the connection with automatic logging
//
// Usage:
//
//	compogo.WithComponents(
//	    dbClient.Component,
//	    sqlite3.Component,
//	)
//
// Then inject the client into your components:
//
//	type UserRepository struct {
//	    db sqlite3.Client
//	}
var Component = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			NewConfig,
			NewSQLite,
		)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.DSN, DsnFieldName, DsnDefault, "sqlite dsn string connection")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
}
