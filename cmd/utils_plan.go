package cmd

import (
	"fmt"
	"strings"

	"github.com/hashicorp/nomad/api"
)

var (
	indentationStyle = "  "
)

// getDiffSimbol returns the symbol to use for Plan output
func getDiffSimbol(diffType string) string {
	switch diffType {
	case "Added":
		return "+"
	case "Deleted":
		return "-"
	case "Edited":
		return "~"
	default:
		return " "
	}
}

func printObjectDiff(object *api.ObjectDiff, verbosePlan bool, indentation int) {
	// Build indentation level
	indStr := strings.Repeat(indentationStyle, indentation)
	// Print only if verbose or edited
	if object.Type == "Edited" || verbosePlan {
		fmt.Printf("%s%s %s  {\n", indStr, getDiffSimbol(object.Type), object.Name)
		for _, field := range object.Fields {
			printFieldsDiff(field, verbosePlan, indentation+1)
		}

		for _, subobj := range object.Objects {
			printObjectDiff(subobj, verbosePlan, indentation+1)
		}
		fmt.Printf("%s}\n", indStr)
	}
}

func printFieldsDiff(field *api.FieldDiff, verbosePlan bool, indentation int) {
	// Build indentation level
	if field.Type == "Edited" || verbosePlan {
		indStr := strings.Repeat("  ", indentation)
		marker := getDiffSimbol(field.Type)

		if field.Type == "Edited" {
			fmt.Printf("%s%s %s: \"%s\" -> \"%s\" ", indStr, marker, field.Name, field.Old, field.New)
		} else {
			fmt.Printf("%s%s %s: \"%s\" ", indStr, marker, field.Name, field.New)
		}

		if len(field.Annotations) > 0 {
			ann := strings.Join(field.Annotations, ", ")
			fmt.Printf("(%s)", ann)
		}
		fmt.Print("\n")
	}
}
