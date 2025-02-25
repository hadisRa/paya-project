package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initGorm(database Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		database.Port,
		database.Host,
		database.User,
		database.Password,
		database.DBName,
	)

	pgdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	db, err := pgdb.DB()
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return pgdb, nil
}
