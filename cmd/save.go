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
			fmt.Printf("Error parsing command: %v\n", err)
			return
		}

		if err := os.WriteFile(file, []byte(loadedContent), 0644); err != nil {
			fmt.Printf("Failed to save a file: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("File saved successfully.")
	},
}
