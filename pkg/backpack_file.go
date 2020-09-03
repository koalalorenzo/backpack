package pkg

import (
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
	return
}

// WriteBackpackToFile will write a Backpack to file
func WriteBackpackToFile(b Backpack, path string) (err error) {
	bb, err := msgpack.Marshal(&b)
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, bb, 0744)
}
