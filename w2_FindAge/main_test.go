package main

import (
	"testing"

	"github.com/tarathep/goWorkshops/findAge"
)

func TestFindDiffTH(t *testing.T) {
	input := []string{"15", "06", "2539", "TH"}
	if (findAge.Find(input)) != "23 ปี  6 เดือน  9 วัน" {
		t.Error()
	}

}
func TestFindDiffEN(t *testing.T) {
	input := []string{"15", "06", "1996", "EN"}
	if (findAge.Find(input)) != "23 year(s)  6 month(s)  9 day(s)" {
		t.Error()
	}
}
