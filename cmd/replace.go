package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ezrantn/vex/dsl"
	"github.com/ezrantn/vex/helper"
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
		findList, replaceList, file, err := parser.ParseReplaceCommand()
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

		text := string(content)
		found := false

		for _, find := range findList {
			if caseInsensitive {
				if strings.Contains(strings.ToLower(text), strings.ToLower(find)) {
					found = true
					break
				}
			} else {
				if strings.Contains(text, find) {
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Fprintf(os.Stderr, "None of the specified terms in 'find' exist in the file '%s'\n", file)
			return
		}

		for i, find := range findList {
			replace := replaceList[i]
			if caseInsensitive {
				text = helper.ReplaceAllIgnoreCase(text, find, replace)
			} else {
				text = strings.ReplaceAll(text, find, replace)
			}
		}

		err = os.WriteFile(file, []byte(text), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file '%s': %v\n", file, err)
			return
		}
		fmt.Printf("Successfully replaced %d item(s) in file '%s'\n", len(findList), file)
	},
}

func init() {
	replaceCmd.Flags().BoolVarP(&caseInsensitive, "ignore-case", "i", false, "Perform case-insensitive replacements")
}
