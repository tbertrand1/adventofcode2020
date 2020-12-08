package main

import (
	"testing"

	"../../utils/files"
	"../../utils/tests"
)

func testInputs() []int {
	return []int{1721, 979, 366, 299, 675, 1456}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 514579, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, 542619, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 241861950, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, 32858450, part2(inputs))
}
