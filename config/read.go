package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Read will parse provided config file and return ImagesMetadata.
func Read(configPath string) (im *ImagesMetadata, err error) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	json.Unmarshal(data, &im)
	return im, nil
}
