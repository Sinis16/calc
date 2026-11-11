package calculator

import (
	"errors"
	"testing"
)

// Tests for operands in go,
func TestBinaryOperations(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(a, b float64) (float64, error)
		a, b    float64
		want    float64
		wantErr error
	}{
		// Add
		{"add/positive", Add, 10, 2, 12, nil},
		{"add/negative result", Add, -9, 3, -6, nil},
		{"add/zero operand", Add, 1, 0, 1, nil},
		{"add/overflow", Add, 1e308, 1e308, 0, ErrNotFinite},

		// Subtract
		{"subtract/positive", Subtract, 10, 2, 8, nil},
		{"subtract/negative result", Subtract, -9, 3, -12, nil},
		{"subtract/zero operand", Subtract, 1, 0, 1, nil},
		{"subtract/overflow", Subtract, 1e308, -1e308, 0, ErrNotFinite},

		// Multiply
		{"multiply/positive", Multiply, 10, 2, 20, nil},
		{"multiply/negative result", Multiply, -9, 3, -27, nil},
		{"multiply/zero operand", Multiply, 1, 0, 0, nil},
		{"multiply/overflow", Multiply, 1e308, 10, 0, ErrNotFinite},

		// Divide
		{"divide/positive", Divide, 10, 2, 5, nil},
		{"divide/negative result", Divide, -9, 3, -3, nil},
		{"divide/by zero", Divide, 1, 0, 0, ErrDivisionByZero},
		{"divide/overflow", Divide, 1e308, 0.1, 0, ErrNotFinite},

		// Power
		{"power/positive", Power, 10, 2, 100, nil},
		{"power/negative exponent", Power, 4, -2, 0.0625, nil},
		{"power/negative base", Power, -4, 2, 16, nil},
		{"power/fractional exponent", Power, 9, 0.5, 3, nil},
		{"power/zero exponent", Power, 10, 0, 1, nil},
		{"power/overflow", Power, 1e308, 2, 0, ErrNotFinite},
		{"power/zero to negative", Power, 0, -1, 0, ErrNotFinite},

		// Percentage
		{"percentage/positive", Percentage, 20, 50, 10, nil},
		{"percentage/negative result", Percentage, -10, 50, -5, nil},
		{"percentage/zero base", Percentage, 10, 0, 0, nil},
		{"percentage/overflow", Percentage, 1e308, 1000, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.fn(tc.a, tc.b)
			assertResult(t, got, err, tc.want, tc.wantErr)
		})
	}
}

// Test for Sqrt specifically due to working with only 1 operand
func TestSqrt(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		want    float64
		wantErr error
	}{
		{"positive", 100, 10, nil},
		{"zero", 0, 0, nil},
		{"negative input", -9, 0, ErrNonReal},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Sqrt(tc.a)
			assertResult(t, got, err, tc.want, tc.wantErr)
		})
	}
}

// assertResult checks a calculator result against the expected value and error.
func assertResult(t *testing.T, got float64, err error, want float64, wantErr error) {
	t.Helper()
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}