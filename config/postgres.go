package config

import (
	"database/sql"
	"fmt"
	"log"
)

func initPostgreSQL(database Database) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
			database.Port,
			database.Host,
			database.User,
			database.Password,
			database.DBName,
		),
	)
	if err != nil {
		log.Panicln(err)
	}

	db.SetMaxOpenConns(database.SetMaxOpenConns)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
