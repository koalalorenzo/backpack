package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/conn"
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
	planCmd.Flags().BoolP("diff", "d", false, "show the differences of changes applied")
	planCmd.Flags().Bool("verbose", false, "show the full changes applied not just top level")
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

	values := getValuesFromCLIInput(cmd)

	verbosePlan, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		log.Fatalf("Error parsing CLI flags (verbose): %s", err)
	}

	debugFlag, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatalf("Error parsing CLI flags (debug): %s", err)
	}

	// Populate the template into job files 💪
	bts, err := templating.BuildHCL(&b, values)
	if err != nil {
		log.Fatalf("Error building the HCL files: %s", err)
	}

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
	if err != nil {
		log.Fatal("Error preparing the output table:", err)
	}

	defer rt.Close()
	w := tabwriter.NewWriter(wt, 3, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Recap of Job Plans")
	fmt.Fprintln(w, "File Name\tCheck Index\tDiff Type\tPlan warnings")

	// For each job file perform the plan! 🚀
	for name, hcl := range bts {
		job, err := client.GetJob(string(hcl))
		if err != nil {
			log.Fatalf("Error obtaining job %s: %s", name, err)
		}

		// always show the diff for plan
		p, err := client.Plan(job, true)
		if err != nil {
			log.Fatalf("Error running %s: %s", name, err)
		}
		// Write in the table in the buffer output
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t\n", name, p.JobModifyIndex, p.Diff.Type, p.Warnings)

		// Write in the output the diff from the previous output
		fmt.Printf("Plan for job %s\n", name)
		fmt.Printf("%s Job: \"%s\"\n", getDiffSimbol(p.Diff.Type), *job.ID)
		for _, field := range p.Diff.Fields {
			printFieldsDiff(field, verbosePlan, 1)
		}
		for _, object := range p.Diff.Objects {
			printObjectDiff(object, verbosePlan, 1)
		}
		// Print the task groups changes
		for _, tg := range p.Diff.TaskGroups {
			// If no changes happed during the plan, continue to the next Task Group
			if tg.Type == "None" {
				continue
			}

			// Build indentation level
			indStr := strings.Repeat(indentationStyle, 1)

			fmt.Printf("%s%s Task Group: \"%s\"\n", indStr, getDiffSimbol(tg.Type), tg.Name)

			for _, field := range tg.Fields {
				printFieldsDiff(field, verbosePlan, 2)
			}

			for _, obj := range tg.Objects {
				printObjectDiff(obj, verbosePlan, 2)
			}

			for _, task := range tg.Tasks {
				// Build indentation level
				indStr := strings.Repeat(indentationStyle, 2)

				ann := strings.Join(task.Annotations, ", ")
				fmt.Printf("%s%s Task: \"%s\" (%s)\n", indStr, getDiffSimbol(task.Type), task.Name, ann)

				for _, field := range task.Fields {
					printFieldsDiff(field, verbosePlan, 3)
				}

				for _, obj := range task.Objects {
					printObjectDiff(obj, verbosePlan, 3)
				}

			}
		}
		fmt.Println()
	}
	// Flushes all the table output after all the plans output.
	w.Flush()
	wt.Close()
	output, err := ioutil.ReadAll(rt)
	if err != nil {
		log.Fatal("Error reading the output table after operation completed:", err)
	}
	os.Stdout.Write(output)
}
