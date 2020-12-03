package main

import (
	"testing"

	utils "../../utils"
)

func testInputs() []string {
	return []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	utils.AssertEquals(t, "part1()", 7, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := utils.ReadLinesOfFile(filename)
	utils.AssertEquals(t, "part1()", 284, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	utils.AssertEquals(t, "part2()", 336, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := utils.ReadLinesOfFile(filename)
	utils.AssertEquals(t, "part2()", 3510149120, part2(inputs))
}
