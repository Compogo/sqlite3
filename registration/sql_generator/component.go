package sql_generator

import (
	"github.com/Compogo/compogo/component"
	sqlGenerator "github.com/Compogo/db-sql-generator"
	"github.com/Compogo/sqlite3"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

// Component is a Compogo component that registers the SQLite3 dialect
// with the SQL generator. It has no runtime behavior, only init-time registration.
//
// Add this component if you need SQL generation for SQLite3:
//
//	compogo.WithComponents(sqlite3.Component, sql_generator.Component)
var Component = &component.Component{
	Dependencies: component.Components{
		sqlGenerator.Component,
		sqlite3.Component,
	},
}

func init() {
	sqlGenerator.Registration(sqlite3.SQLite, "sqlite3")
}
