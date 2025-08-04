package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vityasyyy/sharedlib/logger"
)

func MustConnect(driver string, dbURL string) *sqlx.DB {
	db, err := sqlx.Connect(driver, dbURL)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to DB")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to ping DB")
	}

	return db
}
