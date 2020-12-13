package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs() []string {
	return []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 295, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 156, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 1068781, part2(testInputs()))
	tests.AssertEquals(t, 3417, part2([]string{"0", "17,x,13,19"}))
	tests.AssertEquals(t, 754018, part2([]string{"0", "67,7,59,61"}))
	tests.AssertEquals(t, 779210, part2([]string{"0", "67,x,7,59,61"}))
	tests.AssertEquals(t, 1261476, part2([]string{"0", "67,7,x,59,61"}))
	tests.AssertEquals(t, 1202161486, part2([]string{"0", "1789,37,47,1889"}))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 404517869995362, part2(inputs))
}
