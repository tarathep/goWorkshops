package calendar

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Date is model
type Date struct {
	Weekday string
	Day     string
	Month   string
	Year    string
}

func GetDate(dateString string) Date {
	dt, _ := time.Parse("02-01-2006", dateString)
	return Date{
		fmt.Sprint(dt.Weekday()),
		fmt.Sprint(dt.Day()),
		fmt.Sprint(dt.Month()),
		fmt.Sprint(dt.Year()),
	}
}

func GetWeekDayTH(weekDayEN string) string {
	switch weekDayEN {
	case "Monday":
		return "วันจันทร์"
	case "Tuesday":
		return "วันอังคาร"
	case "Wednesday":
		return "วันพุธ"
	case "Thursday":
		return "วันพฤหัส"
	case "Friday":
		return "วันศุกร์"
	case "Saturday":
		return "วันเสาร์"
	case "Sunday":
		return "วันอาทิตย์"
	}
	return "วันไรวะ"
}
func GetMonthTH(monthEN string) string {
	switch monthEN[0:3] {
	case "Jan":
		return "ม.ค."
	case "Fab":
		return "ก.พ."
	case "Mar":
		return "มี.ค."
	case "Apr":
		return "เม.ย"
	case "May":
		return "พ.ค"
	case "Jun":
		return "มิ.ย."
	case "Jul":
		return "ก.ค."
	case "Aug":
		return "ส.ค."
	case "Sep":
		return "ก.ย."
	case "Oct":
		return "ต.ค."
	case "Nov":
		return "พ.ย."
	case "Dec":
		return "ธ.ค."
	}
	return "เดือนอะไรวะ"
}
func GetYearTH(yearEN string) string {
	year, _ := strconv.Atoi(yearEN)
	return fmt.Sprint(year + 543)
}

func Validation(sp []string) (bool, string) {
	i := 0
	if !(len(sp) == 3 || len(sp) == 4) {
		fmt.Println("error input param format (dd,mm,yyyy,en)")
		return false, ""
	}

	for _, d := range sp {
		switch i {
		case 0:
			// date
			iD, err := strconv.Atoi(d)
			if err != nil {
				fmt.Println("error input convert date format")
				return false, ""
			}
			if !(iD > 0 && iD < 32) {
				fmt.Println("error input date format")
				return false, ""
			}
			break
		case 1:
			// month
			iM, err := strconv.Atoi(d)
			if err != nil {
				fmt.Println("error input convert month format")
				return false, ""
			}
			if !(iM > 0 && iM < 13) {
				fmt.Println("error input month format")
				return false, ""
			}
			break
		case 2:
			// year
			iY := strings.Trim(d, "\n")
			if len(iY) != 4 {
				fmt.Println("error input year format (yyyy)")
				return false, ""
			}
			break
		case 3:
			// en/th
			language := strings.Trim(d, "\n")
			if language == "th" {
				return true, "th"
			} else if language == "en" {
				return true, "en"
			} else if language == "" {
				return true, "en"
			}
			fmt.Println("error input option language (EN/TH)")
			return false, ""
		}
		i++
	}
	return true, "en"
}

//DateEN is Print ENG DATE Function format WeekDay/Day/Month/Year
func DateEN(date Date) string {
	return fmt.Sprint(date.Weekday + " " + date.Day + " " + date.Month + " " + date.Year)
}

//DateTH is Print THA DATE Function format WeekDay/Day/Month/Year
func DateTH(date Date) string {
	return fmt.Sprint(GetWeekDayTH(date.Weekday) + " " + date.Day + " " + GetMonthTH(date.Month) + " " + GetYearTH(date.Year))
}

//--Workshop2
