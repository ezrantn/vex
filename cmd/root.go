package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vex",
	Short: "Vex is a powerful text processing tool.",
	Long:  "Vex is a command-line tool for advanced text processing with features like regex, replace, filtering, and more.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		loadCmd,
		saveCmd,
		replaceCmd,
		versionCmd,
		filterCmd,
		regexCmd,
		countCmd,
		formatCmd,
	)
}
