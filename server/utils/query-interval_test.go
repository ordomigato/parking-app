package utils

import (
	"testing"
	"time"

	"github.com/ordomigato/parking-app/models"
)

func TestGenerateCycleStartDate(t *testing.T) {
	var tests = []struct {
		name string
		dil  int
		dimu models.DurationMeasurementUnit
		rd   time.Time
		cd   time.Time
		ans  time.Time
	}{
		// Test Months
		{
			"Test month if current day is beyond refence date day",
			1,
			models.Months,
			time.Date(2000, 02, 03, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 04, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 03, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"Test if current day is below refence date day",
			1,
			models.Months,
			time.Date(2000, 02, 03, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 02, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 01, 03, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"Test if current dil + dimu go below year start and you current day is above reference date day",
			3,
			models.Months,
			time.Date(2000, 02, 03, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 04, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2023, 12, 03, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"Test if current dil + dimu go below year start and you current day is below reference date day",
			3,
			models.Months,
			time.Date(2000, 02, 03, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 02, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2023, 11, 03, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"Test if date is the same but hour is below reference hour",
			1,
			models.Months,
			time.Date(2000, 02, 03, 6, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 03, 2, 0, 0, 0, time.Now().Location()),
			time.Date(2024, 01, 03, 6, 0, 0, 0, time.Now().Location()),
		},
		{
			"Test if date and hour is the same but minute is below reference hour",
			1,
			models.Months,
			time.Date(2000, 02, 03, 6, 30, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 03, 6, 12, 0, 0, time.Now().Location()),
			time.Date(2024, 01, 03, 6, 30, 0, 0, time.Now().Location()),
		},
		{
			"Test date and time is the exact same",
			1,
			models.Months,
			time.Date(2000, 02, 03, 6, 30, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 03, 6, 30, 0, 0, time.Now().Location()),
			time.Date(2024, 02, 03, 6, 30, 0, 0, time.Now().Location()),
		},
	}

	for i, tt := range tests {
		result := GenerateCycleStartDate(tt.dil, tt.dimu, tt.rd, tt.cd)
		if result != tt.ans {
			t.Errorf("Test Name %v: %s", i, tt.name)
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, tt.ans)
		}
	}
}

// func TestGenerateReferenceDate(t *testing.T) {
// 	result, _ := GenerateReferenceDate(
// 		time.Now(),
// 	)
// 	if result != time.Now() {
// 		t.Errorf("Result was incorrect, got: %s, want: %s.", result, time.Now())
// 	}
// }
