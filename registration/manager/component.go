package manager

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	dbClient "github.com/Compogo/db-client"
	"github.com/Compogo/db-client/client"
	"github.com/Compogo/sqlite3"
)

// Component is a Compogo component that registers the SQLite3 client
// with the db-client manager. It has no runtime behavior, only init-time registration.
//
// Add this component if you want to use SQLite3 through the unified db-client interface:
//
//	compogo.WithComponents(sqlite3.Component, manager.Component)
var Component = &component.Component{
	Dependencies: component.Components{
		dbClient.Component,
		sqlite3.Component,
	},
}

func init() {
	dbClient.Registration(sqlite3.SQLite, func(container container.Container) (client.Client, error) {
		var c sqlite3.Client

		err := container.Invoke(func(client sqlite3.Client) {
			c = client
		})
		if err != nil {
			return nil, err
		}

		return c, nil
	})
}
