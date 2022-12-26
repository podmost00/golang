package math

import "testing"

func TestGetAverageReturnsAverage(t *testing.T) {
	testData := []float64{-1.5, -8.24, 2.47}
	expected := -2.424
	res, err := GetAverage(testData...)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res != expected {
		t.Errorf("Wrong result. Supposed %f, got %f", expected, res)
	}
}

func TestGetAverageWhenOneParameter(t *testing.T) {
	expected := -3.92
	res, err := GetAverage(expected)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res != expected {
		t.Errorf("Wrong result. Supposed %f, got %f", expected, res)
	}
}

func TestGetAverageThrowsError(t *testing.T) {
	_, err := GetAverage()
	if err == nil {
		t.Errorf("Wrong result. The error was not caught.")
	}
}
