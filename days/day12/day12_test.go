package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs() []string {
	return []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 25, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 582, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 286, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 52069, part2(inputs))
}
