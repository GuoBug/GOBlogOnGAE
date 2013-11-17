package controller

import (
	"GOBlogOnGAE/gae_my_app/models"
	"GOBlogOnGAE/gae_my_app/views"

	"html/template"
	"log"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("home").Parse(views.HomePage))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)

	/* 先赋值，后取值打印 */

	topics := models.Topic{
		TopicNum: 0,
		Title:    "Test title",
		Content:  "Test content",
		Created:  time.Now(),
	}

	log.Println("DEBUG")
	models.SaveTopic(w, r, &topics)

	topnew, err := models.GetTopic(w, r, 0)

	log.Println(topnew)

	err = templates.ExecuteTemplate(w, "home", topics)
	if err != nil {
		log.Fatal(err)
	}
}
