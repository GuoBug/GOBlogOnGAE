package hello

import (
	"html/template"
	"net/http"
)

func init() {
	router()
}

func router() {
	http.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("./views/home.html"))
	templates.ExecuteTemplate(w, "home.html", nil)
}
