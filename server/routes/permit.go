package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
)

func CreatePermit(c *fiber.Ctx) error {
	payload := new(models.PermitCreateRequest)
	// Extract the credentials from the request body
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	now := time.Now()

	newPermit := models.Permit{
		FirstName:           payload.FirstName,
		LastName:            payload.LastName,
		Email:               payload.Email,
		PrimaryPhone:        payload.PrimaryPhone,
		VPlate:              payload.VPlate,
		VMake:               payload.VMake,
		VModel:              payload.VModel,
		VColor:              payload.VColor,
		Duration:            payload.Duration,
		DurationMeasurement: payload.DurationMeasurement,
		UpdatedAt:           now,
		CreatedAt:           now,
	}

	if err := initializers.DB.Create(&newPermit).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to create permit: %v", err)})
	}

	return c.JSON(newPermit)
}
