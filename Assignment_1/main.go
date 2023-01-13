package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const baseURL = "0.0.0.0:8080"

type Todo struct {
	Name string `json:"Task"`
}

// var Todos []*Todo

var Todos = []*Todo{{
	Name: "Satu",
}, {
	Name: "Dua",
}}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos", Get).Methods(http.MethodGet)

	log.Println("Listening in URL : " + baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}

func Create(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	Todos = append(Todos, &t)
	w.Write([]byte("Success add todo " + t.Name))
}

func Get(w http.ResponseWriter, r *http.Request) {

	todosRes, error := json.Marshal(Todos)

	if error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(todosRes)
	} else {
		w.Write([]byte("Error : " + error.Error()))
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)

	// Todos = delete(, &t)
	w.Write([]byte("Success delete todo " + t.Name))
}

/*
func main() {

	http.HandleFunc("/helloworld", PrintHello)

	log.Println("Listening in URL : " + baseURL)
	_ = http.ListenAndServe(baseURL, nil)
}

func PrintHello(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write([]byte("Hello World! HTTP GET"))
	} else {
		w.Write([]byte("Hello World! HTTP NON GET"))
	}

	/*
		if _, err := w.Write([]byte("Hello World!")); err != nil {
			log.Println("Error : ", err)
		}
	/
}
*/
