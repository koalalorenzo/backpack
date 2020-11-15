package pkg

import (
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

// UnpackInDirectory will extract files of a pack into a directory.
func UnpackInDirectory(b *Pack, dirPath string) (err error) {
	// Unpack Templates
	for n, f := range b.Templates {
		terr := ioutil.WriteFile(filepath.Join(dirPath, n), f, 0744)
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}
	}

	// Unpack Documentation
	for n, f := range b.Documentation {
		terr := ioutil.WriteFile(filepath.Join(dirPath, n), f, 0744)
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}
	}

	if err != nil {
		return err
	}

	// Unpack Default Values
	err = ioutil.WriteFile(filepath.Join(dirPath, "values.yaml"), b.DefaultValues, 0744)
	if err != nil {
		return err
	}

	bb, err := yaml.Marshal(b)
	if err != nil {
		return
	}

	return ioutil.WriteFile(filepath.Join(dirPath, "backpack.yaml"), bb, 0744)
}
