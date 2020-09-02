package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/bundle"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack [path]",
	Short: "Build a Backpack file from a directory/template",
	Long: `Generate a Backpack file from a directory containing backpack.yaml and
the nomad job files written as go templates. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := bundle.GetBundleFromDirectory(args[0])
		if err != nil {
			log.Fatal(err)
		}
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		writeTo := filepath.Join(cwd, fmt.Sprintf("%s.nomadbackpack", b.Name))

		fileFlag := cmd.Flag("file").Value.String()
		if fileFlag != "" {
			writeTo = fileFlag
		}

		bundle.WriteBundleToFile(*b, writeTo)
	},
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().String("file", "", "path of the file to write into")
}
