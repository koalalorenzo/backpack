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

func TestBackpackUnpack(t *testing.T) {
	// extract the test file name
	_, filename, _, _ := runtime.Caller(0)
	bundlePackageGoDir := filepath.Dir(filename)
	bundleTestDirPath := filepath.Join(bundlePackageGoDir, "test_bundle/")

	tempDir, err := ioutil.TempDir(os.TempDir(), "backpack-unpack")
	assert.NoError(t, err)

	// Get the bundle
	b, err := GetBackpackFromFile(filepath.Join(bundleTestDirPath, "redis.backpack"))
	assert.NoError(t, err)

	t.Logf("Temp dir: %s", tempDir)
	err = UnpackBackpackInDirectory(&b, tempDir)
	assert.NoError(t, err)

	// Validate that the files are correct
	bckpBytes, err := ioutil.ReadFile(filepath.Join(tempDir, "backpack.yaml"))
	assert.NoError(t, err)
	backapckHash := sha256.Sum256(bckpBytes)
	t.Logf("%x: backpack.yaml", backapckHash)
	assert.Equal(t, "516b882dfb71f8a8399bec4825cf43e473b2a0c03aeeeeec2cdaf87c6b5b96ec", fmt.Sprintf("%x", backapckHash))

	exampleBytes, err := ioutil.ReadFile(filepath.Join(tempDir, "example.nomad"))
	assert.NoError(t, err)
	exampleFileHash := sha256.Sum256(exampleBytes)
	t.Logf("%x: example.nomad", exampleFileHash)
	assert.Equal(t, "c4ae96f071e121008a4d374f4ff23729c223394080f834706e439d0e8a73b3be", fmt.Sprintf("%x", exampleFileHash))
}
