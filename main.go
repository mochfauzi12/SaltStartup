package main

import (
	"SaltStartup/handler"
	"SaltStartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/salt_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	input := user.LoginInput{
		Email:    "mochfauzi56@gmail.com",
		Password: "password",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("Cannot Display Email Address")
		fmt.Println(err.Error())

	}
	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/z1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

}
