package templating

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/hashicorp/go-multierror"
	"gitlab.com/qm64/backpack/pkg"
)

// BuildHCL will gather the Backpack templates, the default values and
// custom values to generate proper HCL that can be sent to nomad
func BuildHCL(bpk *pkg.Backpack, cv pkg.ValuesType) (o map[string][]byte, err error) {
	// Merge Values
	values := mergeValues(bpk.DefaultValues, cv)
	o = map[string][]byte{}

	for n, ot := range bpk.Templates {
		t, terr := template.New(n).Parse(string(ot))
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}

		// Apply the values
		var buf bytes.Buffer
		err = t.Execute(&buf, values)
		if terr != nil {
			err = multierror.Append(err, terr)
			continue
		}

		to, err := ioutil.ReadAll(&buf)
		if err != nil {
			err = multierror.Append(err, terr)
			continue
		}

		// Save it in the output
		o[n] = to
	}

	// in case of any error
	if err != nil {
		return nil, err
	}

	return o, nil
}
