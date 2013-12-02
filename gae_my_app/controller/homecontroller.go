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
	templates.Parse(views.TopicTemplate)

	/* 先赋值，后取值打印 */
	var topics [2]models.Topic

	topics[0] = models.Topic{
		Id:      1,
		Title:   "Test title",
		Content: "Test content",
		Created: time.Now(),
	}

	models.SaveTopic(w, r, &topics[0])

	topics[1] = models.Topic{
		Id:      2,
		Title:   "2222222222",
		Content: "T2222",
		Created: time.Now(),
	}

	models.SaveTopic(w, r, &topics[1])

	//models.SaveCategroy(w, r)

	topnew, _ := models.GetAllTopic(w, r)

	category, _ := models.GetCategory(w, r, "事实上")

	log.Printf("***********************  %v", category)

	err := templates.ExecuteTemplate(w, "home", topnew)
	if err != nil {
		log.Fatal(err)
	}
}

func CategorySideController(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.New("categorySide").Parse(views.CategorySidePage))
	templates.New("header").Parse(views.HeadTemplateHtml)

	category, _ := models.GetAllCategory(w, r)

	err := templates.Execute(w, category)
	if err != nil {
		log.Fatal(err)
	}

}
