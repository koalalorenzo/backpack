package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/conn"
	"gitlab.com/qm64/backpack/pkg"
	"gitlab.com/qm64/backpack/templating"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status [path]",
	Aliases: []string{"state", "alloc"},
	Args:    cobra.ExactArgs(1),
	Short:   "Check the status of all the jobs in a pack",
	Long: `Run this command to know the status of all the jobs in a pack. It will
check and provide useful information. By default it shows the allocations that
are running, or the first allocation available. If you want to see all the
previous allocation you can use the option --all (or -a).

`,
	Run: statusRun,
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolP("all", "a", false, "show all allocations")
	statusCmd.Flags().BoolP("unpacked", "u", false, "instead of reading from a file, read from a directory")
}

var showAllocationsWithStatus = map[string]bool{
	api.AllocClientStatusRunning: true,
	api.AllocClientStatusPending: true,
}

// This is the actual command..
func statusRun(cmd *cobra.Command, args []string) {
	b := getPackFromCLIInput(cmd, args)
	var err error

	client, err := conn.NewClient()
	if err != nil {
		log.Fatalf("Error creating new Nomad Client: %s", err)
	}

	showAllAlloc, err := cmd.Flags().GetBool("all")
	if err != nil {
		log.Fatalf("Error parsing CLI flags (all): %s", err)
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
	fmt.Fprintf(w, "Status of the jobs' allocations from \"%s\" backpack:\n", b.Name)
	fmt.Fprintln(w, "Job ID\tAlloc ID\tStatus/Desired\tModified At\tError")

	for name, hcl := range bts {
		job, err := client.GetJobFromCode(string(hcl))
		if err != nil {
			log.Fatalf("Error obtaining job %s: %s", name, err)
		}

		jobResult, err := client.GetJobStatus(*job.ID)
		if err != nil {
			fmt.Fprintf(w, "%s\t\t%s\t\t%s\n", *job.ID, *job.Status, err)
			continue
		}

		allocations, err := client.GetJobAllocations(*job.ID)
		if err != nil {
			fmt.Fprintf(w, "%s\t\t%s\t\t%s\n", *job.ID, *jobResult.Status, err)
			continue
		}

		// fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", *jobResult.ID, "(check allocations)", *jobResult.Status, "", "")
		for i, alloc := range allocations {
			_, ok := showAllocationsWithStatus[alloc.ClientStatus]
			if !showAllAlloc && len(allocations) > 1 && i != 0 && !ok {
				continue
			}
			lt := time.Unix(0, alloc.ModifyTime).Format(time.RFC3339)
			allocID := sanitizeUUIDPrefix(alloc.ID)
			fmt.Fprintf(w, "%s\t%s\t%s/%s\t%s\t\n", *jobResult.ID, allocID, alloc.ClientStatus, alloc.DesiredStatus, lt)
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
