package controller

import (
	i "github.com/conorsheppard/user-api-go/internal/api/service/interface"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService i.UserService
}

func (uc UserController) SetUpUserRoutes(router *gin.Engine) {
	router.POST("/user/create", func(context *gin.Context) {
		uc.userService.Create(context)
	})
	router.GET("/user/all", func(context *gin.Context) {
		uc.userService.GetAll(context)
	})
	router.PUT("/user/update/:id", func(context *gin.Context) {
		uc.userService.Update(context)
	})
	router.DELETE("/user/delete/:id", func(context *gin.Context) {
		uc.userService.Delete(context)
	})
}

func NewUserController(userService i.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
