package hello

import (
	"GOBlogOnGAE/gae_my_app/views"
	"log"
	"net/http"
	"text/template"
)

func init() {
	router()
}

func router() {
	http.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	/*
		names, err := myServer.GetFileName("./views")

		if err != nil {
			return
		}
	*/
	/*
		for _, call := range names {
			log.Println(call)
			templates = template.Must(template.ParseFiles(call))
		}
	*/
	templates := template.Must(template.New("home").Parse(views.homePage))
	templates.New("header").Parse(views.headTemplateHtml)

	log.Println("DEBUG")

	templates.ExecuteTemplate(w, "home.html", nil)
}
