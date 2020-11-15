package pkg

import (
	"encoding/base64"
	"sort"

	"github.com/hashicorp/go-multierror"
)

// Pack is the structure of the package/file that will use to export, share
// exchange templates, docs and configuration
type Pack struct {
	Name         string            `yaml:"name"`
	Version      string            `yaml:"version"`    // Please use semver
	Dependencies map[string]string `yaml:",omitempty"` // URLs for dependencies? TBD

	// Templates are the .nomad files that with DefaultValues to be replaced
	Templates FilesMapType `yaml:"-"`

	// JobsEvalIDs are used to store the Jobs IDs once the templates are applied
	JobsEvalIDs map[string]string `yaml:"-"`

	// DefaultValues are the specified as a yaml file. It is a []byte instead of
	// ValuesType (map[string]interface{}) because in this way we will preserve
	// comments and inline documentation. This can help a lot to getting the right
	// configuration before deploying a backpack, without having to deal with
	// online documentation versioning.
	DefaultValues []byte `yaml:"-"`

	// Documentation contains the Markdown files (.md) files. This is used to
	// provide additional information when the values.yaml inline doc is not good
	// enough due to yaml limits.
	Documentation map[string][]byte `yaml:"-"`

	// BackpackVersion will help in case the Struct changes over time (as semver)
	BackpackVersion string `yaml:"backpack_version,omitempty"`
}

// SortTemplates is used to ensure that files names sorting is respected.
// This is useful to define an "order" to follow when applying resources if that
// is needed.
func (b *Pack) SortTemplates() {
	nm := make(FilesMapType, len(b.Templates))
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

// FilesMapType is useful type to specify what kind of Mapping we are using to
// store files in the pack.
type FilesMapType map[string][]byte

func decodeB64FilesMap(ra FilesMapType) (x FilesMapType, err error) {
	for n, b64f := range ra {
		var f []byte
		f, terr := base64.StdEncoding.DecodeString(string(b64f))
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}
		ra[n] = f
	}

	if err != nil {
		return FilesMapType{}, err
	}

	return ra, err
}

func encodeB64FilesMap(ra FilesMapType) (FilesMapType, error) {
	// Encode Templates to Base64
	for n, f := range ra {
		ra[n] = []byte(base64.StdEncoding.EncodeToString(f))
	}
	return ra, nil
}
