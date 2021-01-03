package main

import (
	"../../utils/tests"
	"testing"
)

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 67384529, part1("389125467"))
}

func TestPart1(t *testing.T) {
	tests.AssertEquals(t, 49725386, part1(puzzleInput))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 149245887792, part2("389125467"))
}

func TestPart2(t *testing.T) {
	tests.AssertEquals(t, 538935646702, part2(puzzleInput))
}
