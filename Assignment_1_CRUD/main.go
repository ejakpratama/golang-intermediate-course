package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "crud/docs"

	_ "github.com/go-sql-driver/mysql"

	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

const (
	baseURL = "0.0.0.0:8080"
	dbURL   = "localhost:3306"
)

type Todo struct {
	ID   string `json:"ID"`
	Name string `json:"Task"`
}

var Todos = []*Todo{{
	ID:   "1",
	Name: "Satu",
}, {
	ID:   "2",
	Name: "Dua",
}}

// @title			Todo Application
// @version		1.0
// @description	This is a todo list test golang application
// @contact.name	Muhammad Reza Pratama
// @contact.email	reza.blaze26@gmail.com
// @host			localhost:8080
// @BasePath		/
func main() {

	// log.SetFlags(log.Llongfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// connect DB
	db, err := sql.Open("mysql", "root:r00?R00t@tcp("+dbURL+")/todo")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("DB Connect Success")
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos", Get).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", Put).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", Delete).Methods(http.MethodDelete)
	r.HandleFunc("/todos/{id}", GetByID).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL(baseURL+"/swagger/doc.json"), //The url pointing to API definition
	// ))
	log.Fatal(http.ListenAndServe(baseURL, r))
}

// Get All is a handler for get todos API
//
//	@Summary		GetAll new todos
//	@Description	get all todo list
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/todos [get]
func Get(w http.ResponseWriter, r *http.Request) {

	todosRes, error := json.Marshal(Todos)

	if error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(todosRes)
	} else {
		w.Write([]byte("Error : " + error.Error()))
	}

}

func GetDB(w http.ResponseWriter, r *http.Request) {

	todosRes, error := json.Marshal(Todos)

	if error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(todosRes)
	} else {
		w.Write([]byte("Error : " + error.Error()))
	}

}

// Create is a handler for create todos API
//
//	@Summary		Create new todos
//	@Description	get string by ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/todos [post]
func Create(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	Todos = append(Todos, &t)
	w.Write([]byte("Success add todo " + t.Name))
}

// Delete is a handler for create todos API
//
//	@Summary		Delete new todos
//	@Description	Delete string by ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/todos/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			w.Write([]byte("Success delete todo " + id))
			return
		}
	}
}

// Put is a handler for create todos API
//
//	@Summary		Update new todos
//	@Description	Update todos by ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/todos/{id} [put]
func Put(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			var t Todo
			decoder := json.NewDecoder(r.Body)
			_ = decoder.Decode(&t)
			Todos[i] = &t
			w.Write([]byte("Success update todo " + t.ID))
			return
		}
	}
}

// GetByID is a handler for get todos API
//
//	@Summary		GetByID new todos
//	@Description	get all todo list
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	string
//	@Router			/todos/{id} [get]
func GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			todosRes, _ := json.Marshal(Todos[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(todosRes)
		}
	}
}
