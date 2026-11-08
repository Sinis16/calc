package calculator

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	error
	}{
		{"positive", 10, 2, 12, nil},
		{"negative result", -9, 3, -6, nil},
		{"by zero", 1, 0, 1, nil},
		{"overflow", 1e308, 1e308, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Add(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	error
	}{
		{"positive", 10, 2, 8, nil},
		{"negative result", -9, 3, -12, nil},
		{"by zero", 1, 0, 1, nil},
		{"overflow", 1e308, -1e308, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Subtract(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	error
	}{
		{"positive", 10, 2, 20, nil},
		{"negative result", -9, 3, -27, nil},
		{"by zero", 1, 0, 0, nil},
		{"overflow", 1e308, 10, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Multiply(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	error
	}{
		{"positive", 10, 2, 5, nil},
		{"negative result", -9, 3, -3, nil},
		{"by zero", 1, 0, 0, ErrDivisionByZero},
		{"overflow", 1e308, 0.1, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	error
	}{
		{"positive", 10, 2, 100, nil},
		{"Negative", 4, -2, 0.0625, nil},
		{"Negative 2", -4, 2, 16, nil},
		{"Sqrt", 9, 0.5, 3, nil},
		{"by zero", 10, 0, 1, nil},
		{"overflow", 1e308, 2, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Power(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name	string
		a		float64
		want	float64
		wantErr	error
	}{
		{"positive", 100, 10, nil},
		{"zero", 0, 0, nil},
		{"negative result", -9, 0, ErrNonReal},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Sqrt(tc.a)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	error
	}{
		{"positive", 20, 50, 10, nil},
		{"negative result", -10, 50, -5, nil},
		{"by zero", 10, 0, 0, nil},
		{"overflow", 1e308, 1000, 0, ErrNotFinite},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Percentage(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
