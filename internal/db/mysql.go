package db

import (
	"fmt"
	"user/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMYSQL() (*gorm.DB, error) {
	dsn := "root:cxy07181123@tcp(127.0.0.1:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	fmt.Println("MySQL connected successfully")

	return database, nil

}
