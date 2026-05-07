package main

import (
	"fmt"
	"log"
	"net/http"
)

const greeting = "Hi there new pipeline build from github add webhooks :)"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", newMux()))
}
