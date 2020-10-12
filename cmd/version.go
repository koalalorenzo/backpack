package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Specify a version and versionTag
	version        = "unstable"
	versionGitHash = "HEAD"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Run:   versionRun,
	Short: "Shows Backpack version",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Backpack version: %s (%s)\n", version, versionGitHash)
	fmt.Println("More info: https://backpack.qm64.tech")
}
