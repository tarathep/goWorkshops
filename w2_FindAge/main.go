package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tarathep/goWorkshops/findAge"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input : ")
	text, _ := reader.ReadString('\n')
	inputs := strings.Split(text, ",")

	result := findAge.Find(inputs)
	fmt.Println(result)
}
