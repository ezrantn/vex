package cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ezrantn/vex/dsl"
	"github.com/ezrantn/vex/helper"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter [word=file]",
	Short: "Filter a word in a file",
	Long:  "Filtering a word based on your input, then display verbose information regarding that file",
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		parser := dsl.NewParser(input)
		word, file, err := parser.ParseFilterCommand()
		if err != nil {
			fmt.Printf("Error parsing command: %v\n", err)
			return
		}

		matches := []string{}
		lineNumbers := []int{}
		totalMatches := 0

		fileReader, err := helper.OpenFile(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		scanner := bufio.NewScanner(fileReader)
		lineNumber := 1

		for scanner.Scan() {
			line := scanner.Text()
			count := strings.Count(line, word)

			if count > 0 {
				matches = append(matches, line)
				lineNumbers = append(lineNumbers, lineNumber)
				totalMatches += count
			}

			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading file: %v", err)
			return
		}

		fmt.Printf("Pattern: \"%s\"\n", word)
		fmt.Printf("File: %s\n", file)
		fmt.Printf("Total Matches: %d\n\n", totalMatches)

		for i, match := range matches {
			fmt.Printf("Match %d: Line %d: %s\n", i+1, lineNumbers[i], match)
		}
	},
}
