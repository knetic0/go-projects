package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

type D struct {
	Fname string
	Sname string
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("fname")
	sname := r.FormValue("sname")
	d := D{Fname: fname, Sname: sname}
	tmpl.ExecuteTemplate(w, "process.gohtml", d)
}

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}
