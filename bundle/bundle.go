package bundle

type Bundle struct {
	Name    string
	Version string // Please use semver

	// Templates are the actual job files
	Templates []string

	// Variables are the key, value that will be replaced in the job files. The
	// values specified here are the default ones
	Variables map[string]string

	// FormatVersion will help in case the Struct changes over time (as semver)
	FormatVersion string
}
