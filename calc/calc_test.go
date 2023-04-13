package calc_test

import (
	"testing"

	"calculater/calc"
)

func TestAddition(t *testing.T) {
	exp := 8
	res := calc.Addition(5, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}
func TestSubstraction(t *testing.T) {
	exp := 8
	res := calc.Substraction(11, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}
