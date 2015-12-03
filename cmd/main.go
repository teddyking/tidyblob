package main

import (
	"fmt"
	"os"

	"github.com/teddyking/tidyblob"
)

func main() {
	blobsYMLFilePath := fmt.Sprintf("%s/config/blobs.yml", pwd())
	boshPackagesDirPath := fmt.Sprintf("%s/packages", pwd())

	if isBoshReleaseDirectory() {
		staleBlobs := staleBlobs(blobsYMLFilePath, boshPackagesDirPath)
		printBlobs(staleBlobs)
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

func printBlobs(blobs []string) {
	var output string

	if len(blobs) > 0 {
		output = fmt.Sprintf("Found stale blobs:\n")

		for _, blob := range blobs {
			output = fmt.Sprintf("%s\n%s", output, blob)
		}
		output = fmt.Sprintf("%s\n", output)
	} else {
		output = fmt.Sprintf("Didn't detect any stale blobs\n")
	}

	fmt.Fprintf(os.Stdout, output)
}

func staleBlobs(blobsYMLFilePath, boshPackagesDirPath string) []string {
	staleBlobs, err := tidyblob.StaleBlobs(blobsYMLFilePath, boshPackagesDirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get determine stale blobs\n")
		os.Exit(1)
	}

	return staleBlobs
}
