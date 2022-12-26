package travelCalculator

import (
	"errors"
	"fmt"
)

const (
	SUMMER   = 0
	WINTER   = 1
	BULGARIA = 0
	GERMANY  = 1
	POLAND   = 2
)

func calculatePrice(days int64, selectedCountry int, selectedSeason int,
	isLuxury bool, isWithGuide bool) (string, error) {
	var res float64

	if selectedCountry == BULGARIA && selectedSeason == SUMMER {
		res = float64(days) * BULGARIA_SUMMER
	} else if selectedCountry == BULGARIA && selectedSeason == WINTER {
		res = float64(days) * BULGARIA_WINTER
	} else if selectedCountry == GERMANY && selectedSeason == SUMMER {
		res = float64(days) * GERMANY_SUMMER
	} else if selectedCountry == GERMANY && selectedSeason == WINTER {
		res = float64(days) * GERMANY_WINTER
	} else if selectedCountry == POLAND && selectedSeason == SUMMER {
		res = float64(days) * POLAND_SUMMER
	} else if selectedCountry == POLAND && selectedSeason == WINTER {
		res = float64(days) * POLAND_WINTER
	} else {
		return "", errors.New("incorrect parameters")
	}

	if isLuxury == true {
		res += res * LUXURY
	}

	if isWithGuide == true {
		res += float64(days) * GUIDE_DAY
	}

	return fmt.Sprintf("%.2f", res) + " $", nil
}
