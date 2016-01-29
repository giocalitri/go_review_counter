package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
)

const APP_VERSION = "0.1"

var (
	versionFlag *bool   = flag.Bool("v", false, "Print the version number.")
	sentinel    *string = flag.String("sentinel", ":1:", "String to look for to count something as \"reviewed\"")
	repos       *string = flag.String("repos", "", "Repos to search through")
)

func parseArgs() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
		os.Exit(0)
	}
	if *repos == "" {
		flag.PrintDefaults()
	}
	repoList := strings.Split(*repos, " ")

	return repoList
}

func main() {
	parseArgs()
}
