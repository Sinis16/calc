// Receives the JSON, converts to struct, validates, calls calulator, and gives back resulting JSON

package api

import (
	"encoding/json"
	"net/http"

	"github.com/Sinis16/calc/backend/internal/calculator"
)

type calculateRequest struct {
	Operation string `json:"operation"`
	A		  *float64 `json:"a"`
	B 		  *float64 `json:"b"`
}

type calculateResponse struct {
	Result float64 `json:"result"`
}

type errorResponse struct {
	Error string `json:"error"`
}

var binaryOps = map[string]func(a, b float64) (float64, error){
	"add":		calculator.Add,
	"subtract": calculator.Subtract,
	"multiply": calculator.Multiply,
	"divide": calculator.Divide,
	"power": calculator.Power,
	"percentage": calculator.Percentage,
}

func HandleCalculate(w http.ResponseWriter, r *http.Request) {
	var req calculateRequest
	
	//Here is decoded the JSON body into req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.A == nil {
		writeError(w, http.StatusBadRequest, "missing operand a")
		return
	}

	var result float64
	var err error

	if req.Operation == "sqrt" {
		result, err = calculator.Sqrt(*req.A)
	} else {
		op, ok := binaryOps[req.Operation]
		if !ok {
			writeError(w, http.StatusBadRequest, "unsupported operation: "+req.Operation)
			return
		}
		if req.B == nil {
			writeError(w, http.StatusBadRequest, "missing operand b")
			return
		}
		result, err = op(*req.A, *req.B)
		
	}

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	writeJSON(w, http.StatusOK, calculateResponse{Result: result})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errorResponse{Error: msg})
}