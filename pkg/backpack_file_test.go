package pkg

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestBackpackWrite(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "backpack-")
	defer os.Remove(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	b := Backpack{
		Name:          "hello-world",
		Version:       "0.1.0",
		DefaultValues: []byte(`datacenters: ["dc1", "dc2"]`),
	}

	err = WriteBackpackToFile(b, tmpFile.Name())
	if err != nil {
		t.Error(err)
	}
}
