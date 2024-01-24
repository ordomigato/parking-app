package routes

import (
	"fmt"
	"net/http"
	"strings"
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

	payload.VPlate = strings.ToUpper(payload.VPlate)

	form := models.Form{}
	if err := initializers.DB.Model(&models.Form{}).Preload("Path").Where("form_id = ?", formid).First(&form).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(("unable to find form")))
	}

	now := time.Now()
	exp, err := CalculateExpiryDate(form.CycleData, now, payload.Duration)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse((err.Error())))
	}

	// if (form.CycleData.)

	recentPermits := []models.Permit{}
	cycleStateDate := utils.GenerateCycleStartDate(
		form.CycleData.ResetInterval.Value,
		form.CycleData.ResetInterval.Unit,
		form.CycleData.ResetInterval.RefDate,
		now,
	)
	if err := initializers.DB.Model(&models.Permit{}).Where("created_at > ? AND expiry > ? AND v_plate = ?", cycleStateDate, cycleStateDate, payload.VPlate).Find(&recentPermits).Error; err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find previous permits"))
	}

	if len(recentPermits) > 0 {
		// loop through permits and add up values
		for i := 0; i < len(recentPermits); i++ {
			// handle if permit is not expired
			if recentPermits[i].Expiry.After(now) {
				// TODO attempt to add on to permit expiry
				return c.Status(http.StatusBadRequest).JSON(
					utils.GenerateServerErrorResponse(fmt.Sprintf("a permit for plate %s already exists", payload.VPlate)))
			}
		}
	}

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
		Expiry:       exp,
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

func CalculateExpiryDate(cd models.CycleData, ct time.Time, d int) (time.Time, error) {
	switch cd.ResetInterval.Unit {
	case models.Months:
		switch cd.DurationLimit.Unit {
		case models.Minutes:
			return CalculateExpiryForMinutesInMonth(
				cd.ResetInterval.RefDate,
				ct,
				d,
			), nil
		case models.Hours:
			return CalculateExpiryForHoursInMonth(
				cd.ResetInterval.RefDate,
				ct,
				d,
			), nil
		case models.Days:
			return CalculateExpiryForDaysInMonth(
				cd.ResetInterval.RefDate,
				ct,
				d,
			), nil
		}
	}

	return ct, fmt.Errorf("cycle data combination is incorrect")
}

func DeletePermit(c *fiber.Ctx) error {
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

	permitid, err := uuid.Parse(c.Params("permitid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("permit_id is not a uuid: %v", formid)))
	}

	if err := initializers.DB.Where("permit_id = ?", permitid).Delete(&models.Permit{}).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Failed to delete permits"))
	}

	return c.SendStatus(http.StatusNoContent)
}

func CalculateExpiryForDaysInMonth(rt time.Time, ct time.Time, d int) time.Time {
	exp := time.Date(
		ct.Year(),
		ct.Month(),
		ct.Day()+d,
		rt.Hour(),
		rt.Minute(),
		rt.Second(),
		rt.Nanosecond(),
		rt.Location(),
	)
	return exp
}

func CalculateExpiryForHoursInMonth(rt time.Time, ct time.Time, d int) time.Time {
	exp := time.Date(
		ct.Year(),
		ct.Month(),
		ct.Day(),
		ct.Hour()+d,
		ct.Minute(),
		rt.Second(),
		rt.Nanosecond(),
		rt.Location(),
	)
	return exp
}

func CalculateExpiryForMinutesInMonth(rt time.Time, ct time.Time, d int) time.Time {
	exp := time.Date(
		ct.Year(),
		ct.Month(),
		ct.Day(),
		ct.Hour(),
		ct.Minute()+d,
		rt.Second(),
		rt.Nanosecond(),
		rt.Location(),
	)
	return exp
}
