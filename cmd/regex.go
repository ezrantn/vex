package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/ezrantn/vex/dsl"
	"github.com/spf13/cobra"
)

var regexCmd = &cobra.Command{
	Use:   "match [regex]=[input]",
	Short: "Pattern matching using regex",
	Long:  "Perform regex searches within text",
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		parser := dsl.NewParser(input)
		regex, file, err := parser.ParseRegexCommand()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Expected syntax: [regex]=file\n")
			return
		}

		re, err := regexp.Compile(regex)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid regex pattern - %v\n", err)
			return
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Cannot open file - %v\n", err)
			return
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		matchedLines := []string{}
		lineNumber := 0

		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			if re.MatchString(line) {
				matchedLines = append(matchedLines, fmt.Sprintf("Line %d: %s", lineNumber, line))
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}

		if len(matchedLines) > 0 {
			fmt.Printf("Found %d matching line(s):\n\n", len(matchedLines))
			for _, matchedLine := range matchedLines {
				fmt.Println(matchedLine)
			}
		} else {
			fmt.Printf("No lines matched the pattern '%s' in file '%s'\n", regex, file)
		}
	},
}
