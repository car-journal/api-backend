package config

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbOnce sync.Once

const (
	DB_DIALECT = "DB_DIALECT"
	DB_HOST    = "DB_HOST"
	DB_PORT    = "DB_PORT"
	DB_NAME    = "DB_NAME"
	DB_USER    = "DB_USER"
	DB_PASS    = "DB_PASS"
	ASCENDING  = "ASC"
	DESCENDING = "DESC"
	RANDOM     = "random"
	NULLS_LAST = "NULLS LAST"
)

var postgresConfig = map[string]string{
	DB_DIALECT: "postgres",
	DB_HOST:    "localhost",
	DB_PORT:    "5432",
	DB_NAME:    "volleyball_db",
	DB_USER:    "postgres",
	DB_PASS:    "postgres",
}

func ConnectGormPostgres() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Get(DB_HOST),
			Get(DB_PORT),
			Get(DB_USER),
			Get(DB_PASS),
			Get(DB_NAME),
		)
		db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
		if nil != err {
			panic(err)
		}
		log.Printf(`Connected to database postgres (%s) at %s:%s with user "%s"`,
			Get(DB_NAME),
			Get(DB_HOST),
			Get(DB_PORT),
			Get(DB_USER),
		)
	})

	return db
}
