package bundle

type Bundle struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"` // Please use semver

	// Templates are the .nomad files that with variables to be replaced
	Templates map[string][]byte `yaml:"templates"`

	// Variables are the key, value that will be replaced in the job files. The
	// values specified here are the default ones
	Variables map[string]interface{} `yaml:"values,omitempty"`

	// BackpackVersion will help in case the Struct changes over time (as semver)
	BackpackVersion string `yaml:"backpack_version,omitempty"`
}
