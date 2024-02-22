package controller

import (
	"auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthController struct {
	GithubConfig *oauth2.Config
}

func NewAuthController(githubConfig *oauth2.Config) *AuthController {
	return &AuthController{
		GithubConfig: githubConfig,
	}
}

func (ac *AuthController) Github(ctx *gin.Context) {
	//generate auth url
	AuthService := service.NewAuthService(ac.GithubConfig)
	err := AuthService.Github()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
