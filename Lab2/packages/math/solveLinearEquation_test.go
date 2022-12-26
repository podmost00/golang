package math

import "testing"

func TestSolveLinearEquation(t *testing.T) {
	k, x, b := 2.4, 3.0, 9.8
	expected := 17.0
	res := SolveLinearEquation(k, x, b)
	if res != expected {
		t.Errorf("Wrong result. Supposed %f, got %f", expected, res)
	}
}
