package templating

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/qm64/backpack/pkg"
)

func TestBuildHCL(t *testing.T) {
	bp := pkg.Backpack{
		Name:    "test-backpack",
		Version: "0.1.0",
		Templates: map[string][]byte{
			"job1.nomad": []byte("{{ .variable }}-{{ .default }}"),
			"job2.nomad": []byte("{{ .default }}-{{ .variable }}"),
		},
		DefaultValues: map[string]interface{}{"variable": "old", "default": 1},
	}

	// Defining new values to apply
	nv := map[string]interface{}{"variable": "new"}

	o, err := BuildHCL(&bp, nv)
	assert.NoError(t, err)

	assert.Equal(t, []byte("new-1"), o["job1.nomad"])
	assert.Equal(t, []byte("1-new"), o["job2.nomad"])
}
