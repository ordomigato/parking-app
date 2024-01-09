package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
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

	now := time.Now()

	newWorkspace := models.Workspace{
		Name:      payload.Name,
		UpdatedAt: now,
		CreatedAt: now,
	}

	if err := initializers.DB.Create(&newWorkspace).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to create workspace: %v", err)})
	}

	// result := initializers.DB.Where("role_name <> ?", "Admin").Find(models.Role{})

	newclientWorkspace := models.ClientWorkspace{
		WorkspaceID: newWorkspace.WorkspaceID,
		ClientID:    c.Locals("client").(models.ClientResponse).ClientID,
		RoleID:      "Admin",
	}
	if err := initializers.DB.Create(&newclientWorkspace).Error; err != nil {
		initializers.DB.Delete(&newWorkspace)
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": "Unable to create client workspace relationship"})
	}

	return c.JSON(fiber.Map{
		"data": newWorkspace,
	})
}

func GetWorkspaces(c *fiber.Ctx) error {
	clientId := c.Locals("client").(models.ClientResponse).ClientID

	// TODO - clean this up with a single query
	clientWorkspaces := []models.ClientWorkspace{}
	if err := initializers.DB.Find(&clientWorkspaces, "client_id = ?", clientId).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": "Unable to find client workspace relationship"})
	}
	workspaces := []models.Workspace{}
	workspaceIds := []uuid.UUID{}

	for _, cw := range clientWorkspaces {
		workspaceIds = append(workspaceIds, cw.WorkspaceID)
	}

	if err := initializers.DB.Find(&workspaces, workspaceIds).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find workspaces: %v", err)})
	}

	return c.JSON(&workspaces)
}
