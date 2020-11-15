package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/pkg"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:     "pack [path]",
	Aliases: []string{"package"},
	Short:   "Build a Backpack file (pack) from a directory/template",
	Long: `Generate a Backpack file (pack) from a directory containing the various
jobs, metadata and documentation.

The directory should have these files:
- backpack.yaml (containing metadata)
- values.yaml (containing the default values for the templates)
- *.nomad (representing the various go templates of Nomad Jobs)
- *.md (useful documentation)

This command performs the opposite of "backpack unpack [...]" command
`,
	Args: cobra.ExactArgs(1),
	Run:  packRun,
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("file", "f", "", "path of the file to create and write into")
}

func packRun(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	b, err := pkg.GetPackFromDirectory(args[0])
	if err != nil {
		log.Fatalf("Error generating the backpack from the directory: %s", err)
	}
	writeTo := filepath.Join(cwd, fmt.Sprintf("%s-%s.backpack", b.Name, b.Version))

	fileFlag, _ := cmd.Flags().GetString("file")
	if fileFlag != "" {
		writeTo = fileFlag
	}

	err = pkg.WritePackToFile(*b, writeTo)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}
