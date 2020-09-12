package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/pkg"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs [path]",
	Args:  cobra.ExactArgs(1),
	Short: "Unpack the documentation in a new directory",
	Long: `Each backpack comes with its own documentation. Use this command to 
extract the documentation for a specific backpack file. This will make sure
to have the right documentation for the right backpack file

This command accepts one argument that is backpack to extract the documentation 
from (path or URL). Unless specified via -d or --dir, the files will be 
extracted in a new directory in the CWD, with the name and version of the 
backpack
`,
	Run: unpackDocsRun,
}

func init() {
	unpackCmd.AddCommand(docsCmd)
	docsCmd.Flags().StringP("dir", "d", "", "specifies the directory to write into")
}

func unpackDocsRun(cmd *cobra.Command, args []string) {
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

	// Checks if a custom directory has been specified, otherwise unpack in
	// the backpack name.
	directory := cmd.Flag("dir").Value.String()
	if directory == "" {
		directory = filepath.Join(cwd, fmt.Sprintf("%s-%s", b.Name, b.Version))
		err = os.Mkdir(directory, 0744)
		if err != nil {
			log.Fatalf("Error creating directory: %s", err)
		}
	}

	for n, by := range b.Documentation {
		path := filepath.Join(directory, n)
		err = ioutil.WriteFile(path, by, 0744)
		if err != nil {
			log.Fatalf("Error writing files: %s", err)
		}
	}
}
