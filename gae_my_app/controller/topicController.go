package controller

import (
	"GOBlogOnGAE/gae_my_app/models"
	"GOBlogOnGAE/gae_my_app/views"

	"html/template"
	"log"
	"net/http"
)

func TopicController(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("topicPage").Parse(views.TopicHTML))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)
	templates.Parse(views.TopicTemplate)

	topnew := models.GetAllTopic(w, r)

	log.Println("in Topic Controller")

	log.Println(topnew)

	err := templates.Execute(w, topnew)
	if err != nil {
		log.Fatal(err)
	}
}
