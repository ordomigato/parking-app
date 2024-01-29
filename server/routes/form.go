package routes

import (
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

func GetForms(c *fiber.Ctx) error {
	wsid, err := uuid.Parse(c.Params("wsid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", wsid)))
	}

	forms := []models.Form{}

	if err := initializers.DB.Where("workspace_id = ?", wsid).Preload("Path").Find(&forms).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("unable to find forms: %v", wsid)))
	}

	return c.JSON(forms)
}

func GetForm(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", err)))
	}
	form := models.Form{}
	if err := initializers.DB.Model(&models.Form{}).Preload("Path").Where("form_id = ?", formid).First(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find form"))
	}
	return c.JSON(form)
}

func CreateForm(c *fiber.Ctx) error {
	wsid, err := uuid.Parse(c.Params("wsid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", err)))
	}

	payload := new(models.FormCreateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	now := time.Now()

	form := models.Form{
		WorkspaceID: wsid,
		Name:        payload.Name,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if payload.CycleData != nil {
		form.CycleData.DurationLimit.Value = payload.CycleData.DurationLimit.Value
		form.CycleData.DurationLimit.Unit = payload.CycleData.DurationLimit.Unit
		form.CycleData.ResetInterval.Value = payload.CycleData.ResetInterval.Value
		form.CycleData.ResetInterval.Unit = payload.CycleData.ResetInterval.Unit
		form.CycleData.ResetInterval.RefDate = payload.CycleData.ResetInterval.RefDate
	}

	if err := initializers.DB.Create(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find forms"))
	}

	p := payload.Path

	// update path
	path := models.Path{
		Path:        p,
		WorkspaceID: wsid,
		FormID:      form.FormID,
	}
	if err := initializers.DB.Create(&path).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to create path"))
	}
	return c.JSON(form)
}

func UpdateForm(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("form id is not a uuid: %v", formid)))
	}

	payload := new(models.FormUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	now := time.Now()

	form := models.Form{
		Name:      payload.Name,
		CycleData: payload.CycleData,
		UpdatedAt: now,
	}

	if err := initializers.DB.Model(&models.Form{}).Where("form_id = ?", formid).Updates(form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to update form"))
	}

	if payload.CycleData == nil {
		if err := initializers.DB.Model(&models.Form{}).Where("form_id = ?", formid).Update("cycle_data", gorm.Expr("NULL")).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("unable to force null cycle data"))
		}
	}

	return c.SendStatus(http.StatusNoContent)
}

func DeleteForm(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("form_id is not a uuid: %v", formid)))
	}

	wsid, err := uuid.Parse(c.Params("wsid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("workspace id is not a uuid: %v", wsid)))
	}

	if err := initializers.DB.Where("form_id = ? AND workspace_id = ?", formid, wsid).Delete(&models.Path{}).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Failed to delete path"))
	}

	if err := initializers.DB.Where("form_id = ?", formid).Delete(&models.Permit{}).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Failed to delete permits"))
	}

	if err := initializers.DB.Delete(&models.Form{}, formid).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Failed to delete form"))
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
			utils.GenerateServerErrorResponse("unable to find form via path"))
	}

	form := models.Form{}

	if err := initializers.DB.Where("form_id = ?", path.FormID).First(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find form"))
	}

	return c.JSON(&form)
}
