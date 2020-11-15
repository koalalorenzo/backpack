package pkg

import (
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

func getAllFileWithExtension(ext, basePath string) (FilesMapType, error) {
	tempMap := FilesMapType{}
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ext {
			continue
		}

		templateBytes, terr := ioutil.ReadFile(filepath.Join(basePath, f.Name()))
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
	return tempMap, nil
}

// GetPackFromDirectory will look in a given path for backpack.yaml and .nomad
// files to pack them together in a Pack struct.
func GetPackFromDirectory(dirPath string) (b *Pack, err error) {
	b = &Pack{}

	// Get the backpack.yaml file
	bpBytes, err := ioutil.ReadFile(filepath.Join(dirPath, "backpack.yaml"))
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bpBytes, b)
	if err != nil {
		return nil, err
	}

	// Get the values.yaml file as []byte
	b.DefaultValues, err = ioutil.ReadFile(filepath.Join(dirPath, "values.yaml"))
	if err != nil {
		return nil, err
	}

	// Get all the .nomad packages (HCL Job templates)
	tempMap, err := getAllFileWithExtension(".nomad", dirPath)
	if err != nil {
		return nil, err
	}

	// Get all the .md packages (Documentation)
	docsMap, err := getAllFileWithExtension(".md", dirPath)
	if err != nil {
		return nil, err
	}

	b.Templates = tempMap
	b.Documentation = docsMap

	b.SortTemplates()

	return b, nil
}
