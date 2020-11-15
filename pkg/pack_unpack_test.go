package pkg

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackUnpack(t *testing.T) {
	// extract the test file name
	_, filename, _, _ := runtime.Caller(0)
	bundlePackageGoDir := filepath.Dir(filename)
	bundleTestDirPath := filepath.Join(bundlePackageGoDir, "../test_files/")

	tempDir, err := ioutil.TempDir(os.TempDir(), "backpack-unpack")
	assert.NoError(t, err)

	// Get the bundle
	b, err := GetPackFromFile(filepath.Join(bundleTestDirPath, "redis.backpack"))
	assert.NoError(t, err)

	t.Logf("Temp dir: %s", tempDir)
	err = UnpackInDirectory(&b, tempDir)
	assert.NoError(t, err)

	expectations := map[string]string{
		"backpack.yaml": "6c96179b16779474234b4f54d5294fb2d3bb122d6723f1dce80e2118ead2610c",
		"example.nomad": "55b72efc1de4957d12e23cd2c738ebf0c5677e3d55f77dfe02d04d190f3da2bb",
		"values.yaml":   "117b5a3ca6938a62297d2741671b00d3cd2b0d8febdfcb202579a864a39ef4dc",
		"README.md":     "4fcffd36bdd608eb5892f3a9250c3c9613ee4f9d37f2d2053352b9773fac8519",
	}

	// Validate that the files are correct
	for fileName, hash := range expectations {
		bckpBytes, err := ioutil.ReadFile(filepath.Join(tempDir, fileName))
		assert.NoError(t, err)
		backapckHash := sha256.Sum256(bckpBytes)
		t.Logf("%x: %s (sha256)", backapckHash, fileName)
		assert.Equal(t, hash, fmt.Sprintf("%x", backapckHash))
	}
}
