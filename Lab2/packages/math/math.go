package math

import (
	"errors"
	"math"
)

func FindMin(args ...float64) (float64, error) {
	if len(args) == 0 {
		return 0.0, errors.New("no arguments")
	}

	if len(args) == 1 {
		return args[0], nil
	}

	currentMin := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < currentMin {
			currentMin = args[i]
		}
	}

	return currentMin, nil
}

func GetAverage(args ...float64) (float64, error) {
	if len(args) == 0 {
		return 0.0, errors.New("no arguments")
	}

	if len(args) == 1 {
		return args[0], nil
	}

	sum := 0.0
	for _, item := range args {
		sum += item
	}

	average := round(sum/float64(len(args)), 3)
	return average, nil
}

func SolveLinearEquation(k float64, x float64, b float64) float64 {
	return k*x + b
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))

	intermed := x * pow
	_, frac := math.Modf(intermed)
	
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
