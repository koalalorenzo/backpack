package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/bundle"
)

// unpackCmd represents the unpack command
var unpackCmd = &cobra.Command{
	Use:   "unpack [file.backpack]",
	Short: "Opens a Backpack file to explore content",
	Long: `Explodes the backpack inside a directory. This is useful to edit a 
Backpack, inspecting it or seeing default values.

The Backpack includes:
- backpack.yaml (containing metadata and default values)
- *.nomad (representing the various go templates of Nomad Jobs)
`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		b, err := bundle.GetBundleFromFile(args[0])
		if err != nil {
			log.Fatalf("Error parsing the bundle: %s", err)
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

		err = bundle.UnpackBundleInDirectory(&b, directory)
		if err != nil {
			log.Fatalf("Error unpacking: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unpackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	unpackCmd.Flags().StringP("dir", "d", "", "specifies the directory to write into")
}
