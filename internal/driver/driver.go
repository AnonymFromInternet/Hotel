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

var dbConn = &DB{}

const maxOpenDbBConn = 15
const maxIdleDBConn = 15
const maxDBLifeTime = 15 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	db, err := newDB(dsn)
	if err != nil {
		panic(err)
	}

	// Configuring DB
	db.SetMaxOpenConns(maxOpenDbBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)
	// Configuring DB

	dbConn.SQL = db

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// newDB creates new database for the app
func newDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// Testing connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// Testing connection

	return db, nil
}

// testDB tries to ping the database
func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
