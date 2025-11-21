package controllers

import (
	"fmt"
	"myblog/app/helpers"
	"myblog/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *UserController) Register(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"name":     "required|max_len:100",
		"email":    "required|email",
		"password": "required|min_len:6",
	})

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"error": validator.Errors().All(),
		})
	}

	var existingUser models.User
	email := ctx.Request().Input("email")
	err = facades.Orm().Query().Where("email", email).First(&existingUser)
	if err == nil && existingUser.ID > 0 {
		return ctx.Response().Json(http.StatusConflict, http.Json{
			"email": "Email already registered",
		})
	}

	password := ctx.Request().Input("password")
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to hash password",
		})
	}

	user := models.User{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("email"),
		Password: string(hashed),
	}

	if err := facades.Orm().Query().Create(&user); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	token, err := helpers.GenerateJWT(fmt.Sprint(user.ID))
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to generate token",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"token": token,
		"user": http.Json{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func (r *UserController) Login(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"email":    "required|email",
		"password": "required|min_len:6",
	})

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"error": validator.Errors().All(),
		})
	}

	email := ctx.Request().Input("email")
	var user models.User
	if err := facades.Orm().Query().Where("email", email).First(&user); err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Invalid credentials",
		})
	}

	if user.ID == 0 {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Invalid credentials",
		})
	}

	password := ctx.Request().Input("password")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Invalid credentials",
		})
	}

	token, err := helpers.GenerateJWT(fmt.Sprint(user.ID))
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to generate token",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"token": token,
		"user": http.Json{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
