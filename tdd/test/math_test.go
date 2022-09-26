package test

import (
	"testing"

	"alukart32.com/usage/tdd/internal/math"
)

const (
	ERR_MSG_FORMAT = "%s :: incorrect function result: %q + %q = %d, should be %d"
)

func TestSumIntForOneAndFour(t *testing.T) {
	arg1 := "1"
	arg2 := "4"
	expected := 5
	res := math.SumInt(arg1, arg2)
	if expected != res {
		t.Errorf(ERR_MSG_FORMAT, "math.SumInt", arg1, arg2, res, expected)
	}
}

func TestMultiplyTwoAndFour(t *testing.T) {
	arg1 := 2
	arg2 := 4
	expected := 8
	res := math.Multiply(arg1, arg2)
	if expected != res {
		t.Errorf(ERR_MSG_FORMAT, "math.SumInt", arg1, arg2, res, expected)
	}
}
