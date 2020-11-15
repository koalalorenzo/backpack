package templating

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"

	"gitlab.com/qm64/backpack/pkg"
)

// BuildHCL will gather the pack templates, the default values and custom values
// to generate proper HCL that can be sent to nomad
func BuildHCL(bpk *pkg.Pack, cv pkg.ValuesType) (o map[string][]byte, err error) {
	// Get the default map properly from the raw bytes that contains comments
	dvm := map[string]interface{}{}
	err = yaml.Unmarshal(bpk.DefaultValues, dvm)
	if err != nil {
		return nil, err
	}

	// Merge Values
	values := mergeValues(dvm, cv)
	o = map[string][]byte{}

	for n, ot := range bpk.Templates {
		t, terr := template.New(n).Funcs(sprig.TxtFuncMap()).Parse(string(ot))
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
