package main

import (
	"log"
	"net/http"

	"github.com/Sinis16/calc/backend/internal/api"
)

func main() {
	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, api.NewRouter()))
}