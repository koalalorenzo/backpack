package bundle

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// GetBundleFromDirectory will look in a given path for backpack.yaml and .nomad
// files to bundle them together.
func GetBundleFromDirectory(dirPath string) (b *Bundle, err error) {
	b = &Bundle{}

	// get all the files available in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the backpack.yaml file
	bpBytes, err := ioutil.ReadFile(filepath.Join(dirPath, "backpack.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(bpBytes, b)
	if err != nil {
		log.Fatal(err)
	}

	// Get all the .nomad packages
	tempMap := map[string][]byte{}
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".nomad" {
			continue
		}

		templateBytes, terr := ioutil.ReadFile(filepath.Join(dirPath, f.Name()))
		if terr != nil {
			return nil, terr
		}
		tempMap[f.Name()] = templateBytes
	}
	b.Templates = tempMap

	return b, nil
}
