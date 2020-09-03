package templating

import (
	"gitlab.com/qm64/backpack/pkg"
)

func mergeValues(ms ...pkg.ValuesType) pkg.ValuesType {
	final := pkg.ValuesType{}
	for _, m := range ms {
		for k, v := range m {
			final[k] = v
		}
	}

	return final
}

// ApplyValues will gather the Backpack templates, the default values and
// custom values to generate proper HCL that can be sent to nomad
func ApplyValues(bpk *pkg.Backpack, cv pkg.ValuesType) map[string][]byte {
	// Merge Values
	// final := mergeValues(bpk.DefaultValues, cv)

	return nil
}
