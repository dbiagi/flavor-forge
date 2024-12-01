package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

const (
	MaxIdleConnections = 20
	MaxOpenConnections = 200
)

type Database struct {
	conn *sql.DB
}

type IDatabase interface {
	Query() (*sql.Rows, error)
	Ping() error
}

func (db *Database) Query(query string) (*sql.Rows, error) {
	return db.conn.Query(query)
}

func (db *Database) Ping() error {
	return db.conn.Ping()
}

func ConnectToDatabase() (*Database, error) {
	conn, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		slog.Error("Error connecting to database", err)
		return nil, err
	}

	conn.SetMaxIdleConns(MaxIdleConnections)
	conn.SetMaxOpenConns(MaxOpenConnections)

	pingErr := conn.Ping()
	if pingErr != nil {
		slog.Error("Error pinging database", slog.String("error", pingErr.Error()))
		panic(pingErr)
	}

	return &Database{conn: conn}, nil
}
