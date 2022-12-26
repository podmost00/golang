#include<stdio.h>
#include<stdbool.h>

#include "calculateTravelPrice.h"

double calculateTourPrice(int days, int selectedCountry, int selectedSeason, bool isLuxury, bool isWithGuide) {
	double result = 0;

	if (selectedCountry == BULGARIA && selectedSeason == SUMMER) {
		result = days * BULGARIA_SUMMER;
	}
	else if (selectedCountry == BULGARIA && selectedSeason == WINTER) {
		result = days * BULGARIA_WINTER;
	}
	else if (selectedCountry == GERMANY && selectedSeason == SUMMER) {
		result = days * GERMANY_SUMMER;
	}
	else if (selectedCountry == GERMANY && selectedSeason == WINTER) {
		result = days * GERMANY_WINTER;
	}
	else if (selectedCountry == POLAND && selectedSeason == SUMMER) {
		result = days * POLAND_SUMMER;
	}
	else if (selectedCountry == POLAND && selectedSeason == WINTER) {
		result = days * POLAND_WINTER;
	}

	if (isLuxury == true) {
		result += result * LUXURY;
	}

	if (isWithGuide == true) {
		result += days * GUIDE_DAY;
	}

	return result;
}