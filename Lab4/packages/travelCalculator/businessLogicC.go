package travelCalculator

/*
#cgo CFLAGS: -I${SRCDIR}/callC
#cgo LDFLAGS: ${SRCDIR}/callC/libcallC.a
#include<stdio.h>
#include<stdbool.h>
#include "calculateTravelPrice.h"
*/
import "C"

import "fmt"

func calculatePriceC(days int64, selectedCountry int, selectedSeason int,
	isLuxury bool, isWithGuide bool) string {

	result := (float64)(C.calculateTourPrice(C.int(days), C.int(selectedCountry), C.int(selectedSeason),
		C.bool(isLuxury), C.bool(isWithGuide)))

	return fmt.Sprintf("%.2f", result) + " $"
}
