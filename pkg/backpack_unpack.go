package pkg

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

// UnpackBackpackInDirectory will write the Backpack's backpack into a directory.
func UnpackBackpackInDirectory(b *Backpack, dirPath string) (err error) {
	for n, b64f := range b.Templates {
		// Decode Base64
		var f []byte
		f, terr := base64.StdEncoding.DecodeString(string(b64f))
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}

		terr = ioutil.WriteFile(filepath.Join(dirPath, n), f, 0744)
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}
	}

	if err != nil {
		return err
	}

	bb, err := yaml.Marshal(b)
	if err != nil {
		return
	}

	return ioutil.WriteFile(filepath.Join(dirPath, "backpack.yaml"), bb, 0744)
}
