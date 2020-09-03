package pkg

type ValuesType map[string]interface{}

type TemplatesType map[string][]byte

type Backpack struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"` // Please use semver

	// Templates are the .nomad files that with DefaultValues to be replaced
	Templates TemplatesType `yaml:"-"`

	// DefaultValues are the key, value that will be replaced in the job files. The
	// values specified here are the default ones
	DefaultValues ValuesType `yaml:"values,omitempty"`

	// BackpackVersion will help in case the Struct changes over time (as semver)
	BackpackVersion string `yaml:"backpack_version,omitempty"`
}
