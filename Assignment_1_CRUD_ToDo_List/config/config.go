package config

const (
	BaseUrl = "0.0.0.0:8080"
	DbUrl   = "localhost:3306"
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
