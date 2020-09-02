package bundle

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBudleReadDirectory(t *testing.T) {
	// extract the test file name
	_, filename, _, _ := runtime.Caller(0)
	bundlePackageGoDir := filepath.Dir(filename)
	bundleTestDirPath := filepath.Join(bundlePackageGoDir, "test_bundle/")

	t.Logf("path: %s", bundleTestDirPath)

	b, err := GetBundleFromDirectory(bundleTestDirPath)
	assert.NoError(t, err)

	assert.Equal(t, "redis", b.Name)
	assert.Equal(t, "0.1.0", b.Version)
	assert.Equal(t, "[\"dc1\", \"dc2\"]", b.Variables["datacenters"])
}
