package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Postgres *sql.DB

func InitPostgres() error {
	var sourceStr string
	if Config.PROD {
		sourceStr = "host=%s port=%s user=%s password=%s dbname=%s sslmode=require"
	} else {
		sourceStr = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	}

	psqlInfo := fmt.Sprintf(
		sourceStr,
		Config.PG_HOST,
		Config.PG_PORT,
		Config.PG_USER,
		Config.PG_PASSWORD,
		Config.PG_DATABASE,
	)

	// Open a connection to the database.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	Postgres = db

	return nil
}
