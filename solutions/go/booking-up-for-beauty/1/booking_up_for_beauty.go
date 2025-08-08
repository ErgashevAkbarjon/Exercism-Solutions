package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	layouts := []string{
		"1/2/2006 15:04:05",
		"January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
	}

	var dateTime time.Time
	for _, l := range layouts {
		var err error
		dateTime, err = time.Parse(l, date)
		if err == nil {
			break
		}
	}

	return dateTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateTime := Schedule(date)
	return dateTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hour := Schedule(date).Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateTime := Schedule(date)
	layout := "Monday, January 2, 2006, at 15:04"

	return "You have an appointment on " + dateTime.Format(layout) + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
