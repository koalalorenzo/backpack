package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "backpack",
		Short: "Package Manager for Hashicorp Nomad",
		Long: `Backpack allows you to deploy a suite of jobs with a templating system.
Please read more at https://backpack.qm64.tech/

Copyright © 2020 Lorenzo Setale https://setale.me
This program and its source code is licensed under the terms of the
GNU Lesser General Public License v3 (LGPLv3). Please refer to the source code 
for the full license text

Happy Backpacking! 🎒😀 
`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
