package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Postgres *sql.DB

func InitPostgres() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Config.PG_HOST,
		Config.PG_PORT,
		Config.PG_USER,
		Config.PG_PASSWORD,
		Config.PG_DATABASE,
	)

	// Open a connection to the database.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	Postgres = db
}
