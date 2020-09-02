package bundle

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// UnpackBundleInDirectory will write the bundle's backpack into a directory.
func UnpackBundleInDirectory(b *Bundle, dirPath string) (err error) {
	for n, f := range b.Templates {
		err = ioutil.WriteFile(filepath.Join(dirPath, n), f, 0744)
		if err != nil {
			return
		}
	}

	bb, err := yaml.Marshal(b)
	if err != nil {
		return
	}

	return ioutil.WriteFile(filepath.Join(dirPath, "backpack.yaml"), bb, 0744)
}
