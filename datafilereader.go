package datafilereader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadLinesFromFile reads the given file line by line, and returns a slice of strings
// containing the lines read, or an error is once occurred.
func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

// ReadInts parses integers from a slice of strings, returning a slice of ints,
// or an error is one occurred.
func ReadInts(filename string) ([]int, error) {

	ints := make([]int, 0)
	lines, err := ReadLinesFromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read lines from file: %v", err)
	}

	for _, line := range lines {
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return ints, fmt.Errorf("unable to parse int: %v", err)
		}
		ints = append(ints, int(i))
	}

	return ints, nil
}
