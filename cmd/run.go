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
	Short: "Starts the jobs of a backpack",
	Long:  `It allows to run different jobs specified in the backpack`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b, err := pkg.GetBackpackFromFile(args[0])
		if err != nil {
			log.Fatalf("Error parsing the backpack: %s", err)
		}

		client, err := conn.NewClinet()
		if err != nil {
			log.Fatalf("Error creating new Nomad Client: %s", err)
		}

		vfPath := cmd.Flag("values").Value.String()
		values, err := pkg.ValuesFromFile(vfPath)
		if err != nil {
			log.Fatalf("Error reading the value file: %s", err)
		}

		// Populate the template into job files 💪
		bts, err := templating.BuildHCL(&b, values)
		if err != nil {
			log.Fatalf("Error building the HCL files: %s", err)
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
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("values", "v", "", "specifies the file to use for values and ensure to populate the Go Templates")
	runCmd.MarkFlagRequired("values")
}
