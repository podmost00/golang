#ifndef CALLC_H

#define CALLC_H 

#define SELECTED_TREE 0
#define SELECTED_METAL 1
#define SELECTED_METAL_PLASTIC 2
#define SELECTED_ONE_GLASS 0
#define SELECTED_TWO_GLASS 1

#define SILL 350.0
#define METAL_PLASTIC1 1.5
#define METAL_PLASTIC2 2.0
#define METAL1 0.5
#define METAL2 1.0
#define TREE1 2.5
#define TREE2 3.0

double calculateWindowPrice(int selectedMaterial, int selectedGlass,
	bool sillChecked, double height, double width);

#endif