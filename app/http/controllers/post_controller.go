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

func (r *PostController) Store(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"title": "required|max_len:255",
		"body":  "required",
	})

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"errors": validator.Errors().All(),
		})
	}

	post := models.Post{
		Title:  ctx.Request().Input("title"),
		Body:   ctx.Request().Input("body"),
		Status: "DRAFT",
	}

	if err := facades.Orm().Query().Create(&post); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusCreated, post)
}

func (r *PostController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var post models.Post
	if err := facades.Orm().Query().Where("id", id).First(&post); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"error": "Post not found",
		})
	}

	validator, err := ctx.Request().Validate(map[string]string{
		"title": "required|max_len:255",
		"body":  "required",
	})

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"errors": validator.Errors().All(),
		})
	}

	post.Title = ctx.Request().Input("title")
	post.Body = ctx.Request().Input("body")
	if ctx.Request().Input("status") != "" {
		post.Status = ctx.Request().Input("status")
	}

	if err := facades.Orm().Query().Save(&post); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, post)
}

func (r *PostController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var post models.Post
	if err := facades.Orm().Query().Where("id", id).First(&post); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"error": "Post not found",
		})
	}

	if _, err := facades.Orm().Query().Delete(&post); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"success": true})
}
