package utils

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

// AssertEquals verify that a test is in error
func AssertEquals(t *testing.T, context string, expected int, actual int) {
	if expected != actual {
		t.Errorf("%s test fail; expected %d; actual %d", context, expected, actual)
	}
}

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

// Find returns the index of a int in an array of ints
func Find(values []int, value int) (int, bool) {
	for i, val := range values {
		if val == value {
			return i, true
		}
	}
	return -1, false
}
