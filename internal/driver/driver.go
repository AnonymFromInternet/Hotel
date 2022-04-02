package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dataBaseConnection = &DB{}

const maxOpenedDBConnections = 33
const maxIdleDBConnections = 5
const maxDBConnectionLifeTime = 5 * time.Minute

// ConnectSQL returns an DB type object which contains database
func ConnectSQL(dataSourceName string) (*DB, error) {
	connection, err := NewDatabase(dataSourceName)
	if err != nil {
		panic(err)
	}

	// Connection Settings:
	connection.SetMaxOpenConns(maxOpenedDBConnections)
	connection.SetMaxIdleConns(maxIdleDBConnections)
	connection.SetConnMaxLifetime(maxDBConnectionLifeTime)
	// Connection Settings

	dataBaseConnection.SQL = connection

	err = testConnection(connection)
	if err != nil {
		return nil, err
	}

	return dataBaseConnection, nil
}

// testConnection tries to ping the database
func testConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates new connection
func NewDatabase(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = testConnection(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
