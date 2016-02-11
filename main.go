package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)
import (
	"go_review_counter/utils"
	"go_review_counter/config"
)

const APP_VERSION = "0.1"

var (
	versionFlag *bool   = flag.Bool("v", false, "Print the version number.")
	sentinel    *string = flag.String("sentinel", "", "One or more strings to look for to count something as \"reviewed\"")
	repos       *string = flag.String("repos", "", "Repos to search through")
	token       *string = flag.String("token", "", "Your github personal access token")
)


func parseArgs() (confFromCommandline config.RVConfig) {
	flag.Parse()

	// version
	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
		os.Exit(0)
	}
	//// repos to be checked
	//if (utils.IsEmpty(*repos) || utils.IsEmpty(*sentinel)) {
	//	flag.PrintDefaults()
	//}
	if !utils.IsEmpty(*repos) {
		confFromCommandline.Repos = strings.Split(*repos, " ")
	}
	if !utils.IsEmpty(*sentinel) {
		confFromCommandline.Sentinel = strings.Split(*sentinel, " ")
	}
	if !utils.IsEmpty(*token) {
		confFromCommandline.Token = *token
	}
	return
}

func main() {
	// parse the input arguments
	confFromCommandline := parseArgs()
	confFromFile := config.LoadConfigFile()

	finalConfig := config.ChooseConfig(confFromCommandline, confFromFile)

	fmt.Printf("%+v\n", finalConfig)
}
