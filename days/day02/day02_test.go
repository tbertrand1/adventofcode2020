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
	got := part1(testInputs())
	expected := 2
	if got != expected {
		t.Errorf("Part1() = %d; expected %d", got, expected)
	}
}

func TestPart1(t *testing.T) {
	inputs := utils.ReadLinesOfFile("day02.txt")
	got := part1(inputs)
	expected := 398
	if got != expected {
		t.Errorf("Part1() = %d; expected %d", got, expected)
	}
}

func TestPart2WithTestInputs(t *testing.T) {
	got := part2(testInputs())
	expected := 1
	if got != expected {
		t.Errorf("Part2() = %d; expected %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	inputs := utils.ReadLinesOfFile("day02.txt")
	got := part2(inputs)
	expected := 562
	if got != expected {
		t.Errorf("Part2() = %d; expected %d", got, expected)
	}
}
