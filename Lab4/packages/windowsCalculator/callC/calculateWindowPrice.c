#include <stdbool.h>
#include "calculateWindowPrice.h"

double calculateWindowPrice(int selectedMaterial, int selectedGlass,
	bool sillChecked, double height, double width) {
	double price = height * width;

	if (selectedMaterial == SELECTED_TREE && selectedGlass == SELECTED_ONE_GLASS) {
		price *= TREE1;
	}
	else if (selectedMaterial == SELECTED_TREE && selectedGlass == SELECTED_TWO_GLASS) {
		price *= TREE2;
	}
	else if (selectedMaterial == SELECTED_METAL && selectedGlass == SELECTED_ONE_GLASS) {
		price *= METAL1;
	}
	else if (selectedMaterial == SELECTED_METAL && selectedGlass == SELECTED_TWO_GLASS) {
		price *= METAL2;
	}
	else if (selectedMaterial == SELECTED_METAL_PLASTIC && selectedGlass == SELECTED_ONE_GLASS) {
		price *= METAL_PLASTIC1;
	}
	else if (selectedMaterial == SELECTED_METAL_PLASTIC && selectedGlass == SELECTED_TWO_GLASS) {
		price *= METAL_PLASTIC2;
	}

	return price;
}