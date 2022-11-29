package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

type Todos struct {
	Item string
	Done bool
}

type PageDatas struct {
	Title string
	todos []Todos
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func check_todos(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "todo.gohtml", nil)
}

func add_todos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	todo_1 := PageDatas{Title: "Todos", todos: []Todos{{Item: r.FormValue("todo"), Done: false}}}
	tmpl.ExecuteTemplate(w, "todo.gohtml", todo_1)
}

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", index)
	http.HandleFunc("/todo", check_todos)
	http.HandleFunc("/add", add_todos)
	http.ListenAndServe(":8080", nil)
}
