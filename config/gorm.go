package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func initGorm(database Database) (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
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

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
