package main

import (
	"testing"

	files "../../utils/files"
	tests "../../utils/tests"
)

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part1()", 896, part1(inputs))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part2()", 659, part2(inputs))
}
