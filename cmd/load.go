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
			fmt.Printf("Error parsing command: %v\n", err)
			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		loadedContent = string(content)
		fmt.Printf("File '%s' successfully loaded in the system\n", file)
	},
}
