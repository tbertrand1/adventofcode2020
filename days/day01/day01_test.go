package main

import (
	"testing"

	utils "../../utils"
)

func TestPart1WithTestValues(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	got := part1(values)
	expected := 514579
	if got != expected {
		t.Errorf("Part1() = %d; expected %d", got, expected)
	}
}

func TestPart1(t *testing.T) {
	values := utils.ReadLinesAsIntsOfFile("day01.txt")
	got := part1(values)
	expected := 542619
	if got != expected {
		t.Errorf("Part1() = %d; expected %d", got, expected)
	}
}

func TestPart2WithTestValues(t *testing.T) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	got := part2(values)
	expected := 241861950
	if got != expected {
		t.Errorf("Part2() = %d; expected %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	values := utils.ReadLinesAsIntsOfFile("day01.txt")
	got := part2(values)
	expected := 32858450
	if got != expected {
		t.Errorf("Part2() = %d; expected %d", got, expected)
	}
}
