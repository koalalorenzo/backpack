package bundle

import (
	"io/ioutil"
)

// GetBundleFromFile will read and get a Bundle from a file path
func GetBundleFromFile(path string) (b Bundle, err error) {
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = b.Unmarshal(bb)
	return
}

// WriteBundleToFile will write a Bundle to file
func WriteBundleToFile(b *Bundle, path string) (err error) {
	bb, err := b.Marshal()
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, bb, 0744)
}
