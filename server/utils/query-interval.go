package utils

import (
	"time"

	"github.com/ordomigato/parking-app/models"
)

func GenerateCycleStartDate(
	dil int,
	dimu models.DurationMeasurementUnit,
	rd time.Time,
	cd time.Time,
) time.Time {
	var csd time.Time
	switch dimu {
	// case models.None:
	// 	q = ""
	// case models.Minutes:
	// 	q = "Minutes"
	// case models.Hours:
	// 	q = ""
	// case models.Days:
	// q = ""
	case models.Months:
		var m int
		if cd.Day() == rd.Day() {
			// check time
			if cd.Hour() == rd.Hour() {
				if cd.Minute() > rd.Minute() || cd.Minute() == rd.Minute() {
					m = int(cd.Month())
				} else {
					m = int(cd.Month()) - dil
				}
			} else if cd.Hour() > rd.Hour() {
				m = int(cd.Month())
			} else {
				m = int(cd.Month()) - dil
			}
		} else if cd.Day() > rd.Day() {
			m = int(cd.Month()) - dil + 1
		} else {
			m = int(cd.Month()) - dil
		}
		csd = time.Date(
			cd.Year(),
			time.Month(m),
			rd.Day(),
			rd.Hour(),
			rd.Minute(),
			rd.Second(),
			rd.Nanosecond(),
			rd.Location(),
		)
		// case models.Years:
		// 	q = ""
	}
	return csd
}

func GenerateReferenceDate(
	d time.Time,
) (time.Time, error) {
	// set year to sometime before now to ensure reference date is always older than current date.
	rd := time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), 0, 0, d.Location())
	// minus by one year to ensure refence date is
	return rd, nil
}
