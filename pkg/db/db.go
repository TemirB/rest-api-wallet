package db

import (
	"database/sql"
	"log"

	"github.com/TemirB/rest-api-wallet/internal/config"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	connStr := config.SetDBConfig()
	DB, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	return DB
}
