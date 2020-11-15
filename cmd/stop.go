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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:     "stop [path]",
	Aliases: []string{"uninstall", "delete"},
	Args:    cobra.ExactArgs(1),
	Short:   "Stop all the jobs in a pack",
	Long: `This command will stop all the jobs available in a pack. By default it
mimics nomad CLI, that keeps the job as "dead" until the garbage collector
deletes the job. If you want to delete the jobs entirely and lose all the setup
you can pass the option --purge (or -p).
`,
	Run: stopRun,
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().BoolP("purge", "p", false, "Delete the jobs, without waiting for the GC to fully delete them")
	stopCmd.Flags().BoolP("unpacked", "u", false, "instead of reading from a file, read from a directory")
}

// This is the actual command..
func stopRun(cmd *cobra.Command, args []string) {
	b := getPackFromCLIInput(cmd, args)
	var err error

	client, err := conn.NewClient()
	if err != nil {
		log.Fatalf("Error creating new Nomad Client: %s", err)
	}

	purgeJob, err := cmd.Flags().GetBool("purge")
	if err != nil {
		log.Fatalf("Error parsing CLI flags (purge): %s", err)
	}

	// Populate the template into job files 💪
	bts, err := templating.BuildHCL(&b, pkg.ValuesType{})
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
	fmt.Fprintf(w, "Recap of stopping jobs in \"%s\" backpack:\n", b.Name)
	fmt.Fprintln(w, "File Name\tJob ID\tStatus\tError if any")
	// For each job file go and stop 🚀
	for name, hcl := range bts {
		job, err := client.GetJobFromCode(string(hcl))
		if err != nil {
			log.Fatalf("Error obtaining job %s: %s", name, err)
		}

		// Stop the job
		_, err = client.Stop(*job.ID, purgeJob)
		if err != nil {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", name, *job.ID, *job.Status, err)
			continue
		}

		if !purgeJob {
			jobAfter, err := client.GetJobStatus(*job.ID)
			if err != nil {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", name, *job.ID, *job.Status, err)
				continue
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", name, *jobAfter.ID, *jobAfter.Status, "")
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", name, *job.ID, "purged", "")
		}

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
