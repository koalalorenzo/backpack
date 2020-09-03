package pkg

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// GetBackpackFromDirectory will look in a given path for backpack.yaml and .nomad
// files to Backpack them together.
func GetBackpackFromDirectory(dirPath string) (b *Backpack, err error) {
	b = &Backpack{}

	// get all the files available in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// Get the backpack.yaml file
	bpBytes, err := ioutil.ReadFile(filepath.Join(dirPath, "backpack.yaml"))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bpBytes, b)
	if err != nil {
		return nil, err
	}

	// Get all the .nomad packages
	tempMap := TemplatesType{}
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".nomad" {
			continue
		}

		templateBytes, terr := ioutil.ReadFile(filepath.Join(dirPath, f.Name()))
		if terr != nil {
			return nil, terr
		}

		// Encode files in base64
		b64file := base64.StdEncoding.EncodeToString(templateBytes)

		tempMap[f.Name()] = []byte(b64file)
	}
	b.Templates = tempMap

	return b, nil
}
