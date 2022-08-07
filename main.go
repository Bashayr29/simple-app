package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Json struct {
	Data string `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Json{}
	data.Data = "Hi there!"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
