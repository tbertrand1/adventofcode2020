package main

import (
	"testing"

	files "../../utils/files"
	tests "../../utils/tests"
)

func testInputs() []string {
	return []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part1()", 2, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part1()", 398, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part2()", 1, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part2()", 562, part2(inputs))
}
