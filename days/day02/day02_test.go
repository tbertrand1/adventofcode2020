package main

import (
	"testing"

	utils "../../utils"
)

func testInputs() []string {
	return []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	utils.AssertEquals(t, "part1()", 2, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := utils.ReadLinesOfFile(filename)
	utils.AssertEquals(t, "part1()", 398, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	utils.AssertEquals(t, "part2()", 1, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := utils.ReadLinesOfFile(filename)
	utils.AssertEquals(t, "part2()", 562, part2(inputs))
}
