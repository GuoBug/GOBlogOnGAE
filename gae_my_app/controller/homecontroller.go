package controller

import (
	"GOBlogOnGAE/gae_my_app/views"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.New("home").Parse(views.HomePage))

	templates.New("header").Parse(views.HeadTemplateHtml)
	templates.New("navbar").Parse(views.NavbarTemplateHtml)

	/* 先赋值，后取值打印 */

	Title := "Test title"

	log.Println("DEBUG")

	err := templates.Execute(w, Title)
	if err != nil {
		log.Fatal(err)
	}
}
