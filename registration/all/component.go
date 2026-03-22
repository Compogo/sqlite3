package all

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	dbClient "github.com/Compogo/db-client"
	dbMigrator "github.com/Compogo/db-migrator"
	dbSqlGenerator "github.com/Compogo/db-sql-generator"
	"github.com/Compogo/sqlite3"
	"github.com/Compogo/sqlite3/registration/manager"
	"github.com/Compogo/sqlite3/registration/migrator"
	sqlGenerator "github.com/Compogo/sqlite3/registration/sql_generator"
)

// Component is a composite Compogo component that provides everything needed
// for SQLite3 database integration:
//   - Client: ready-to-use database connection
//   - Migrator: schema migration support
//   - SQL Generator: query builder with SQLite3 dialect
//   - Manager: unified db-client interface
//
// Usage:
//
//	compogo.WithComponents(sqlite3.AllComponent)
//
// This automatically configures the migrator and sql-generator to use the
// SQLite3 driver when the client is registered.
var Component = &component.Component{
	Dependencies: component.Components{
		manager.Component,
		migrator.Component,
		sqlGenerator.Component,
		sqlite3.Component,
	},
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(managerCfg dbClient.Config, migratorCfg *dbMigrator.Config, generatorCfg *dbSqlGenerator.Config) {
			if managerCfg.Driver != sqlite3.SQLite {
				return
			}

			migratorCfg.Driver = sqlite3.SQLite
			generatorCfg.Driver = sqlite3.SQLite
		})
	}),
}
