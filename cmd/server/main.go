package main

import (
	"log"
	"user/internal/controller"
	"user/internal/db"
	"user/internal/repository"
	"user/internal/router"
	"user/internal/service"
)

func main() {
	database, err := db.InitMYSQL()
	if err != nil {
		log.Fatal("MYSQL 连接失败：", err)
	}
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	r := router.InitRouter(userController)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务启动失败：", err)
	}
}
