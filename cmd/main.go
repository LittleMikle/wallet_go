package main

import (
	"github.com/LittleMikle/wallet_go/pkg/config"
	"github.com/LittleMikle/wallet_go/pkg/db"
	"github.com/LittleMikle/wallet_go/pkg/models"
	"github.com/LittleMikle/wallet_go/pkg/models/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.CreateConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	database, err := db.NewConnection(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	err = models.MigrateWallets(database)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	//DB = pointer on gorm.DB
	//db = value from our DataBase
	r := handlers.Repository{
		DB: database,
	}
	app := fiber.New()
	//access to struct inside method
	r.SetupRoutes(app)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
