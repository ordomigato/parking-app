package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
	"gorm.io/gorm"
)

func CreateBlacklistEntry(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("form id is not a uuid: %v", formid)))
	}

	payload := new(models.BlacklistCreateRequest)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	now := time.Now()

	newBlacklist := models.Blacklist{
		FormID:    formid,
		VPlate:    strings.ToUpper(payload.VPlate),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := initializers.DB.Create(&newBlacklist).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("Plate already added to blacklist"))
		}
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("Unable to blacklist %s", payload.VPlate)))
	}

	return c.SendStatus(http.StatusNoContent)
}

func GetBlacklist(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("form id is not a uuid: %v", formid)))
	}

	blacklist := []models.Blacklist{}

	if err := initializers.DB.Scopes(utils.Paginate(c)).Where("form_id = ?", formid).Find(&blacklist).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Unable to get blacklist"))
	}

	var count int64
	if err := initializers.DB.Model(&models.Blacklist{}).Where("form_id = ?", formid).Count(&count).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to count blacklist entries"))
	}

	var response []string

	for i := 0; i < len(blacklist); i++ {
		response = append(response, blacklist[i].VPlate)
	}

	return c.JSON(fiber.Map{
		"data":  response,
		"count": count,
	})
}

func DeleteBlacklistEntry(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("form id is not a uuid: %v", formid)))
	}

	p := c.Params("plate")

	plate := models.Blacklist{
		FormID: formid,
		VPlate: p,
	}
	if err := initializers.DB.Delete(&plate).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("Unable to remove %s from blacklist", p)))
	}

	return c.SendStatus(http.StatusNoContent)
}
