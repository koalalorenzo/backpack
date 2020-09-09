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
	Use:   "pack [path]",
	Short: "Build a Backpack file from a directory/template",
	Long: `Generate a Backpack file from a directory containing backpack.yaml and
the nomad job files written as go templates

The directory should have these files:
- backpack.yaml (containing metadata)
- values.yaml (containing the default values for the templates)
- *.nomad (representing the various go templates of Nomad Jobs)
- *.md (useful documentation)

This command performs the opposite of "unpack" command
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

	b, err := pkg.GetBackpackFromDirectory(args[0])
	if err != nil {
		log.Fatalf("Error generating the backpack from the directory: %s", err)
	}
	writeTo := filepath.Join(cwd, fmt.Sprintf("%s-%s.backpack", b.Name, b.Version))

	fileFlag := cmd.Flag("file").Value.String()
	if fileFlag != "" {
		writeTo = fileFlag
	}

	err = pkg.WriteBackpackToFile(*b, writeTo)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}
