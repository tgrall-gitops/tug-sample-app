package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title     string
	Version   string
	Namespace string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	namespace := os.Getenv("KUBE_NS")

	p := Page{"Go Simple Application", "1.0.5", namespace}
	t := template.New("Template")
	t = template.Must(t.ParseFiles("static/page.tmpl"))
	err := t.ExecuteTemplate(w, "layout", p)
	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

func Sum(x int, y int) int {
	return x + y
}
