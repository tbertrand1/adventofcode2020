package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 3517*3593*2797*3167, part1(inputs))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 1792, part2(inputs))
}
