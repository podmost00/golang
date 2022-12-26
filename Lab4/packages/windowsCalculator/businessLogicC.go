package windowsCalculator

/*
#cgo CFLAGS: -I${SRCDIR}/callC
#cgo LDFLAGS: ${SRCDIR}/callC/libcallC.a
#include <stdbool.h>
#include "calculateWindowPrice.h"
*/
import "C"

import (
	"fmt"
)

func calculatePriceC(selectedMaterial int, selectedGlass int,
	sillChecked bool, height float64, width float64) string {

	result := (float64)(C.calculateWindowPrice(C.int(selectedMaterial), C.int(selectedGlass),
		C.bool(sillChecked), C.double(height), C.double(width)))

	return fmt.Sprintf("%.2f", result) + "грн"
}
