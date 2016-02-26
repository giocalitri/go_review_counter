package config

import (
	"go_review_counter/utils"
)

func ChooseConfig(confFromCommandline RVConfig, confFromFile RVConfig) (finalConfig RVConfig)  {

	if len(confFromCommandline.Repos) > 0 {
		finalConfig.Repos = confFromCommandline.Repos
	} else {
		finalConfig.Repos = confFromFile.Repos
	}

	if len(confFromCommandline.Sentinel) > 0 {
		finalConfig.Sentinel = confFromCommandline.Sentinel
	} else {
		finalConfig.Sentinel = confFromFile.Sentinel
	}

	if !utils.IsEmpty(confFromCommandline.Token) {
		finalConfig.Token = confFromCommandline.Token
	} else {
		finalConfig.Token = confFromFile.Token
	}

	return
}
