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
	http.HandleFunc("/topic", controller.TopicController)
	http.HandleFunc("/topic/add", controller.TopicAddController)
	http.HandleFunc("/topic/del", controller.TopicDelController)
	http.HandleFunc("/categorySide", controller.CategorySideController)

}
