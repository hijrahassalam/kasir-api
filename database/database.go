package database

import (
	"database/sql"
	"log"
	"strings"

	// postgres driver - using pgx
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Add simple protocol mode to avoid prepared statement cache issues
	if !strings.Contains(connectionString, "default_query_exec_mode") {
		if strings.Contains(connectionString, "?") {
			connectionString += "&default_query_exec_mode=simple_protocol"
		} else {
			connectionString += "?default_query_exec_mode=simple_protocol"
		}
	}

	// Open database
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}