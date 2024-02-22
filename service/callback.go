package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type CallbackService struct {
	GithubConfig *oauth2.Config
}

func NewCallbackService(githubConfig *oauth2.Config) *CallbackService {
	return &CallbackService{
		GithubConfig: githubConfig,
	}
}

func (cbs *CallbackService) Github(ctx *gin.Context, code string) (string, error) {
	verifier := oauth2.GenerateVerifier()
	tok, err := cbs.GithubConfig.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		return "", err
	}
	log.Printf("access token : %v \n refresh token : %v \n", tok.AccessToken, tok.RefreshToken)
	return tok.AccessToken, nil
}
