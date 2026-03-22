package sqlite3

import (
	"database/sql"

	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/db-client/client"
	"github.com/Compogo/db-client/driver"
	logger2 "github.com/Compogo/db-client/logger"
	_ "github.com/mattn/go-sqlite3"
)

// SQLite is the driver identifier for SQLite3 databases.
// It implements the driver.Driver interface and is used for
// registration in db-client, db-migrator, and db-sql-generator.
const SQLite driver.Driver = "sqlite3"

// Client is the interface for SQLite3 database operations.
// It embeds the standard client.Client interface.
type Client client.Client

// sqliteClient wraps *sql.DB and implements the Client interface.
// It provides methods for SQL operations and driver identification.
type sqliteClient struct {
	*sql.DB
}

// NewSQLite creates a new SQLite3 database client.
// It opens a connection using the provided configuration and
// wraps it with a logger decorator for automatic query logging.
//
// The logger is used to log all SQL queries at DEBUG level.
// Returns an error if the connection fails.
func NewSQLite(logger logger.Logger, config *Config) (Client, error) {
	conn, err := sql.Open(string(SQLite), config.DSN)
	if err != nil {
		return nil, err
	}

	return logger2.NewLogger(
		&sqliteClient{conn},
		logger,
	), nil
}

func (m *sqliteClient) SQL() *sql.DB {
	return m.DB
}

func (m *sqliteClient) Driver() driver.Driver {
	return SQLite
}
