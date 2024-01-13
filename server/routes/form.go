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
	wpid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}
	fmt.Println(wpid)

	forms := []models.Form{}

	if err := initializers.DB.Where("workspace_id = ?", wpid).Find(&forms).Error; err != nil {
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
	if err := initializers.DB.Find(&form, formId).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find form: %v", err)})
	}
	return c.JSON(form)
}

func CreateForm(c *fiber.Ctx) error {
	wpid, err := uuid.Parse(c.Params("id"))
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
		WorkspaceID:               wpid,
		Name:                      payload.Name,
		Path:                      payload.Path,
		SubmissionConstraintType:  payload.SubmissionConstraintType,
		SubmissionConstraintLimit: payload.SubmissionConstraintLimit,
		CreatedAt:                 now,
		UpdatedAt:                 now,
	}

	if err := initializers.DB.Create(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("unable to find forms: %v", err)})
	}
	return c.JSON(form)
}

func UpdateForm(c *fiber.Ctx) error {
	formId, err := uuid.Parse(c.Params("formId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("id is not a uuid: %v", err)})
	}
	payload := new(models.FormUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_message": err.Error(),
		})
	}

	now := time.Now()

	form := models.Form{
		Name:                      payload.Name,
		Path:                      payload.Path,
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

	// TODO confirm user is admin of this form

	formId, err := uuid.Parse(c.Params("formId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("form_id is not a uuid: %v", err)})
	}

	if err := initializers.DB.Delete(&models.Form{}, formId).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error_message": fmt.Sprintf("Failed to delete: %v", err)})
	}

	return c.SendStatus(http.StatusNoContent)
}
