package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func TestPart1WithTestInputs(t *testing.T) {
	testInputs := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
	tests.AssertEquals(t, 71, part1(testInputs))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 27870, part1(inputs))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 3173135507987, part2(inputs))
}
