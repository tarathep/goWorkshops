package dateTime

import (
	"fmt"
	"time"

	"github.com/tarathep/goWorkshops/calendar"
)

type DateTime struct {
	WeekDay string
	Day     string
	Month   string
	Year    string
}

func (dateTime DateTime) GetWeekDay(language string) string {
	if language == "TH" {
		return calendar.GetWeekDayTH(dateTime.WeekDay)
	}
	return dateTime.WeekDay
}
func (dateTime DateTime) Init() {
	dt, _ := time.Parse("02-01-2006", dateTime.Day+"-"+dateTime.Month+"-"+dateTime.Year)
	dateTime = DateTime{
		fmt.Sprint(dt.Weekday()),
		fmt.Sprint(dt.Day()),
		fmt.Sprint(dt.Month()),
		fmt.Sprint(dt.Year()),
	}
}
