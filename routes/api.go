package routes

import (
	"github.com/goravel/framework/facades"

	"myblog/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	postContoller := controllers.NewPostController()
	facades.Route().Get("/posts", postContoller.Index)
	facades.Route().Get("/posts/{id}", postContoller.Show)
}
