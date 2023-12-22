package file

import (
	"bufio"
	"os"
)

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line as needed (you can replace this with your logic)
		result += line + "\n"
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}
