package routes

import (
	"casamento_api/src/controller"
	"casamento_api/src/controller/auth"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.RouterGroup) {

	r.POST("/login", controller.LoginGuest)
	r.POST("/guest", controller.CreateGuest)

	authGroup := r.Group("/auth")
	authGroup.Use(auth.VerifyTokenMiddleware)
	{
		authGroup.POST("/present", controller.CreatePresent)
		authGroup.PUT("/present", controller.SelectPresent)
		authGroup.GET("/present/selected", controller.GetSelectedPresents)
		authGroup.GET("/present", controller.GetAllPresents)
	}
}
