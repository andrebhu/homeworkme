package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Answer struct {
	Subject  string
	Filename string
}

var templates = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/homework.html"))

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", logging(index))
	log.Fatal(http.ListenAndServe(":8080", r))
}
