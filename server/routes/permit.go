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

func checkRecentPermits(
	rp []models.Permit,
	csd time.Time,
	now time.Time,
	payload models.PermitCreateRequest,
	f models.Form,
) error {
	if len(rp) > 0 {
		totalDuration := 0.0
		// loop through permits and add up values
		for i := 0; i < len(rp); i++ {
			p := rp[i]

			if p.Expiry != nil {
				// handle if permit is not expired
				if p.Expiry.After(now) {
					// TODO attempt to add on to permit expiry
					return fmt.Errorf("a permit for plate %s already exists", payload.VPlate)
				}

				var delta time.Duration

				if p.CreatedAt.Before(csd) {
					delta = csd.Sub(*p.Expiry)
				} else {
					delta = p.CreatedAt.Sub(*p.Expiry)
				}

				switch f.CycleData.DurationLimit.Unit {
				case models.Minutes:
					totalDuration += math.Abs(delta.Minutes())
				case models.Hours:
					totalDuration += math.Abs(delta.Hours())
				case models.Days:
					totalDuration += math.Abs(delta.Hours() / 24)
				}
			}
		}

		if int(totalDuration) >= f.CycleData.DurationLimit.Value {
			return fmt.Errorf("plate %v has already exceed the maximum available duration", payload.VPlate)
		}

		if int(totalDuration)+payload.Duration > f.CycleData.DurationLimit.Value {
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
	exp, err := CalculateExpiryDate(form.CycleData, now, payload.Duration)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse((err.Error())))
	}

	if form.CycleData != nil {
		// check if entered duration is bigger than possible duration
		if payload.Duration > form.CycleData.DurationLimit.Value {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("duration is too large"))
		}
		// check if entered duration is too small
		if payload.Duration < 0 {
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
		if err := initializers.DB.Model(&models.Permit{}).Where("created_at > ? AND expiry > ? AND v_plate = ?", cycleStartDate, cycleStartDate, payload.VPlate).Find(&recentPermits).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(
				utils.GenerateServerErrorResponse("unable to find previous permits"))
		}

		if err := checkRecentPermits(recentPermits, cycleStartDate, now, *payload, form); err != nil {
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

func CalculateExpiryDate(cd *models.CycleData, ct time.Time, d int) (*time.Time, error) {
	if cd == nil {
		// set time to zero
		return nil, nil
	}

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

	return nil, nil
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

func CalculateExpiryForDaysInMonth(rt time.Time, ct time.Time, d int) *time.Time {
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
	return &exp
}

func CalculateExpiryForHoursInMonth(rt time.Time, ct time.Time, d int) *time.Time {
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
	return &exp
}

func CalculateExpiryForMinutesInMonth(rt time.Time, ct time.Time, d int) *time.Time {
	exp := time.Date(
		ct.Year(),
		ct.Month(),
		ct.Day(),
		ct.Hour(),
		ct.Minute()+d,
		ct.Second(),
		ct.Nanosecond(),
		rt.Location(),
	)
	return &exp
}
