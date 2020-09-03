package pkg

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackpackPackDirectory(t *testing.T) {
	// extract the test file name
	_, filename, _, _ := runtime.Caller(0)
	bundlePackageGoDir := filepath.Dir(filename)
	bundleTestDirPath := filepath.Join(bundlePackageGoDir, "test_bundle/")

	b, err := GetBackpackFromDirectory(bundleTestDirPath)
	assert.NoError(t, err)

	assert.Equal(t, "redis", b.Name)
	assert.Equal(t, "0.1.0", b.Version)
	assert.Equal(t, "[\"dc1\", \"dc2\"]", b.DefaultValues["datacenters"])
}
