package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, "{ \"msg\" : \"Hello world\", \"version\" : \"v1.0.17\" }")
		})

	log.Fatalf(
		"error: %s",
		http.ListenAndServe(":8080", nil))
}
