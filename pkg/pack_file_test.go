package pkg

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPackWrite(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "backpack-")
	defer os.Remove(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	b := Pack{
		Name:          "hello-world",
		Version:       "0.1.0",
		DefaultValues: []byte(`datacenters: ["dc1", "dc2"]`),
	}

	err = WritePackToFile(b, tmpFile.Name())
	if err != nil {
		t.Error(err)
	}
}
