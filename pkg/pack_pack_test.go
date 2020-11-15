package pkg

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestPackPackDirectory(t *testing.T) {
	// extract the test file name
	_, filename, _, _ := runtime.Caller(0)
	bundlePackageGoDir := filepath.Dir(filename)
	bundleTestDirPath := filepath.Join(bundlePackageGoDir, "../test_files/backpack/")

	b, err := GetPackFromDirectory(bundleTestDirPath)
	assert.NoError(t, err)

	assert.Equal(t, "redis", b.Name)
	assert.Equal(t, "0.1.0", b.Version)

	values := map[string]interface{}{}
	err = yaml.Unmarshal(b.DefaultValues, values)
	assert.NoError(t, err)

	assert.Equal(t, `["dc1", "dc2"]`, values["datacenters"])
	// WriteBackpackToFile(*b, filepath.Join(bundlePackageGoDir, "../test_files/new_pack.backpack"))
}
