package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/gorilla/mux"
)

type Homework struct {
	Subject  string
	Filename string
}

var templates = template.Must(template.ParseFiles("tmpl/index.html"))
var files []map[string]interface{}

func retrieveFile(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fn, err := jsonparser.GetString(data, "filename")
	subject, err := jsonparser.GetString(data, "subject")
	check := false

	for _, m := range files {
		if (m["filename"] == fn) && (m["subject"] == subject) {
			check = true
		}
	}

	if check {
		// send data to backend
		w.Write([]byte("sucesss!"))
		log.Println("Send data to backend!")
	} else {
		w.Write([]byte("nonexistent homework"))
		log.Println("homework not found :(")
	}

}

// retrieves available homeworks on the server
func getHomeworks() {
	requestURL := "http://localhost:8000/directory"

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	jsonparser.ArrayEach(resBody, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		m := make(map[string]interface{})
		m["filename"], _ = jsonparser.GetString(value, "filename")
		m["subject"], _ = jsonparser.GetString(value, "subject")
		files = append(files, m)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	var err = templates.ExecuteTemplate(w, "index.html", files)
	if err != nil {
		panic(err)
	}
}

func main() {
	getHomeworks()

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/homework", retrieveFile)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":1234", r))
}
