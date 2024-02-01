package routes

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
	"gorm.io/gorm"
)

func CreateWorkspace(c *fiber.Ctx) error {
	payload := new(models.WorkspaceCreateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	now := time.Now()

	newWorkspace := models.Workspace{
		Name:      payload.Name,
		Path:      payload.Path,
		UpdatedAt: now,
		CreatedAt: now,
	}

	if err := initializers.DB.Create(&newWorkspace).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("Unable to create workspace. Path has already been claimed."))
		}
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to create workspace"))
	}

	newclientWorkspace := models.ClientWorkspace{
		WorkspaceID: newWorkspace.WorkspaceID,
		ClientID:    c.Locals("client").(models.ClientResponse).ClientID,
		RoleID:      "Admin",
	}
	if err := initializers.DB.Create(&newclientWorkspace).Error; err != nil {
		initializers.DB.Delete(&newWorkspace)
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Unable to create client workspace relationship"))
	}

	return c.JSON(newWorkspace)
}

func GetWorkspaces(c *fiber.Ctx) error {
	clientId := c.Locals("client").(models.ClientResponse).ClientID

	fmt.Println(clientId)

	// TODO - clean this up with a single query
	clientWorkspaces := []models.ClientWorkspace{}
	if err := initializers.DB.Find(&clientWorkspaces, "client_id = ?", clientId).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Unable to find client workspace relationship"))
	}

	fmt.Println(clientWorkspaces)
	workspaces := []models.Workspace{}
	workspaceIds := []uuid.UUID{}

	for _, cw := range clientWorkspaces {
		workspaceIds = append(workspaceIds, cw.WorkspaceID)
	}

	if err := initializers.DB.Where("workspace_id = ?", workspaceIds).Find(&workspaces).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find workspaces"))
	}
	fmt.Println(workspaceIds)
	fmt.Println(workspaces)

	return c.JSON(workspaces)
}

func UpdateWorkspace(c *fiber.Ctx) error {
	wsid, err := uuid.Parse(c.Params("wsid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", wsid)))
	}

	payload := new(models.WorkspaceUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	now := time.Now()

	updatedWorkspace := models.Workspace{
		Name: payload.Name,
		// Path:      payload.Path,
		UpdatedAt: now,
	}
	if err := initializers.DB.Model(&models.Workspace{}).Where("workspace_id = ?", wsid).Updates(updatedWorkspace).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Failed to delete"))
	}

	// TODO: Update all forms paths if path has changed

	return c.SendStatus(http.StatusNoContent)
}

func DeleteWorkspace(c *fiber.Ctx) error {

	// TODO confirm user is admin of this workspace

	wsid, err := uuid.Parse(c.Params("wsid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", wsid)))
	}

	// DELETE WORKSPACE
	if err := initializers.DB.Delete(&models.Workspace{}, wsid).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("Failed to delete workspace. Please ensure all forms are deleted first."))
		}
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(("Failed to delete workspace")))
	}

	// DELETE CLIENT WORKSPACE RELATIONSHIP
	if err := initializers.DB.Delete(&models.ClientWorkspace{}, "workspace_id = ?", wsid).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(("Failed to delete client workspace relation")))
	}
	return c.SendStatus(http.StatusNoContent)
}
