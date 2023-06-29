package pg

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var connectionString string = "user=postgres password=1234 dbname=fullVagas sslmode=disable host=localhost port=5432"

var instanceDb *sql.DB = nil

func GetConnection() *sql.DB {
	if instanceDb == nil {
		conn, err := sql.Open("postgres", connectionString)
		if err != nil {
			panic("could not connect to the database: " + err.Error())
		}

		if err := conn.Ping(); err != nil {
			panic("could not connect to the database: " + err.Error())
		}

		instanceDb = conn
	}
	return instanceDb
}
