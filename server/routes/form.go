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

func GetForms(c *fiber.Ctx) error {
	wsID, err := uuid.Parse(c.Params("wsID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}

	forms := []models.Form{}

	if err := initializers.DB.Where("workspace_id = ?", wsID).Find(&forms).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find forms: %v", err)})
	}

	return c.JSON(forms)
}

func GetForm(c *fiber.Ctx) error {
	formId, err := uuid.Parse(c.Params("formId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}
	form := models.Form{}
	if err := initializers.DB.Model(&models.Form{}).Preload("Path").Where("form_id = ?", formId).First(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find form: %v", err)})
	}
	return c.JSON(form)
}

func CreateForm(c *fiber.Ctx) error {
	wsID, err := uuid.Parse(c.Params("wsID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}

	payload := new(models.FormCreateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	now := time.Now()

	form := models.Form{
		WorkspaceID: wsID,
		Name:        payload.Name,
		// Questions:                 payload.Questions,
		SubmissionConstraintType:  payload.SubmissionConstraintType,
		SubmissionConstraintLimit: payload.SubmissionConstraintLimit,
		CreatedAt:                 now,
		UpdatedAt:                 now,
	}

	if err := initializers.DB.Create(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find forms: %v", err)})
	}

	p := payload.Path

	// update path
	path := models.Path{
		Path:        p,
		WorkspaceID: wsID,
		FormID:      form.FormID,
	}
	if err := initializers.DB.Create(&path).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to create path: %v", err)})
	}
	return c.JSON(form)
}

func UpdateForm(c *fiber.Ctx) error {
	formId, err := uuid.Parse(c.Params("formId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("form id is not a uuid: %v", err)})
	}

	payload := new(models.FormUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	now := time.Now()

	form := models.Form{
		Name: payload.Name,
		// Questions:                 payload.Questions,
		SubmissionConstraintType:  payload.SubmissionConstraintType,
		SubmissionConstraintLimit: payload.SubmissionConstraintLimit,
		UpdatedAt:                 now,
	}

	if err := initializers.DB.Model(&models.Form{}).Where("form_id = ?", formId).Updates(form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to update form: %v", err)})
	}

	return c.JSON(form)
}

func DeleteForm(c *fiber.Ctx) error {
	formId, err := uuid.Parse(c.Params("formId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("form_id is not a uuid: %v", err)})
	}

	wsID, err := uuid.Parse(c.Params("wsID"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("workspace id is not a uuid: %v", err)})
	}

	if err := initializers.DB.Where("form_id = ? AND workspace_id = ?", formId, wsID).Delete(&models.Path{}).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to path: %v", err)})
	}

	if err := initializers.DB.Delete(&models.Form{}, formId).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to delete: %v", err)})
	}

	return c.SendStatus(http.StatusNoContent)
}

func GetFormInfo(c *fiber.Ctx) error {
	wpPath := c.Params("workspacePath")
	formPath := c.Params("formPath")

	p := "/" + wpPath + "/" + formPath

	path := models.Path{}

	if err := initializers.DB.Where("path = ?", p).First(&path).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find forms: %v", err)})
	}

	form := models.Form{}

	if err := initializers.DB.Where("form_id = ?", path.FormID).First(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find forms: %v", err)})
	}

	return c.JSON(&form)
}
