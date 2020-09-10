package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/pkg"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "values [path]",
	Args:  cobra.ExactArgs(1),
	Short: "Extracts the default values of a backpack into a yaml files",
	Long: `Extracts the default values of a backpack (specified as first argument
as either an existing path or URL) into a yaml file. If no file path is 
specified in the option -f or --file a new file called values.yaml will be 
created
`,
	Run: unpackValuesRun,
}

func init() {
	unpackCmd.AddCommand(valuesCmd)
	valuesCmd.Flags().StringP("file", "f", "", "path of the file to create and write into")
}

func unpackValuesRun(cmd *cobra.Command, args []string) {
	// get a file from URL or Path
	p := getAUsablePathOfFile(args[0])

	b, err := pkg.GetBackpackFromFile(p)
	if err != nil {
		log.Fatalf("Error parsing the backpack: %s", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outputFile := cmd.Flag("file").Value.String()
	if outputFile == "" {
		outputFile = filepath.Join(cwd, "values.yaml")
	}

	err = ioutil.WriteFile(outputFile, b.DefaultValues, 0744)
	if err != nil {
		log.Fatalf("Error writing in the file %s: %s", outputFile, err)
	}
}
