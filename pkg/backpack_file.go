package pkg

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/vmihailenco/msgpack"
)

// GetBackpackFromFile will read and get a Backpack from a file path
func GetBackpackFromFile(path string) (b Backpack, err error) {
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = msgpack.Unmarshal(bb, &b)
	if err != nil {
		return Backpack{}, err
	}

	// Decode Templates to Base64
	b.Templates, err = decodeB64FilesMap(b.Templates)
	if err != nil {
		return Backpack{}, err
	}
	b.SortTemplates()

	// Decode Documentation
	b.Documentation, err = decodeB64FilesMap(b.Documentation)
	if err != nil {
		return Backpack{}, err
	}

	// Decode Default Values
	b.DefaultValues, err = base64.StdEncoding.DecodeString(string(b.DefaultValues))
	if err != nil {
		return Backpack{}, err
	}

	return
}

// WriteBackpackToFile will write a Backpack to file
func WriteBackpackToFile(b Backpack, path string) (err error) {
	// Encode to base64
	b.Templates, _ = encodeB64FilesMap(b.Templates)
	b.Documentation, _ = encodeB64FilesMap(b.Documentation)
	b.DefaultValues = []byte(base64.StdEncoding.EncodeToString(b.DefaultValues))

	bb, err := msgpack.Marshal(&b)
	if err != nil {
		return
	}

	return ioutil.WriteFile(path, bb, 0744)
}
