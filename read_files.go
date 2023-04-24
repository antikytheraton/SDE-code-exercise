package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: err: %w", err)
	}
	defer file.Close()

	var out []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		out = append(out, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: err: %w", err)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("error empty file")
	}
	return out, nil
}
