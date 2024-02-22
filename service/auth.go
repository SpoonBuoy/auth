package service

import (
	"fmt"

	"golang.org/x/oauth2"
)

type AuthService struct {
	GithubConfig *oauth2.Config
}

func NewAuthService(githubConfig *oauth2.Config) *AuthService {
	return &AuthService{
		GithubConfig: githubConfig,
	}
}

func (as *AuthService) Github() error {
	verifier := oauth2.GenerateVerifier()
	//generate auth url
	url := as.GithubConfig.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	//ctx.JSON(200, gin.H{"message": "success"})
	return nil
}
