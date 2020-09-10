package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/conn"
	"gitlab.com/qm64/backpack/pkg"
	"gitlab.com/qm64/backpack/templating"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [path]",
	Args:  cobra.ExactArgs(1),
	Short: "Starts the jobs of a backpack",
	Long: `It allows to run different jobs specified in the backpack.
It accepts one argument that is the path of the file, but if the option 
--unpacked (or -u is) passed it consider the first argument as the path of an
unpacked backpack directory that will be used instead of a file.
`,
	Run: runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("values", "v", "", "specifies the file to use for values and ensure to populate the Go Templates")
	runCmd.Flags().BoolP("unpacked", "u", false, "instead of reading from a file, read from a directory")
	runCmd.Flags().Bool("debug", false, "prints the jobs on stdout instead of sending them to nomad")
}

// This is the actual command..
func runRun(cmd *cobra.Command, args []string) {
	b := pkg.Backpack{}
	var err error

	readFromDir := cmd.Flag("unpacked").Value.String()
	if readFromDir == "false" {
		// get a file from URL or Path
		p := getAUsablePathOfFile(args[0])

		b, err = pkg.GetBackpackFromFile(p)
		if err != nil {
			log.Fatalf("Error parsing the backpack: %s", err)
		}
	} else {
		// If we have to read from directory instead args[0] is a path
		d, err := pkg.GetBackpackFromDirectory(args[0])
		b = *d
		if err != nil {
			log.Fatalf("Error parsing the unpacked backpack: %s", err)
		}
	}

	client, err := conn.NewClinet()
	if err != nil {
		log.Fatalf("Error creating new Nomad Client: %s", err)
	}

	vfPath := cmd.Flag("values").Value.String()
	values := pkg.ValuesType{}
	if vfPath != "" {
		values, err = pkg.ValuesFromFile(vfPath)
		if err != nil {
			log.Fatalf("Error reading the value file: %s", err)
		}
	}

	// Populate the template into job files 💪
	bts, err := templating.BuildHCL(&b, values)
	if err != nil {
		log.Fatalf("Error building the HCL files: %s", err)
	}

	debugFlag := cmd.Flag("debug").Value.String()
	if debugFlag == "true" {
		for name, hcl := range bts {
			log.Printf("File: %s\n", name)
			fmt.Println(string(hcl))
		}
		return
	}

	// For each job file run it! 🚀
	// then store the job ID in the backpack to show it afterwards.
	jIDs := map[string]string{}
	for name, hcl := range bts {
		jID, err := client.Run(string(hcl))
		if err != nil {
			log.Fatalf("Error running %s: %s", name, err)
		}
		jIDs[name] = jID
	}
	b.JobsEvalIDs = jIDs

	// Prepare a table for the output
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
	fmt.Fprintln(w, "File Name\tJob ID\t")
	for n, jID := range b.JobsEvalIDs {
		fmt.Fprintf(w, "%s\t%s\t\n", n, jID)
	}
	w.Flush()
}
