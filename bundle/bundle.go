package bundle

import (
	"github.com/vmihailenco/msgpack"
)

type Bundle struct {
	Name    string
	Version int

	// Templates are the actual job files
	Templates []string
	// Variables are the key, value that will be replaced in the job files
	Variables map[string]string

	// FormatVersion will help in case the Struct changes over time
	FormatVersion int
}

func (b *Bundle) Marshal() ([]byte, error) {
	return msgpack.Marshal(b)
}

func (b *Bundle) Unmarshal(content []byte) error {
	return msgpack.Unmarshal(content, b)
}
