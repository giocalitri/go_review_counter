package config

import (
	"os"
)

const ENV_GITHUB_VAR = "GITHUB_TOKEN"

func GetGithubToken() string {
	return os.Getenv(ENV_GITHUB_VAR)
}

