package contain

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Contains(filePath string, substring string) (bool, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return false, err
		}
		if err == io.EOF && line == "" {
			break
		}
		if err == io.EOF && line != "" {
			if contains := strings.Contains(line, substring); contains {
				return true, nil
			}
			break
		}
		if contains := strings.Contains(line, substring); contains {
			return true, nil
		}
	}

	return false, nil
}
