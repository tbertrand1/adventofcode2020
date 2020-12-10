package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs1() []int {
	return []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
}

func testInputs2() []int {
	return []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 35, part1(completeInputs(testInputs1())))
	tests.AssertEquals(t, 220, part1(completeInputs(testInputs2())))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, 2310, part1(completeInputs(inputs)))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 8, part2(completeInputs(testInputs1())))
	tests.AssertEquals(t, 19208, part2(completeInputs(testInputs2())))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	tests.AssertEquals(t, 64793042714624, part2(completeInputs(inputs)))
}
