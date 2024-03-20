package some_func

import (
	"errors"
	"time"
)

type Vector struct {
	X, Y int
}

func (v Vector) Len() int {
	return v.X * v.Y
}

func Sum(x, y int) int {
	return x + y
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func SSum(x, y []int) ([]int, error) {
	if len(x) != len(y) {
		return nil, errors.New("must be same length")
	}
	z := make([]int, len(x))
	for i, a := range x {
		z[i] = a + y[i]
	}
	return z, nil
}

func LongCalc(x, y int) int {
	time.Sleep(1 * time.Second)
	return x + y
}
