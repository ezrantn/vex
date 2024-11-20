package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ezrantn/vex/dsl"
	"github.com/spf13/cobra"
)

var caseInsensitive bool

var replaceCmd = &cobra.Command{
	Use:   "replace [find:replace=file]",
	Short: "Find and replace text in a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		parser := dsl.NewParser(input)
		find, replace, file, err := parser.ParseReplaceCommand()
		if err != nil {
			fmt.Printf("Error parsing command: %v\n", err)
			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		text := string(content)

		if caseInsensitive {
			text = strings.ReplaceAll(
				strings.ToLower(text),
				strings.ToLower(find),
				replace,
			)
		} else {
			text = strings.ReplaceAll(text, find, replace)
		}

		err = os.WriteFile(file, []byte(text), 0644)
		if err != nil {
			fmt.Printf("Error writing to a file: %v\n", err)
			return
		}
		fmt.Printf("Successfully replaced '%s' with '%s' in file '%s'\n", find, replace, file)
	},
}

func init() {
	replaceCmd.Flags().BoolVarP(&caseInsensitive, "ignore-case", "i", false, "Perform case-insensitive replacements")
}
