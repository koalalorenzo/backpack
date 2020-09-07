package pkg

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/hashicorp/go-multierror"
	"github.com/vmihailenco/msgpack"
)

// GetBackpackFromFile will read and get a Backpack from a file path
func GetBackpackFromFile(path string) (b Backpack, err error) {
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = msgpack.Unmarshal(bb, &b)

	// Decode Templates to Base64
	for n, b64f := range b.Templates {
		var f []byte
		f, terr := base64.StdEncoding.DecodeString(string(b64f))
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}
		b.Templates[n] = f
	}

	b.SortTemplates()
	return
}

// WriteBackpackToFile will write a Backpack to file
func WriteBackpackToFile(b Backpack, path string) (err error) {
	// Encode Templates to Base64
	for n, f := range b.Templates {
		b.Templates[n] = []byte(base64.StdEncoding.EncodeToString(f))
	}

	bb, err := msgpack.Marshal(&b)
	if err != nil {
		return
	}

	return ioutil.WriteFile(path, bb, 0744)
}
