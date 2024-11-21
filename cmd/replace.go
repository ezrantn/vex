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
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Expected syntax: find:replace=file\n")
			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", file, err)
			return
		}

		if input != string(content) {
			fmt.Fprintf(os.Stderr, "We can't find the word '%s' in the file '%s'\n", find, file)
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
			fmt.Fprintf(os.Stderr, "Error writing to file '%s': %v\n", file, err)
			return
		}
		fmt.Printf("Successfully replaced '%s' with '%s' in file '%s'\n", find, replace, file)
	},
}

func init() {
	replaceCmd.Flags().BoolVarP(&caseInsensitive, "ignore-case", "i", false, "Perform case-insensitive replacements")
}
