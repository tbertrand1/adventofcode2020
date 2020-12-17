package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs() []string {
	return []string{
		".#.",
		"..#",
		"###",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 112, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 426, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 848, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 1892, part2(inputs))
}
