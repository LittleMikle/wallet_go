package config

import (
	"fmt"
	"github.com/LittleMikle/wallet_go/pkg/db"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func CreateConfig() (Conf *db.Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	} else {
		log.Info().Msg("Config created successfully")
	}

	Conf = &db.Config{
		PG_Host:     os.Getenv("POSTGRES_HOST"),
		PG_Port:     os.Getenv("POSTGRES_PORT"),
		PG_User:     os.Getenv("POSTGRES_USER"),
		PG_Password: os.Getenv("POSTGRES_PASSWORD"),
		PG_Name:     os.Getenv("POSTGRES_NAME"),
		PG_SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	return Conf, nil
}
