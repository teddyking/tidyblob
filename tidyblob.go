package tidyblob

import (
	"io/ioutil"

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

type Config map[string]BlobAttributes

func Blobs(blobs_yml_file_path string) ([]Blob, error) {
	var blobs []Blob
	var config Config

	bytes, err := ioutil.ReadFile(blobs_yml_file_path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	for name, attributes := range config {
		blob := Blob{Name: name, Attributes: &attributes}
		blobs = append(blobs, blob)
	}

	return blobs, nil
}
