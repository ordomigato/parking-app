package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
)

// check if user is able to edit workspace
func WorkspaceAdmin(c *fiber.Ctx) error {
	wsID, err := uuid.Parse(c.Params("wsID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}

	clientID := c.Locals("client").(models.ClientResponse).ClientID

	clientWS := models.ClientWorkspace{}

	if err := initializers.DB.First(&clientWS, "client_id = ? AND workspace_id = ?", clientID, wsID).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"error_message": fmt.Sprintln("unable to find workspace")})
	}

	if clientWS.RoleID == "Admin" {
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"error_message": fmt.Sprintln("unable to find workspace")})
	}
}
