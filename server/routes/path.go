package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
)

func UpdateFormPath(c *fiber.Ctx) error {
	formId, err := uuid.Parse(c.Params("formId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("form id is not a uuid: %v", err)})
	}

	wsID, err := uuid.Parse(c.Params("wsID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("workspace id is not a uuid: %v", err)})
	}

	payload := new(models.PathUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	p := payload.Path

	path := models.Path{
		Path: p,
	}
	if err := initializers.DB.Where("form_id = ? AND workspace_id = ?", formId, wsID).Updates(&path).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to update path: %v", err)})
	}

	return c.SendStatus(http.StatusNoContent)
}
