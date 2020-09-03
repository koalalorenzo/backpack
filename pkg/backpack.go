package pkg

type ValuesType map[string]interface{}

type Backpack struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"` // Please use semver

	// Templates are the .nomad files that with variables to be replaced
	Templates map[string][]byte `yaml:"-"`

	// Variables are the key, value that will be replaced in the job files. The
	// values specified here are the default ones
	Variables ValuesType `yaml:"variables,omitempty"`

	// BackpackVersion will help in case the Struct changes over time (as semver)
	BackpackVersion string `yaml:"backpack_version,omitempty"`
}
