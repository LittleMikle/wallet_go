package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Wallet struct {
	ID      uint    `json:"id"`
	Owner   *string `json:"owner"`
	Balance *int    `json:"balance"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_wallet", r.CreateWallet)
	//api.Delete("/delete_wallet/:id", r.DeleteWallet)
	//api.Get("/get_wallet/:id", r.GetWalletByID)
	//api.Get("/books", r.GetWallets)
}
