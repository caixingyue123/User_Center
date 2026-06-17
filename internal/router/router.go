package router

import (
	"user/internal/controller"
	"user/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(userController *controller.UserController) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}
	user := api.Group("")
	user.Use(middleware.AuthMiddleWare())
	{
		user.GET("", userController.List)
		user.GET("/profile", userController.GetProfile)
		user.PUT("/profile", userController.UpdateProfile)
		user.DELETE("/:id", userController.Delete)
	}
	return r
}