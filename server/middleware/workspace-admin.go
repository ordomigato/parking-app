package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

// check if user is able to edit workspace
func WorkspaceAdmin(c *fiber.Ctx) error {
	wsID, err := uuid.Parse(c.Params("wsID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", wsID)))
	}

	clientID := c.Locals("client").(models.ClientResponse).ClientID

	clientWS := models.ClientWorkspace{}

	if err := initializers.DB.First(&clientWS, "client_id = ? AND workspace_id = ?", clientID, wsID).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(
			utils.GenerateServerErrorResponse("unable to find workspace"))
	}

	if clientWS.RoleID == "Admin" {
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).JSON(
			utils.GenerateServerErrorResponse(("unable to find workspace")))
	}
}
