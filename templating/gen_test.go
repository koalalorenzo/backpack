package templating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergingValues(t *testing.T) {
	val1 := map[string]interface{}{
		"default": "stuff",
		"hello":   "world",
		"how":     map[string]string{"are": "you"},
	}

	val2 := map[string]interface{}{
		"Iamgood": map[string]string{"and": "you"},
		"hello":   "mondo",
		"how":     "is",
	}

	final := mergeValues(val1, val2)

	assert.Equal(t, "mondo", final["hello"])
	assert.Equal(t, "is", final["how"])
	assert.Equal(t, "stuff", final["default"])
}
