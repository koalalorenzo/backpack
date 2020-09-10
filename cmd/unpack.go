package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/pkg"
)

// unpackCmd represents the unpack command
var unpackCmd = &cobra.Command{
	Use:   "unpack [file.backpack]",
	Args:  cobra.ExactArgs(1),
	Run:   unpackRun,
	Short: "Opens a Backpack file to explore content",
	Long: `Explodes the backpack inside a directory. This is useful to edit a
Backpack, inspecting it or seeing default values

This command performs the opposite of "pack" command

This command accepts one argument that is backpack to extract the data
from. Unless specified via -d or --dir, the files will be extracted in a new
directory in the CWD, with the name and version of the backpack

The Backpack includes:
- backpack.yaml (containing metadata)
- values.yaml (containing the default values for the templates)
- *.nomad (representing the various go templates of Nomad Jobs)
- *.md (useful documentation)
`,
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("dir", "d", "", "specifies the directory to write into")
}

func unpackRun(cmd *cobra.Command, args []string) {
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

	err = pkg.UnpackBackpackInDirectory(&b, directory)
	if err != nil {
		log.Fatalf("Error unpacking: %s", err)
	}
}
