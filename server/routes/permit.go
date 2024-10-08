package routes

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

func getAbsoluteDuration(f models.Form, startDate time.Time, endDate time.Time) int {
	startEndDelta := getTimeDuration(startDate, endDate)
	duration := 0.0

	switch f.CycleData.DurationLimit.Unit {
	case models.Minutes:
		duration = math.Abs(startEndDelta.Minutes())
	case models.Hours:
		duration = math.Abs(startEndDelta.Hours())
	case models.Days:
		duration = math.Abs(startEndDelta.Hours() / 24)
	}

	return int(duration)
}

func getTimeDuration(startDate time.Time, endDate time.Time) time.Duration {
	return startDate.Sub(endDate)
}

func checkRecentPermits(
	rp []models.Permit,
	cycleStartDate time.Time,
	payload models.PermitCreateRequest,
	f models.Form,
) error {
	if len(rp) > 0 {
		requestedDuration := getAbsoluteDuration(f, *payload.StartDate, *payload.EndDate)

		// loop through permits and add up values
		totalDuration := 0
		for i := 0; i < len(rp); i++ {
			p := rp[i]

			if p.EndDate != nil {
				fmt.Println(p.EndDate)
				// handle if permit is not expired
				if p.EndDate.After(*payload.StartDate) {
					return fmt.Errorf("a permit for plate %s already exists until %s", payload.VPlate, p.EndDate.Format("Mon Jan 2 15:04"))
				}

				// if permit started before cycle start resets, use cycle start date
				if p.StartDate.Before(cycleStartDate) {
					totalDuration += getAbsoluteDuration(f, cycleStartDate, *p.EndDate)
				} else {
					totalDuration += getAbsoluteDuration(f, *p.StartDate, *p.EndDate)
				}
			}
		}

		if int(totalDuration) >= f.CycleData.DurationLimit.Value {
			return fmt.Errorf("plate %v has already exceed the maximum available duration", payload.VPlate)
		}

		if int(totalDuration)+requestedDuration > f.CycleData.DurationLimit.Value {
			return fmt.Errorf("plate %v does not have enough time left for the requested duration", payload.VPlate)
		}
	}

	return nil
}

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

	// check blacklist
	bl := initializers.DB.Where("form_id = ? AND v_plate = ?", formid, payload.VPlate).Find(&models.Blacklist{})

	if bl.RowsAffected > 0 {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Plate has been blacklisted"))
	}

	// setup expiry
	now := time.Now()
	fmt.Println(payload.EndDate)
	duration := getAbsoluteDuration(form, *payload.StartDate, *payload.EndDate)
	fmt.Println(payload.EndDate)
	fmt.Println(duration)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse((err.Error())))
	}

	if err != nil {
		return fmt.Errorf("start date is incorrect: %s", err)
	}

	if form.CycleData != nil {
		// check if entered duration is bigger than possible duration
		if duration > form.CycleData.DurationLimit.Value {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("duration is too large"))
		}
		// check if entered duration is too small
		if duration < 0 {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("duration is too small"))
		}
		// check recent permits
		recentPermits := []models.Permit{}

		cycleStartDate := utils.GenerateCycleStartDate(
			form.CycleData.ResetInterval.Value,
			form.CycleData.ResetInterval.Unit,
			form.CycleData.ResetInterval.RefDate,
			now,
		)
		if err := initializers.DB.Model(&models.Permit{}).Where("created_at > ? AND end_date > ? AND v_plate = ?", cycleStartDate, cycleStartDate, payload.VPlate).Find(&recentPermits).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("unable to find previous permits"))
		}

		if err := checkRecentPermits(recentPermits, cycleStartDate, *payload, form); err != nil {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse(err.Error()))
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
		StartDate:    payload.StartDate,
		EndDate:      payload.EndDate,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := initializers.DB.Create(&newPermit).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to create permit"))
	}

	return c.JSON(newPermit)
}

func UpdatePermit(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	permitid, err := uuid.Parse(c.Params("permitid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	payload := new(models.Permit)
	// Extract the credentials from the request body
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	payload.VPlate = strings.ToUpper(payload.VPlate)

	now := time.Now()

	permit := models.Permit{
		PermitID:     payload.PermitID,
		FormID:       payload.FormID,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PrimaryPhone: payload.PrimaryPhone,
		VPlate:       payload.VPlate,
		VMake:        payload.VMake,
		VModel:       payload.VModel,
		VColor:       payload.VColor,
		StartDate:    payload.StartDate,
		EndDate:      payload.EndDate,
		CreatedAt:    payload.CreatedAt,
		UpdatedAt:    now,
	}

	if err := initializers.DB.Model(&models.Permit{}).Where("form_id = ? AND permit_id = ?", formid, permitid).Updates(permit).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find permit"))
	}

	return c.JSON(permit)
}

func GetPermit(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	permitid, err := uuid.Parse(c.Params("permitid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	permit := models.Permit{}

	if err := initializers.DB.Model(&models.Permit{}).Where("form_id = ? AND permit_id = ?", formid, permitid).First(&permit).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find permit"))
	}

	return c.JSON(permit)
}

func GetPermits(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("id is not a uuid: %v", formid)))
	}

	permits := []models.Permit{}

	sqlWhere := fmt.Sprintf("form_id = '%s'", formid)

	if s := c.Query("search"); s != "" {
		// TODO export to utility function
		col := []string{"v_plate", "v_model", "v_make", "v_color", "first_name", "last_name", "email", "primary_phone"}
		colSql := fmt.Sprintf(" AND v_plate ILIKE '%%%s%%'", s)

		for i := 0; i < len(col); i++ {
			if i == 0 {
				colSql = colSql + fmt.Sprintf(" AND %s ILIKE '%%%s%%'", col[i], s)
			} else {
				colSql = colSql + fmt.Sprintf(" OR %s ILIKE '%%%s%%'", col[i], s)
			}
		}

		sqlWhere = sqlWhere + colSql
	}

	if err := initializers.DB.Scopes(utils.Paginate(c)).Where(sqlWhere).Find(&permits).Error; err != nil {
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

func DownloadPermits(c *fiber.Ctx) error {
	formid, err := uuid.Parse(c.Params("formid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("form_id is not a uuid: %v", formid)))
	}

	from := c.Query("from")
	to := c.Query("to")

	permits := []models.Permit{}

	if err := initializers.DB.Scopes(utils.Paginate(c)).Where("form_id = ? AND created_at > ? AND created_at < ?", formid, from, to).Find(&permits).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("unable to find permits"))
	}

	return c.JSON(permits)
}
