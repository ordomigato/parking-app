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
)

func CreatePermit(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	payload := new(models.PermitCreateRequest)
	// Extract the credentials from the request body
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	form := models.Form{}
	if err := initializers.DB.Model(&models.Form{}).Preload("Path").Where("form_id = ?", formid).First(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(("unable to find form")))
	}

	// recentPermits := []models.Permit{}
	// if err := initializers.DB.Model(&models.Permit{}).Where(fmt.Sprintf("created_at < %v", form.ReferenceTime)).Find(&recentPermits).Error; err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(
	// 		&fiber.Map{"error_message": fmt.Sprintf("unable to find previous permits: %v", err)})
	// }

	// fmt.Println(recentPermits)

	// get all permits from a certain date (duration_reset_timer) to another

	now := time.Now()

	newPermit := models.Permit{
		FormID:       formid,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PrimaryPhone: payload.PrimaryPhone,
		VPlate:       payload.VPlate,
		VMake:        payload.VMake,
		VModel:       payload.VModel,
		VColor:       payload.VColor,
		Expiry:       now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := initializers.DB.Create(&newPermit).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to create permit"))
	}

	return c.JSON(newPermit)
}

func GetPermits(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	permits := []models.Permit{}

	if err := initializers.DB.Scopes(utils.Paginate(c)).Where("form_id = ?", formid).Find(&permits).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find permits"))
	}

	var count int64
	if err := initializers.DB.Model(&models.Permit{}).Where("form_id = ?", formid).Count(&count).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to count permits"))
	}

	return c.JSON(fiber.Map{
		"data":  permits,
		"count": count,
	})
}
