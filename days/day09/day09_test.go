package main

import (
	"testing"

	"../../utils/files"
	"../../utils/tests"
)

func testInputs() []int {
	return []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	invalidNumber, index := part1(testInputs(), 5)
	tests.AssertEquals(t, 127, invalidNumber)
	tests.AssertEquals(t, 14, index)
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	invalidNumber, index := part1(inputs, 25)
	tests.AssertEquals(t, 1504371145, invalidNumber)
	tests.AssertEquals(t, 652, index)
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 62, part2(testInputs(), 127, 14))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, 183278487, part2(inputs, 1504371145, 652))
}
