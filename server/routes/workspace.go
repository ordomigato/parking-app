package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ordomigato/parking-app/models"
)

func CreateWorkspace(c *fiber.Ctx) error {
	payload := new(models.WorkspaceCreateRequest)
	// Extract the credentials from the request body
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": payload,
	})
}
