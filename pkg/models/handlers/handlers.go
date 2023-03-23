package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *Repository) CreateWallet(context *fiber.Ctx) error {
	wallet := Wallet{}

	err := context.BodyParser(&wallet)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&wallet).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create a wallet"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "wallet has been added"})
	return nil
}
