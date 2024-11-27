package cmd

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"

	"github.com/ezrantn/vex/dsl"
	"github.com/ezrantn/vex/helper"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count [word=file]",
	Short: "Count the occurences of a word",
	Long:  "Count the occurences of a word that pops in your file",
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		parser := dsl.NewParser(input)
		pattern, file, err := parser.ParseCountCommand()

		if err != nil {
			fmt.Printf("Error parsing command: %v\n", err)
			return
		}

		fileReader, err := helper.OpenFile(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		defer fileReader.Close()

		scanner := bufio.NewScanner(fileReader)

		totalOccurrences := 0
		lineOccurrences := make(map[int]int)

		searchPattern := strings.TrimSpace(pattern)

		lineNumber := 0
		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()

			lineOccurrencesCount := countWordOccurrences(line, searchPattern)

			if lineOccurrencesCount > 0 {
				totalOccurrences += lineOccurrencesCount
				lineOccurrences[lineNumber] = lineOccurrencesCount
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		if totalOccurrences > 0 {
			fmt.Printf("Total occurrences of '%s': %d\n", searchPattern, totalOccurrences)
		} else {
			fmt.Printf("No occurrences of '%s' found in the file.\n", searchPattern)
		}
	},
}

func countWordOccurrences(line, word string) int {
	line = strings.ToLower(line)
	word = strings.ToLower(word)

	words := splitIntoWords(line)

	count := 0
	for _, w := range words {
		if w == word {
			count++
		}
	}

	return count
}

func splitIntoWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
