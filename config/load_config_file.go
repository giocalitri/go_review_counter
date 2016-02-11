package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// definition of the configuration object
type RVConfig struct {
	Repos    []string
	Token    string
	Sentinel []string
}

func LoadConfigFile() RVConfig {
	// first try to load the configuration from a file
	// example copied by http://stackoverflow.com/a/16466189/257092
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := RVConfig{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	// then override the configuration with environment variables


	return configuration
}
