package main

import (
	"testing"

	utils "../../utils"
)

func TestPart1(t *testing.T) {
	inputs := utils.ReadLinesOfFile(filename)
	utils.AssertEquals(t, "part1()", 896, part1(inputs))
}

func TestPart2(t *testing.T) {
	inputs := utils.ReadLinesOfFile(filename)
	utils.AssertEquals(t, "part2()", 659, part2(inputs))
}
