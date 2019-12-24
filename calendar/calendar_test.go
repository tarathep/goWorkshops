package calendar

import (
	"testing"
)

func TestDateEN(t *testing.T) {

	mockDate := Date{
		Weekday: "Saturday",
		Day:     "15",
		Month:   "June",
		Year:    "1996",
	}
	if !(DateEN(mockDate) == "Saturday 15 June 1996") {
		t.Error()
	}

}

func TestDateTH(t *testing.T) {
	mockDate := Date{
		Weekday: "Saturday",
		Day:     "15",
		Month:   "June",
		Year:    "1996",
	}
	if !(DateTH(mockDate) == "วันเสาร์ 15 มิ.ย. 2539") {
		t.Error()
	}
}
func TestGetDate(t *testing.T) {
	mockDate := "15-06-1996"
	date := Date{"Saturday", "15", "June", "1996"}
	if !(GetDate(mockDate) == date) {
		t.Error()
	}
}

func TestGetWeekDayTH(t *testing.T) {
	if !(GetWeekDayTH("Monday") == "วันจันทร์") {
		t.Error()
	}
	if !(GetWeekDayTH("Tuesday") == "วันอังคาร") {
		t.Error()
	}

	if !(GetWeekDayTH("Thursday") == "วันพฤหัส") {
		t.Error()
	}
	if !(GetWeekDayTH("Friday") == "วันศุกร์") {
		t.Error()
	}
	if !(GetWeekDayTH("Saturday") == "วันเสาร์") {
		t.Error()
	}
	if !(GetWeekDayTH("Sunday") == "วันอาทิตย์") {
		t.Error()
	}
	if !(GetWeekDayTH("??") == "วันไรวะ") {
		t.Error()
	}
}

func TestGetMonthTH(t *testing.T) {
	monthsEN := map[int]string{
		1: "January", 2: "Fabruary", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "Novenber", 12: "December",
	}

	monthsTH := map[int]string{
		1: "ม.ค.", 2: "ก.พ.", 3: "มี.ค.", 4: "เม.ย", 5: "พ.ค", 6: "มิ.ย.",
		7: "ก.ค.", 8: "ส.ค.", 9: "ก.ย.", 10: "ต.ค.", 11: "พ.ย.", 12: "ธ.ค.",
	}

	for i := 1; i <= len(monthsEN); i++ {
		if !(GetMonthTH(monthsEN[i]) == monthsTH[i]) {
			t.Error()
		}

	}
	if !(GetMonthTH("???") == "เดือนอะไรวะ") {
		t.Error()
	}

}
func TestGetYearTH(t *testing.T) {
	if !(GetYearTH("1996") == "2539") {
		t.Error()
	}
}

func TestValidation3Params(t *testing.T) {
	input := []string{"15", "06", "1996"}
	check, lang := Validation(input)
	if !(check && lang == "en") {
		t.Error()
	}
}

func TestValidation4Params(t *testing.T) {
	input := []string{"15", "06", "1996", "th"}
	check, lang := Validation(input)
	if !(check && lang == "th") {
		t.Error()
	}
}

func TestValidationParamsError(t *testing.T) {
	input := []string{"32", "06"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}

func TestValidationDateError(t *testing.T) {
	input := []string{"32", "06", "1996", "th"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}
func TestValidationDateConvertError(t *testing.T) {
	input := []string{"x", "06", "1996", "th"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}
func TestValidationMonthError(t *testing.T) {
	input := []string{"31", "13", "1996", "th"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}
func TestValidationMonthConvertError(t *testing.T) {
	input := []string{"11", "xx", "1996", "th"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}
func TestValidationYearError(t *testing.T) {
	input := []string{"31", "12", "1", "th"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}
func TestValidationLangugeOptionError(t *testing.T) {
	input := []string{"31", "12", "1999", "x"}
	check, _ := Validation(input)
	if check == true {
		t.Error()
	}
}
func TestValidationLangugeOptionEN(t *testing.T) {
	input := []string{"31", "12", "1999", "en"}
	check, lang := Validation(input)
	if !(check && lang == "en") {
		t.Error()
	}
}
func TestValidationLangugeOptionTH(t *testing.T) {
	input := []string{"31", "12", "1999", "th"}
	check, lang := Validation(input)
	if !(check && lang == "th") {
		t.Error()
	}
}
func TestValidationLangugeOptionEmpty(t *testing.T) {
	input := []string{"31", "12", "1999", ""}
	check, lang := Validation(input)
	if !(check && lang == "en") {
		t.Error()
	}
}
