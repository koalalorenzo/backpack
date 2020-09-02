package bundle

import (
	"io/ioutil"

	"github.com/vmihailenco/msgpack"
)

func (b *Bundle) MarshalMsgPack() ([]byte, error) {
	return msgpack.Marshal(b)
}

func (b *Bundle) UnmarshalMsgPack(content []byte) error {
	return msgpack.Unmarshal(content, b)
}

// GetBundleFromFile will read and get a Bundle from a file path
func GetBundleFromFile(path string) (b *Bundle, err error) {
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = b.UnmarshalMsgPack(bb)
	return
}

// WriteBundleToFile will write a Bundle to file
func WriteBundleToFile(b *Bundle, path string) (err error) {
	bb, err := b.MarshalMsgPack()
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, bb, 0744)
}
