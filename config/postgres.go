package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/car-journal/api-backend/lib/logger"
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
	DBSSLMode  = "DB_SSL_MODE"
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
	DBSSLMode: "disable",
}

func ConnectGormPostgres() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			Get(DBHost),
			Get(DBPort),
			Get(DBUser),
			Get(DBPass),
			Get(DBName),
			Get(DBSSLMode),
		)
		db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
		if nil != err {
			logger.LoggerInterface.Log(err.Error())
			panic(err)
		}
		log.Printf(`Connected to database postgres (%s) at %s:%s with user "%s" and sslmode "%s"`,
			Get(DBName),
			Get(DBHost),
			Get(DBPort),
			Get(DBUser),
			Get(DBSSLMode),
		)
	})

	return db
}
