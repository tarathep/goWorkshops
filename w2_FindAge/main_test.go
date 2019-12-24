package main

import (
	"fmt"
	"testing"

	"github.com/tarathep/goWorkshops/findAge"
)

func TestFindDiffTH(t *testing.T) {
	input := []string{"15", "06", "2539", "TH"}
	fmt.Println(findAge.Find(input))

}
func TestFindDiffEN(t *testing.T) {
	input := []string{"15", "06", "1996", "EN"}
	fmt.Println(findAge.Find(input))
}
