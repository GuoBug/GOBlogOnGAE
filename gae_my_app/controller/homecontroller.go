package controller

import (
	"GOBlogOnGAE/gae_my_app/models"
	"GOBlogOnGAE/gae_my_app/views"

	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("home").Parse(views.HomePage))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)
	templates.Parse(views.TopicTemplate)

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
