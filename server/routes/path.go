package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

func UpdateFormPath(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("form id is not a uuid"))
	}

	wsid, err := uuid.Parse(c.Params("wsid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("workspace id is not a uuid"))
	}

	payload := new(models.PathUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	p := payload.Path

	path := models.Path{
		Path: p,
	}
	if err := initializers.DB.Where("form_id = ? AND workspace_id = ?", formid, wsid).Updates(&path).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to update path"))
	}

	return c.SendStatus(http.StatusNoContent)
}
