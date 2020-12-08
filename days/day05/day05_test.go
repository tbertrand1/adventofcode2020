package main

import (
	"testing"

	"../../utils/files"
	"../../utils/tests"
)

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 896, part1(inputs))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 659, part2(inputs))
}
