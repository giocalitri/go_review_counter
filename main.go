package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)
import (
	"go_review_counter/utils"
)

const APP_VERSION = "0.1"

var (
	versionFlag *bool   = flag.Bool("v", false, "Print the version number.")
	sentinel    *string = flag.String("sentinel", ":1:", "String to look for to count something as \"reviewed\"")
	repos       *string = flag.String("repos", "", "Repos to search through")
	token       *string = flag.String("token", "", "Your github personal access token")
)
var repoList []string

func parseArgs() {
	flag.Parse()

	// version
	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
		os.Exit(0)
	}
	// repos to be checked
	if utils.IsEmpty(*repos) {
		flag.PrintDefaults()
	}
	repoList = strings.Split(*repos, " ")
}

func main() {
	// parse the input arguments
	parseArgs()
}
