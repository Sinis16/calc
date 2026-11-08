package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrNotFinite = errors.New("result is not a finite number")
	ErrNonReal = errors.New("result is not a real number")
)

func Add(a, b float64) (float64, error) {
	return checkFinite(a + b)
}

func Substract(a, b float64) (float64, error) {
	return checkFinite(a - b)
}

func Multiply(a, b float64) (float64, error) {
	return checkFinite(a * b)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return checkFinite(a / b)
}

func Power(a, b float64) (float64, error) {
	return checkFinite(math.Pow(a, b))
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNonReal
	}
	return CheckFinite(math.Sqrt(a))
}

func Percentage(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return CheckFinite(a * b / 100)
}

func CheckFinite(r float64) (float64, error) {
	if math.IsInf(r, 0) || math.isNan(r) {
		return 0, ErrNotFinite
	}
	return r, nil
}
