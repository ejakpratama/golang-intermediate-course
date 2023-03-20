package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"assignment_1_crud_todo_list/config"
	_ "assignment_1_crud_todo_list/docs"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	Todos        []*Todo
	postgrespool *pgxpool.Pool
	perr         error
)

// @title Todo Application
// @version 1.0
// @description This is a todo list test management application
// @contact.name Muhammad Reza Pratama
// @contact.email reza.blaze26@gmail.com
// @hostconfig.DbHost8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	postgrespool, perr = newPostgresPool(config.DbIp, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	if perr != nil {
		log.Fatal(perr)
	}

	r.HandleFunc("/todos", Get).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todos", Create).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", Put).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", Delete).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// serve http server
	log.Println("Listening in url " + config.BaseUrl)
	log.Fatal(http.ListenAndServe(config.BaseUrl, r))
}

// newPostgresPool builds a pool of pgx client.
func newPostgresPool(host, port, user, password, name string) (*pgxpool.Pool, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
	return pgxpool.Connect(context.Background(), connCfg)
}

// Create is a handler for create todos API
// @Summary Create new todos
// @Description get string by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [post]
func Create(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)

	_, _ = postgrespool.Exec(context.Background(), "insert into todo(id,name) values ($1,$2)", t.ID, t.Name)

	w.Write([]byte("Success add todo " + t.Name))
}

// Put is a handler for create todos API
// @Summary Update new todos
// @Description Update todos by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [put]
func Put(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("id = ", id)

	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	_, _ = postgrespool.Exec(context.Background(), "update todo set name=$1 where id=$2", t.Name, id)

	w.Write([]byte("Success update todo " + t.Name))

}

// Delete is a handler for create todos API
// @Summary Delete new todos
// @Description Delete string by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)

	_, _ = postgrespool.Exec(context.Background(), "delete from todo where id = $1", id)

	w.Write([]byte("Success delete todo " + id))
}

// Get is a handler for get todos API
// @Summary Get new todos
// @Description get all todo list
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [get]
func Get(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	rows, err := postgrespool.Query(context.Background(), "select id,name from todo")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			log.Println(err)
		}
		todos = append(todos, t)
	}
	fmt.Println("todos", todos)

	todosRes, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todosRes)
}

// GetByID is a handler for get todos API
// @Summary GetByID new todos
// @Description get all todo list
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [get]
func GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("id = ", id)
	var todos []Todo
	rows, err := postgrespool.Query(context.Background(), "select id,name from todo where id=$1", id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			log.Println(err)
		}
		todos = append(todos, t)
	}

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			fmt.Println("todos get", todos[i])
			todosRes, _ := json.Marshal(todos[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(todosRes)
		}
	}
}
