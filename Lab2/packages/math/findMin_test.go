package math

import "testing"

func TestFindMinFindsMin(t *testing.T) {
	testData := []float64{-1.5, -3.92, 2.47}
	expected := -3.92
	res, err := FindMin(testData...)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res != expected {
		t.Errorf("Wrong result. Supposed %f, got %f", expected, res)
	}
}

func TestFindMinFindsMinWhenOneParameter(t *testing.T) {
	expected := -3.92
	res, err := FindMin(expected)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res != expected {
		t.Errorf("Wrong result. Supposed %f, got %f", expected, res)
	}
}

func TestFindMinThrowsError(t *testing.T) {
	_, err := FindMin()
	if err == nil {
		t.Errorf("Wrong result. The error was not caught.")
	}
}
