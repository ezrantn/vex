package helper

import (
	"fmt"
	"os"
	"strings"
)

func OpenFile(file string) (*os.File, error) {
	filePath, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	return filePath, nil
}

func ReplaceAllIgnoreCase(text, find, replace string) string {
	lowerText := strings.ToLower(text)
	lowerFind := strings.ToLower(find)
	var result strings.Builder

	i := 0
	for {
		idx := strings.Index(lowerText[i:], lowerFind)
		if idx == -1 {
			// Append the rest of the text
			result.WriteString(text[i:])
			break
		}

		// Append part before the match
		result.WriteString(text[i : i+idx])
		// Append the replacement
		result.WriteString(replace)

		// Move past the matched word
		i += idx + len(find)
	}

	return result.String()
}
