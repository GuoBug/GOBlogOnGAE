package gae_my_app

import (
	"GOBlogOnGAE/gae_my_app/controller"
	"net/http"
)

func init() {
	router()
}

func router() {
	http.HandleFunc("/", controller.Home)
}
