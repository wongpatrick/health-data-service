package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func Connectdb() (*sql.DB, error) {
	// in a real setting, I would put this in the kubernetes config file.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)
	db, errdb := sql.Open("postgres", psqlInfo)
	if errdb != nil {
		return nil, errdb
	}
	err := db.Ping()
	return db, err
}
