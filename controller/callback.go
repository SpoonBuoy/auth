package controller

import (
	"auth/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type CallbackController struct {
	GithubConfig *oauth2.Config
}

func NewCallbackController(githubConfig *oauth2.Config) *CallbackController {
	return &CallbackController{
		GithubConfig: githubConfig,
	}
}

func (cbc *CallbackController) GitHub(ctx *gin.Context) {
	code := ctx.Query("code")
	log.Printf("access token requested for github with the code %v ", code)
	callbackService := service.NewCallbackService(cbc.GithubConfig)
	accessTok, err := callbackService.Github(ctx, code)
	log.Printf("Access Token %v", accessTok)
	if err != nil {
		log.Printf("Error in getting github access token with err : %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
