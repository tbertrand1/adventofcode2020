package files

import (
	"bufio"
	"os"
	"strconv"
)

// ReadLinesOfFile reads a file and returns all the lines
func ReadLinesOfFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()
	return lines
}

// ReadLinesAsIntsOfFile reads a file and returns all the lines as ints
func ReadLinesAsIntsOfFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []int
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, value)
	}

	file.Close()
	return lines
}
