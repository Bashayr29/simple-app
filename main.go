package main

import (
	"fmt"
	"log"
	"net/http"
)

const greeting = "Hi there new pipeline build from github add webhooks :)"

var listenAndServe = http.ListenAndServe

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return mux
}

func run(addr string) error {
	return listenAndServe(addr, newMux())
}

func main() {
	log.Fatal(run(":8080"))
}
