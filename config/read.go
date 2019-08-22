package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
	configExtension := filepath.Ext(configPath)
	switch configExtension {
	case ".json":
		json.Unmarshal(data, &im)
	case ".yaml", ".yml":
		yaml.Unmarshal(data, &im)
	default:
		err = fmt.Errorf("Wrong config extension: '%s'", configExtension)
		return
	}
	return im, nil
}
