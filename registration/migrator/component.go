package migrator

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	migrator "github.com/Compogo/db-migrator"
	"github.com/Compogo/sqlite3"
	"github.com/golang-migrate/migrate/v4/database"
	migratorSQLite "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

// Component is a Compogo component that registers the SQLite3 migration driver
// with the migrator. It has no runtime behavior, only init-time registration.
//
// Add this component if you need database migrations for SQLite3:
//
//	compogo.WithComponents(sqlite3.Component, migrator.Component)
var Component = &component.Component{
	Dependencies: component.Components{
		migrator.Component,
		sqlite3.Component,
	},
}

func init() {
	migrator.Registration(sqlite3.SQLite, func(container container.Container) (database.Driver, error) {
		var c sqlite3.Client

		err := container.Invoke(func(client sqlite3.Client) {
			c = client
		})
		if err != nil {
			return nil, err
		}

		return migratorSQLite.WithInstance(c.SQL(), &migratorSQLite.Config{})
	})
}
