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
	facades.Route().Post("/posts", postContoller.Store)
	facades.Route().Put("/posts/{id}", postContoller.Update)
	facades.Route().Delete("/posts/{id}", postContoller.Destroy)
}
