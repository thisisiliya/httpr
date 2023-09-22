package extract

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) []string {

	var contents []string
	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if text := scanner.Text(); text != "" {

			contents = append(contents, scanner.Text())
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}

	file.Close()
	return contents
}
