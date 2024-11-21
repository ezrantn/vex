package cmd

import (
	"fmt"
	"os"

	"github.com/ezrantn/vex/dsl"
	"github.com/spf13/cobra"
)

var loadedContent string

var loadCmd = &cobra.Command{
	Use:   "load [:=input.txt]",
	Short: "Load a text file into memory",
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

		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", file, err)
			return
		}

		loadedContent = string(content)
		fmt.Printf("Successfully loaded file '%s' into memory.\n", file)
	},
}
