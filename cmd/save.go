package cmd

import (
	"fmt"
	"os"

	"github.com/ezrantn/vex/dsl"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [:=output.txt]",
	Short: "Save the current content to a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		parser := dsl.NewParser(input)
		file, err := parser.ParseFileCommand()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing command: %v\n", err)
			fmt.Fprintf(os.Stderr, "Expected syntax: :=file_name\n")
			return
		}

		if err := os.WriteFile(file, []byte(loadedContent), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "✖ Failed to save content to file '%s': %v\n", file, err)
			os.Exit(1)
		}

		fmt.Printf("✔ Successfully saved content to file '%s'.\n", file)
	},
}
