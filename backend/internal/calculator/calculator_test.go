package calculator

imports (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name	string
		a,b		float64
		want	float64
		wantErr	float64
	}{
		{"positive", 10, 2, 5, nil},
		{"negative result", -9, 3, -3, nil},
		{"by zero", 1, 0, 0, ErrDivisionByZero},
		{"overflow", 1e308, 0.1, ErrOverflow}
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