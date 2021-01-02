package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs() []string {
	return []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 306, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 31754, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 291, part2(testInputs()))
	tests.AssertEquals(t, 105, part2([]string{
		"Player 1:",
		"43",
		"19",
		"",
		"Player 2:",
		"2",
		"29",
		"14",
	}))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 35436, part2(inputs))
}
