package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

func (r *Repository) RegisterClient(context *fiber.Ctx) {
	var payload *models.ClientRegisterRequest

	err := context.BodyParser(&payload)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": err.Error()})
		return
	}

	now := time.Now()

	newClient := models.Client{
		Username:  payload.Username,
		Password:  hashedPassword,
		Verified:  false,
		CreatedAt: now,
		UpdatedAt: now,
		LastLogin: now,
	}

	err = r.DB.Create(&newClient).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not register account"})
		return
	}
}
