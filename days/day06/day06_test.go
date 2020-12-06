package main

import (
	"testing"

	files "../../utils/files"
	tests "../../utils/tests"
)

func testInputs() []string {
	return []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part1()", 11, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part1()", 6680, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part2()", 6, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part2()", 3117, part2(inputs))
}
