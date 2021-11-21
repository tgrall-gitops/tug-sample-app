package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Got to /api/hello")
		})

	http.HandleFunc("/api/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, "{ \"msg\" : \"Hello world\", \"version\" : \"v1.0.15\" }")
		})

	log.Fatalf(
		"error: %s",
		http.ListenAndServe(":8080", nil))
}
