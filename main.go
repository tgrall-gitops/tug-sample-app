package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, world from v1.0.13!")
		})
	log.Fatalf(
		"error: %s",
		http.ListenAndServe(":8080", nil))
}
