package controller

import (
	"GOBlogOnGAE/gae_my_app/models"
	"GOBlogOnGAE/gae_my_app/views"

	"html/template"
	"log"
	"net/http"
	"time"
)

func TopicController(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("topicPage").Parse(views.TopicHTML))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)
	templates.Parse(views.TopicTemplate)

	if r.Method == "POST" {

		r.ParseForm()

		log.Println(r.Form) //这些信息是输出到服务器端的打印信息

		log.Println(r.Form["title"])

		topic := models.Topic{
			Id:         3,
			Uid:        0,
			Title:      r.Form["title"][0],
			Content:    r.Form["content"][0],
			Attachment: "",
			Created:    time.Now(),
			Updated:    time.Now(),
			Views:      0,
			Author:     "Bug",
			ReplyCount: 0,
		}

		log.Println(topic)

		models.SaveTopic(w, r, &topic)
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
