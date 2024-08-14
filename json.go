package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(fmt.Errorf("Error marshalling JSON: %s", err).Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	code, err = w.Write(dat)
	if err != nil {
		log.Fatal("sending Response failed with error: %s", err)
	}
}
