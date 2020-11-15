package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/conn"
	"gitlab.com/qm64/backpack/templating"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run [path]",
	Aliases: []string{"start", "install"},
	Args:    cobra.ExactArgs(1),
	Short:   "Starts the jobs of a pack",
	Long: `It allows to run different jobs specified in the pack.
It accepts one argument that is the path or URL of the file, but if the option 
--unpacked (or -u is) passed it consider the first argument as the path of an
unpacked directory that will be used instead of a file.
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
	b := getPackFromCLIInput(cmd, args)
	var err error

	client, err := conn.NewClient()
	if err != nil {
		log.Fatalf("Error creating new Nomad Client: %s", err)
	}

	values := getValuesFromCLIInput(cmd)

	// Populate the template into job files 💪
	bts, err := templating.BuildHCL(&b, values)
	if err != nil {
		log.Fatalf("Error building the HCL files: %s", err)
	}

	debugFlag, _ := cmd.Flags().GetBool("debug")
	if debugFlag {
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
		job, err := client.GetJobFromCode(string(hcl))
		if err != nil {
			log.Fatalf("Error obtaining job %s: %s", name, err)
		}

		// Run = Register the job
		jr, err := client.Run(job)
		if err != nil {
			log.Fatalf("Error running %s: %s", name, err)
		}

		jIDs[name] = jr.EvalID
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
