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
