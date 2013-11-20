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

	r.ParseForm()

	log.Println(r.Form) //这些信息是输出到服务器端的打印信息

	if r.Method == "POST" {
		topic := models.Topic{}
	}

	topnew := models.GetAllTopic(w, r)

	log.Println("in Topic Controller")

	log.Println(topnew)

	err := templates.Execute(w, topnew)
	if err != nil {
		log.Fatal(err)
	}

}

func TopicAddController(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("AddPage").Parse(views.TopicAddHTML))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)
	templates.Parse(views.TopicTemplate)

	err := templates.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func Get(w http.ResponseWriter, r *http.Request) {
}
