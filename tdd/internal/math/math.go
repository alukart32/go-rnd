package math

import (
	"strconv"
)

func SumInt(a, b string) int {
	a1, e := strconv.Atoi(a)
	if e != nil {
		panic(e)
	}
	b1, e := strconv.Atoi(b)
	if e != nil {
		panic(e)
	}
	return a1 + b1
}

func Multiply(a, b int) int {
	return a * b
}
