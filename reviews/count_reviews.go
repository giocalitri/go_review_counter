package reviews

import (
	"fmt"
)

import (
	//"golang.org/x/oauth2"
	//"github.com/google/go-github/github"
)

import (
	"go_review_counter/config"
	"go_review_counter/utils"
)


func GetReviews(Config config.RVConfig) (listCount [][2]string) {

	token := config.GetGithubToken()
	if utils.IsEmpty(token) {
		fmt.Println("No GitHub Token found in the environment variable", config.ENV_GITHUB_VAR)
		return
	}

	return
}