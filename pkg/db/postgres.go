package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	PG_Host     string
	PG_Port     string
	PG_User     string
	PG_Password string
	PG_Name     string
	PG_SSLMode  string
}

func NewConnection(config *Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		config.PG_User, config.PG_Password, config.PG_Name, config.PG_Host, config.PG_Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	} else {
		log.Info().Msg("Connection to Postgres successful")
	}

	return db, nil
}
