package bundle

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestBundleWrite(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "backpack-")
	defer os.Remove(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	b := Bundle{
		Name:      "hello-world",
		Version:   "0.1.0",
		Variables: map[string]string{"datacenter": "dc1"},
	}

	err = WriteBundleToFile(&b, tmpFile.Name())
	if err != nil {
		t.Error(err)
	}
}
