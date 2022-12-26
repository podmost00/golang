#ifndef CALLC_H

#define BULGARIA_SUMMER 100.0
#define BULGARIA_WINTER 150.0
#define GERMANY_SUMMER 160.0
#define GERMANY_WINTER 200.0
#define POLAND_SUMMER 120.0
#define POLAND_WINTER 180.0
#define GUIDE_DAY 50.0
#define LUXURY 0.2

#define SUMMER 0
#define WINTER 1
#define BULGARIA 0
#define GERMANY 1
#define POLAND 2

double calculateTourPrice(int days, int selectedCountry, int selectedSeason, bool isLuxury, bool isWithGuide);

#endif