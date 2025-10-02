package controllers

import (
	"myblog/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PostController struct {
	// Dependent services
}

func NewPostController() *PostController {
	return &PostController{
		// Inject services
	}
}

func (r *PostController) Index(ctx http.Context) http.Response {
	var posts []models.Post

	facades.Orm().Query().Find(&posts)

	return ctx.Response().Json(http.StatusOK, posts)
}

func (r *PostController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var post models.Post
	err := facades.Orm().Query().Where("id", id).First(&post)

	if err != nil || post.ID == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"error": "Post not found",
		})
	}

	return ctx.Response().Json(http.StatusOK, post)
}
