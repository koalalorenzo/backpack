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
	Use:     "unpack [file.backpack]",
	Aliases: []string{"pull"},
	Args:    cobra.ExactArgs(1),
	Run:     unpackRun,
	Short:   "Opens a pack file to explore content",
	Long: `Explodes/Open the pack inside a directory. This is useful to edit a
pack, inspecting it or seeing default values... or if you are looking for
something inside it and you know it is at the bottom of the backpack!

This command accepts one argument that is the path of a pack to extract the data 
from (path or URL). Unless specified via -d or --dir, the files will be 
extracted in a new directory in the CWD, with the name and version of the 
original pack.

A pack includes the following files:
- backpack.yaml (containing metadata)
- values.yaml (containing the default values for the templates)
- *.nomad (representing the various go templates of Nomad Jobs)
- *.md (useful documentation)

This command performs the opposite of "backpack pack [...]" command.
`,
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("dir", "d", "", "specifies the directory to write into")
}

func unpackRun(cmd *cobra.Command, args []string) {
	// get a file from URL or Path
	p := getAUsablePathOfFile(args[0])
	b, err := pkg.GetPackFromFile(p)
	if err != nil {
		log.Fatalf("Error parsing the backpack: %s", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Checks if a custom directory has been specified, otherwise unpack in
	// the backpack name.
	directory, _ := cmd.Flags().GetString("dir")
	if directory == "" {
		directory = filepath.Join(cwd, fmt.Sprintf("%s-%s", b.Name, b.Version))
		err = os.Mkdir(directory, 0744)
		if err != nil {
			log.Fatalf("Error creating directory: %s", err)
		}
	}

	err = pkg.UnpackInDirectory(&b, directory)
	if err != nil {
		log.Fatalf("Error unpacking: %s", err)
	}
}
