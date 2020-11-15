package pkg

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/vmihailenco/msgpack"
)

// GetPackFromFile will read and get a Pack from a file path and return as a
// Pack struct
func GetPackFromFile(path string) (b Pack, err error) {
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = msgpack.Unmarshal(bb, &b)
	if err != nil {
		return Pack{}, err
	}

	// Decode Templates to Base64
	b.Templates, err = decodeB64FilesMap(b.Templates)
	if err != nil {
		return Pack{}, err
	}
	b.SortTemplates()

	// Decode Documentation
	b.Documentation, err = decodeB64FilesMap(b.Documentation)
	if err != nil {
		return Pack{}, err
	}

	// Decode Default Values
	b.DefaultValues, err = base64.StdEncoding.DecodeString(string(b.DefaultValues))
	if err != nil {
		return Pack{}, err
	}

	return
}

// WritePackToFile will write a Pack to file
func WritePackToFile(b Pack, path string) (err error) {
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
