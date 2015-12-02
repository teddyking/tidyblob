package tidyblob

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Blob struct {
	Name       string
	Attributes *BlobAttributes
}

type BlobAttributes struct {
	ObjectID string `yaml:"object_id"`
	SHA      string `yaml:"sha"`
	Size     int    `yaml:"size"`
}

type Spec struct {
	Name  string   `yaml:"name"`
	Files []string `yaml:"files"`
}

type Config map[string]BlobAttributes

func Blobs(blobs_yml_file_path string) ([]string, error) {
	var blobs []string
	var config Config

	bytes, err := ioutil.ReadFile(blobs_yml_file_path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	for name, _ := range config {
		blobs = append(blobs, name)
	}

	return blobs, nil
}

func RequiredBlobs(bosh_packages_dir_path string) ([]string, error) {
	var requiredBlobs []string

	spec_file_path_pattern := fmt.Sprintf("%s/*/spec", bosh_packages_dir_path)
	spec_file_paths, err := filepath.Glob(spec_file_path_pattern)
	if err != nil {
		return nil, err
	}

	for _, spec_file_path := range spec_file_paths {
		var spec Spec

		bytes, err := ioutil.ReadFile(spec_file_path)
		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(bytes, &spec); err != nil {
			return nil, err
		}

		requiredBlobs = append(requiredBlobs, spec.Files...)
	}

	return uniq(requiredBlobs), nil
}

func StaleBlobs(blobs_yml_file_path, bosh_packages_dir_path string) ([]string, error) {
	var staleBlobs []string

	blobs, err := Blobs(blobs_yml_file_path)
	if err != nil {
		return nil, err
	}

	requiredBlobs, err := RequiredBlobs(bosh_packages_dir_path)
	if err != nil {
		return nil, err
	}

	for _, blob := range blobs {
		if !containsString(requiredBlobs, blob) {
			staleBlobs = append(staleBlobs, blob)
		}
	}

	return staleBlobs, nil
}

func IsBoshReleaseDirectory(boshReleasePath string) (bool, error) {
	requiredDirs := []string{"jobs", "packages", "src"}

	for _, requiredDir := range requiredDirs {
		_, err := os.Stat(fmt.Sprintf("%s/%s", boshReleasePath, requiredDir))
		if os.IsNotExist(err) {
			return false, nil
		}

		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// This func was essentially copied from a golang-nuts post:
// https://groups.google.com/d/msg/golang-nuts/-pqkICuokio/KqJ0091EzVcJ
// It has been renamed from 'removeDuplicates' to 'uniq' and modified to work with strings
func uniq(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func containsString(slice []string, str string) bool {
	for _, sliceElement := range slice {
		if sliceElement == str {
			return true
		}
	}
	return false
}
