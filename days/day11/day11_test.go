package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs() []string {
	return []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 37, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 2472, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 26, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 2197, part2(inputs))
}
