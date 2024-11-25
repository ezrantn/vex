package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Vex",
	Long:  `All software has versions. This is Vex's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.2.2")
	},
}
