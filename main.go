package main

import (
	"auth/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("bullshit programmer...")

	r := gin.Default()
	conf := &oauth2.Config{
		ClientID:     "e52d7c7e965a974ac487",
		ClientSecret: "7eed485bc46cc22590c520fc3489a551c26c2d99",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	authController := controller.NewAuthController(conf)
	callbackController := controller.NewCallbackController(conf)

	r.GET("/github", authController.Github)
	r.GET("/auth", callbackController.GitHub)

	r.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	r.Run(":8080")
}
