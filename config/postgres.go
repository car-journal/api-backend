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
	DBDialect  = "DB_DIALECT"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBName     = "DB_NAME"
	DBUser     = "DB_USER"
	DBPass     = "DB_PASS"
	Ascending  = "ASC"
	Descending = "DESC"
	Random     = "random"
	NullsLast  = "NULLS LAST"
)

var postgresConfig = map[string]string{
	DBDialect: "postgres",
	DBHost:    "localhost",
	DBPort:    "5432",
	DBName:    "volleyball_db",
	DBUser:    "postgres",
	DBPass:    "postgres",
}

func ConnectGormPostgres() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Get(DBHost),
			Get(DBPort),
			Get(DBUser),
			Get(DBPass),
			Get(DBName),
		)
		db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
		if nil != err {
			panic(err)
		}
		log.Printf(`Connected to database postgres (%s) at %s:%s with user "%s"`,
			Get(DBName),
			Get(DBHost),
			Get(DBPort),
			Get(DBUser),
		)
	})

	return db
}
