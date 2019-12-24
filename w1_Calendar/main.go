package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tarathep/goWorkshops/calendar"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Date (DD,MM,YYYY): ")
	text, _ := reader.ReadString('\n')
	sp := strings.Split(text, ",")

	check, lang := calendar.Validation(sp)
	if check {
		date := strings.Trim(sp[0]+"-"+sp[1]+"-"+sp[2], "\n")

		if lang == "en" {
			fmt.Println(calendar.DateEN(calendar.GetDate(date)))
		} else {
			fmt.Println(calendar.DateTH(calendar.GetDate(date)))
		}
	}
}
