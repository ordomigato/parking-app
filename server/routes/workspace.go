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
		Path:      payload.Path,
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

	return c.JSON(newWorkspace)
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

	return c.JSON(workspaces)
}

func UpdateWorkspace(c *fiber.Ctx) error {
	wpid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}

	payload := new(models.WorkspaceUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	now := time.Now()

	updatedWorkspace := models.Workspace{
		Name:      payload.Name,
		Path:      payload.Path,
		UpdatedAt: now,
	}
	if err := initializers.DB.Model(&models.Workspace{}).Where("workspace_id = ?", wpid).Updates(updatedWorkspace).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to delete: %v", err)})
	}

	// TODO: Update all forms paths if path has changed

	return c.SendStatus(http.StatusNoContent)
}

func DeleteWorkspace(c *fiber.Ctx) error {

	// TODO confirm user is admin of this workspace

	wpid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}

	// DELETE ALL FORMS UNDER THIS WORKSPACE
	if err := initializers.DB.Delete(&models.Form{}, "workspace_id = ?", wpid).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to delete forms from workspace: %v", err)})
	}

	// DELETE WORKSPACE
	if err := initializers.DB.Delete(&models.Workspace{}, wpid).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to delete: %v", err)})
	}

	// DELETE CLIENT WORKSPACE RELATIONSHIP
	if err := initializers.DB.Delete(&models.ClientWorkspace{}, "workspace_id = ?", wpid).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to delete workspace: %v", err)})
	}
	return c.SendStatus(http.StatusNoContent)
}
