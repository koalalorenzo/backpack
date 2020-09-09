package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/pkg"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new unpacked backpack in a directory",
	Long: `Use this command to create the directory structure and some basic files
to modify for new workloads and backpacks. If no value is passed by -d or --dir
a new directory with the name and version of the backapck will be created.

The directory that will be created will have these files:
- backpack.yaml (containing metadata)
- values.yaml (containing the default values for the templates)
- main.nomad (representing a go templates of Nomad Jobs written as HCL)
- README.md (useful documentation)
`,
	Run: createRun,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("dir", "d", "", "specifies the directory to write into")
}

func createRun(cmd *cobra.Command, args []string) {
	name := args[0]
	b := pkg.Backpack{
		Name:         name,
		Version:      "0.1.0",
		Dependencies: map[string]string{"TODO": "http://backpack.qm64.com/example.backpack"},

		DefaultValues: []byte(`# Write the default values for your backpack here
# It is strongly suggested to use inline documentation for quick explanations
# otherwise the *.md files can be used for longer documentation.
datacenters:
	- dc1
	- dc2
`),

		Templates: pkg.FilesMapType{
			"main.nomad": []byte(fmt.Sprintf(`job "%s" {
	datacenters = [{{range .datacenters}} "{{ . }}", {{else}} "dc1" {{end}}]
	type = "service"

	task "%s_nginx" {
		driver = "docker"
		config {
			image = "nginx:alpine"
			ports = ["http", "https"]
		}
		network {
			port "http" {
				static = 80
			}
			port "https" {
				static = 443
			}
		}
}`, name, name)),
		},

		Documentation: pkg.FilesMapType{
			"README.md": []byte(fmt.Sprintf(`# How to deploy %s
Write here a longer documentation full with links and examples or things to 
know when upgrading/downgrading, values combination, etc, etc! 

Have fun! 😄 
`, name)),
		},
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Checks if a custom directory has been specified, otherwise unpack in
	// the backpack name.
	directory := cmd.Flag("dir").Value.String()
	showPath := directory
	if directory == "" {
		showPath = fmt.Sprintf("%s-%s", b.Name, b.Version)
		directory = filepath.Join(cwd, showPath)
		err = os.Mkdir(directory, 0744)
		if err != nil {
			log.Fatalf("Error creating directory: %s", err)
		}
	}

	err = pkg.UnpackBackpackInDirectory(&b, directory)
	if err != nil {
		log.Fatalf("Error unpacking: %s", err)
	}

	fmt.Printf(`Congratulations! Feel free to modify the files written in %s

Your next step will be, creating a backpack file:
  $ backpack pack %s

You can then run the backpack in your nomad local testing:
  $ backpack run %s.backpack 

Happy packing and sharing!
`, directory, showPath, showPath)
}
