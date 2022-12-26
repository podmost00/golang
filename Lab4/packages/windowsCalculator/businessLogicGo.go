package windowsCalculator

import (
	"errors"
	"fmt"
)

const (
	SELECTED_TREE          = 0
	SELECTED_METAL         = 1
	SELECTED_METAL_PLASTIC = 2
	SELECTED_ONE_GLASS     = 0
	SELECTED_TWO_GLASS     = 1
)

func calculatePrice(selectedMaterial int, selectedGlass int,
	sillChecked bool, height float64, width float64) (string, error) {
	price := height * width

	if selectedMaterial == SELECTED_TREE && selectedGlass == SELECTED_ONE_GLASS {
		price *= TREE1
	} else if selectedMaterial == SELECTED_TREE && selectedGlass == SELECTED_TWO_GLASS {
		price *= TREE2
	} else if selectedMaterial == SELECTED_METAL && selectedGlass == SELECTED_ONE_GLASS {
		price *= METAL1
	} else if selectedMaterial == SELECTED_METAL && selectedGlass == SELECTED_TWO_GLASS {
		price *= METAL2
	} else if selectedMaterial == SELECTED_METAL_PLASTIC && selectedGlass == SELECTED_ONE_GLASS {
		price *= METAL_PLASTIC1
	} else if selectedMaterial == SELECTED_METAL_PLASTIC && selectedGlass == SELECTED_TWO_GLASS {
		price *= METAL_PLASTIC2
	} else {
		return "", errors.New("incorrect parameters")
	}

	if sillChecked == true {
		price += SILL
	}

	return fmt.Sprintf("%.2f", price) + " грн", nil
}
