package helper

import (
	"fmt"
	"os"
)

func OpenFile(file string) (*os.File, error) {
	filePath, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	return filePath, nil
}
