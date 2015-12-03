package main

import (
	"fmt"
	"os"

	"github.com/teddyking/tidyblob"
)

func main() {
	if isBoshReleaseDirectory() {
		os.Exit(0)
	} else {
		fmt.Fprintf(os.Stderr, "Sorry, your current directory doesn't look like release directory\n")
		os.Exit(1)
	}
}

func isBoshReleaseDirectory() bool {
	isDir, err := tidyblob.IsBoshReleaseDirectory(pwd())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to determine if current directory looks like a release directory\n")
		os.Exit(1)
	}
	return isDir
}

func pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get current working directory\n")
		os.Exit(1)
	}
	return pwd
}
