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

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:     "plan [path]",
	Aliases: []string{"diff", "dry-run"},
	Args:    cobra.ExactArgs(1),
	Short:   "Plan and check the changes before running a pack",
	Long: `It allows you to plan ahead before running/registering jobs.
It is useful when combined with existing jobs to validate changes. By default
the output shows you a brief summary of changes, but if you want to see the 
full changes being applied for each job in the backpack you can use the option
--verbose. Use the option --diff=false To disable the diff and just check if all 
the jobs can be allocated (dry-run).
`,
	Run: planRun,
}

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.Flags().BoolP("diff", "d", true, "show the differences of changes applied")
	planCmd.Flags().Bool("verbose", false, "show the full changes applied not just top level")
	planCmd.Flags().StringP("values", "v", "", "specifies the file to use for values and ensure to populate the Go Templates")
	planCmd.Flags().BoolP("unpacked", "u", false, "instead of reading from a file, read from a directory")
}

// This is the actual command..
func planRun(cmd *cobra.Command, args []string) {
	b := getPackFromCLIInput(cmd, args)
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

	diffPlan, err := cmd.Flags().GetBool("diff")
	if err != nil {
		log.Fatalf("Error parsing CLI flags (diff): %s", err)
	}

	// Populate the template into job files 💪
	bts, err := templating.BuildHCL(&b, values)
	if err != nil {
		log.Fatalf("Error building the HCL files: %s", err)
	}

	// Prepare a table for the output in a buffer. This is done so that we can
	// have a table after outputting the Plans for each job
	rt, wt, err := os.Pipe()
	if err != nil {
		log.Fatal("Error preparing the output table:", err)
	}

	defer rt.Close()
	w := tabwriter.NewWriter(wt, 3, 0, 4, ' ', 0)
	fmt.Fprintf(w, "Recap of Job Plans in \"%s\" backpack:\n", b.Name)
	if !diffPlan {
		fmt.Fprintln(w, "File Name\tCheck Index\tDry Run status\tPlan warnings")
	} else {
		fmt.Fprintln(w, "File Name\tCheck Index\tDiff type\tDry run status\tWarnings")
	}
	// For each job file perform the plan! 🚀
	for name, hcl := range bts {
		job, err := client.GetJobFromCode(string(hcl))
		if err != nil {
			log.Fatalf("Error obtaining job %s: %s", name, err)
		}

		// always show the diff for plan
		p, err := client.Plan(job, diffPlan)
		if err != nil {
			log.Fatalf("Error running %s: %s", name, err)
		}

		// Makes it clear that there are no Warnings
		if p.Warnings == "" {
			p.Warnings = "None"
		}

		// Check if dry-run was successfull. This helps understanding if the tasks
		// are allocated properly on the nodes or not
		dryRunStatus := "Success"
		if len(p.FailedTGAllocs) != 0 {
			dryRunStatus = fmt.Sprintf("%d allocations failed", len(p.FailedTGAllocs))
		}

		// If we disabled diff just populate the table and skip
		if !diffPlan {
			fmt.Fprintf(w, "%s\t%d\t%s\t%s\t\n", name, p.JobModifyIndex, dryRunStatus, p.Warnings)
			continue
		}

		// Write in the table in the buffer output
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t\n", name, p.JobModifyIndex, p.Diff.Type, dryRunStatus, p.Warnings)

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
