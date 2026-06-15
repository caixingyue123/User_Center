package router

import (
	"user/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(userController *controller.UserController) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")
	user := api.Group("/users")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.GET("", userController.List)
		user.DELETE("/:id", userController.Delete)
	}
	return r
}
