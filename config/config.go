package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var Cfg config

type config struct {
	Database Database
	JWT      JWTConfig
	Redis    Redis
}

type Database struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SetMaxOpenConns int
}

type Redis struct {
	Addr        string
	DB          int
	Password    string
	DailTimeout time.Duration
}

type SetupResult struct {
	PostgreSQLConnection *sql.DB
	RedisConnection      *redis.Client
	GormConnection       *gorm.DB
}

type JWTConfig struct {
	Secret string
}

func LoadConfig() *SetupResult {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	gdb, err := initGorm(Cfg.Database)
	if err != nil {
		panic(fmt.Sprintf("error at connecting to postgresql with gorm database. err: %v, conntion info :  %+v ", err, Cfg.Database))
	}

	pgdb, err := initPostgreSQL(Cfg.Database)
	if err != nil {
		panic(fmt.Sprintf("error at connecting to postgresql database. err: %v, conntion info :  %+v ", err, Cfg.Database))
	}

	rdb, err := initRedis(Cfg.Redis)
	if err != nil {
		panic(fmt.Sprintf("error at connecting to redis database. err: %v, conntion info :  %+v ", err, Cfg.Redis))
	}

	return &SetupResult{
		PostgreSQLConnection: pgdb,
		RedisConnection:      rdb,
		GormConnection:       gdb,
	}
}
