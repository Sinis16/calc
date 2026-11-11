package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateEndpoint(t *testing.T) {
		tests := []struct {
		name       string
		body       string
		wantStatus int
		wantResult float64
		wantError  string
	}{
		// Good ending paths
		{"add", `{"operation":"add","a":2,"b":3}`, http.StatusOK, 5, ""},
		{"subtract", `{"operation":"subtract","a":10,"b":4}`, http.StatusOK, 6, ""},
		{"multiply", `{"operation":"multiply","a":3,"b":-4}`, http.StatusOK, -12, ""},
		{"divide", `{"operation":"divide","a":10,"b":4}`, http.StatusOK, 2.5, ""},
		{"power", `{"operation":"power","a":2,"b":10}`, http.StatusOK, 1024, ""},
		{"percentage", `{"operation":"percentage","a":20,"b":50}`, http.StatusOK, 10, ""},
		{"sqrt ignores b", `{"operation":"sqrt","a":16}`, http.StatusOK, 4, ""},
		{"explicit zero operand", `{"operation":"add","a":0,"b":5}`, http.StatusOK, 5, ""},

		// Domain errors from the calculator package
		{"division by zero", `{"operation":"divide","a":1,"b":0}`, http.StatusBadRequest, 0, "division by zero"},
		{"sqrt of negative", `{"operation":"sqrt","a":-4}`, http.StatusBadRequest, 0, "result is not defined"},
		{"overflow", `{"operation":"multiply","a":1e308,"b":10}`, http.StatusBadRequest, 0, "result is not a finite number"},

		// Request validation
		{"invalid json", `hola`, http.StatusBadRequest, 0, "invalid JSON body"},
		{"missing a", `{"operation":"add","b":3}`, http.StatusBadRequest, 0, "missing operand a"},
		{"missing b", `{"operation":"add","a":3}`, http.StatusBadRequest, 0, "missing operand b"},
		{"unsupported operation", `{"operation":"modulo","a":1,"b":2}`, http.StatusBadRequest, 0, "unsupported operation: modulo"},
		{"wrong type", `{"operation":"add","a":"two","b":3}`, http.StatusBadRequest, 0, "invalid JSON body"},
	}

	router := NewRouter()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(tc.body))
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if tc.wantError != "" {
				var got errorResponse
				if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
					t.Fatalf("decoding error response: %v", err)
				}
				if got.Error != tc.wantError {
					t.Errorf("error = %q, want %q", got.Error, tc.wantError)
				}
				return
			}

			var got calculateResponse
			if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
				t.Fatalf("decoding response: %v", err)
			}
			if got.Result != tc.wantResult {
				t.Errorf("result = %v, want %v", got.Result, tc.wantResult)
			}
		})
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}