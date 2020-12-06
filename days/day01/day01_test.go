package main

import (
	"testing"

	files "../../utils/files"
	tests "../../utils/tests"
)

func testInputs() []int {
	return []int{1721, 979, 366, 299, 675, 1456}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part1()", 514579, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, "part1()", 542619, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part2()", 241861950, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, "part2()", 32858450, part2(inputs))
}
