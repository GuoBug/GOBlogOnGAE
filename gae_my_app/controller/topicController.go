package controller

import (
	"GOBlogOnGAE/gae_my_app/models"
	"GOBlogOnGAE/gae_my_app/views"

	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func TopicController(w http.ResponseWriter, r *http.Request) {

	var i int64

	templates := template.Must(template.New("topicPage").Parse(views.TopicHTML))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)

	//这样可以取得最大的文章号码
	_, i = models.GetAllTopic(w, r)

	log.Printf("Get Topic Count [%d]", i)

	//templates.Parse(views.TopicTemplate)

	if r.Method == "POST" {

		r.ParseForm()

		log.Println(r.Form) //这些信息是输出到服务器端的打印信息

		log.Println(r.Form["title"])

		topic := models.Topic{
			Id:         i + 1,
			Uid:        0,
			Title:      r.Form["title"][0],
			Category:   r.Form["category"][0],
			Content:    r.Form["content"][0],
			Attachment: "",
			Created:    time.Now(),
			Updated:    time.Now(),
			Views:      0,
			Author:     "Bug",
			ReplyCount: 0,
		}

		category := models.Category{
			Title: r.Form["category"][0],
		}

		log.Println(topic)

		models.SaveTopic(w, r, &topic)
		models.SaveCategroy(w, r, &category)
	}

	topnew, _ := models.GetAllTopic(w, r)

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

func TopicDelController(w http.ResponseWriter, r *http.Request) {
	log.Printf("OOOOOOOOOOOOOOOOOOOO\n")

	id, err := strconv.ParseInt(r.URL.RawQuery, 10, 64)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Get Delete Id is [%d]", id)
	models.DeleteTopic(w, r, id)
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
}
