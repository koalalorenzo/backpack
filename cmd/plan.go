package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/conn"
	"gitlab.com/qm64/backpack/pkg"
	"gitlab.com/qm64/backpack/templating"
)

// planCmd represents the run command
var planCmd = &cobra.Command{
	Use:   "plan [path]",
	Args:  cobra.ExactArgs(1),
	Short: "Plan the jobs of a backpack",
	Long:  ``,
	Run:   planRun,
}

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.Flags().BoolP("diff", "d", true, "calculate and show the differences and changes")
	planCmd.Flags().StringP("values", "v", "", "specifies the file to use for values and ensure to populate the Go Templates")
	planCmd.Flags().BoolP("unpacked", "u", false, "instead of reading from a file, read from a directory")
	planCmd.Flags().Bool("debug", false, "prints the jobs on stdout instead of sending them to nomad")
}

// This is the actual command..
func planRun(cmd *cobra.Command, args []string) {
	b := getBackpackFromCLIInput(cmd, args)
	var err error

	client, err := conn.NewClient()
	if err != nil {
		log.Fatalf("Error creating new Nomad Client: %s", err)
	}

	vfPath, _ := cmd.Flags().GetString("values")
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

	debugFlag, _ := cmd.Flags().GetBool("debug")
	if debugFlag {
		for name, hcl := range bts {
			log.Printf("File: %s\n", name)
			fmt.Println(string(hcl))
		}
		return
	}

	// Prepare a table for the output in a buffer. This is done so that we can
	// have a table after outputting the Plans for each job
	rt, wt, err := os.Pipe()
	w := tabwriter.NewWriter(wt, 3, 0, 2, ' ', 0)
	fmt.Fprintln(w, "File Name\tCheck Index\tDiff Type\tPlan warnings")
	showPlanDiff, _ := cmd.Flags().GetBool("diff")

	// For each job file perform the plan! 🚀
	for name, hcl := range bts {
		p, err := client.Plan(string(hcl), showPlanDiff)
		if err != nil {
			log.Fatalf("Error running %s: %s", name, err)
		}
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t\n", name, p.JobModifyIndex, p.Diff.Type, p.Warnings)

	}

	// Flushes all the table output
	w.Flush()
	wt.Close()
	output, err := ioutil.ReadAll(rt)
	if err != nil {
		log.Fatal("Error reading the output table after operation completed:", err)
	}
	os.Stdout.Write(output)
}
