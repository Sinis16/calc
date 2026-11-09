package main

import (
	"log"
	"net/http"

	"github.com/Sinis16/calc/backend/internal/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/calculate", api.HandleCalculate)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}