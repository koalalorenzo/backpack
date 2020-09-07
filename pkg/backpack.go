package pkg

import (
	"sort"
)

type TemplatesType map[string][]byte

// Backpack is
type Backpack struct {
	Name         string            `yaml:"name"`
	Version      string            `yaml:"version"` // Please use semver
	Dependencies map[string]string // URLs for dependencies? TBD

	// Templates are the .nomad files that with DefaultValues to be replaced
	Templates TemplatesType `yaml:"-"`

	// JobsEvalIDs are used to store the Jobs IDs once the templates are applied
	JobsEvalIDs map[string]string `yaml:"-"`

	// DefaultValues are the key, value that will be replaced in the job files. The
	// values specified here are the default ones
	DefaultValues ValuesType `yaml:"values,omitempty"`

	// BackpackVersion will help in case the Struct changes over time (as semver)
	BackpackVersion string `yaml:"backpack_version,omitempty"`
}

// SortTemplates is used to ensure that files names sorting is respected.
// This is useful to define an "order" to follow when applying resources if that
// is needed.
func (b *Backpack) SortTemplates() {
	nm := make(TemplatesType, len(b.Templates))
	sk := make([]string, 0, len(b.Templates))

	// Get the keys (file names)
	for k := range b.Templates {
		sk = append(sk, k)
	}

	sort.Strings(sk)

	// Populate the new map.
	for _, k := range sk {
		nm[k] = b.Templates[k]
	}

	b.Templates = nm
}
