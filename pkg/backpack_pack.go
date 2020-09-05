package pkg

import (
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
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
			err = multierror.Append(err, terr)
			continue
		}
		tempMap[f.Name()] = templateBytes
	}

	// Report the multierror
	if err != nil {
		return nil, err
	}

	b.Templates = tempMap
	return b, nil
}
