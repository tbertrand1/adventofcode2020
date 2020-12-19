package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 71, evalaluateInputPart1("1 + 2 * 3 + 4 * 5 + 6"))
	tests.AssertEquals(t, 51, evalaluateInputPart1("1 + (2 * 3) + (4 * (5 + 6))"))
	tests.AssertEquals(t, 26, evalaluateInputPart1("2 * 3 + (4 * 5)"))
	tests.AssertEquals(t, 437, evalaluateInputPart1("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	tests.AssertEquals(t, 12240, evalaluateInputPart1("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	tests.AssertEquals(t, 13632, evalaluateInputPart1("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 3885386961962, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 80, evalaluateInputPart2("8 * 6 + 4"))
	tests.AssertEquals(t, 231, evalaluateInputPart2("1 + 2 * 3 + 4 * 5 + 6"))
	tests.AssertEquals(t, 51, evalaluateInputPart2("1 + (2 * 3) + (4 * (5 + 6))"))
	tests.AssertEquals(t, 46, evalaluateInputPart2("2 * 3 + (4 * 5)"))
	tests.AssertEquals(t, 1445, evalaluateInputPart2("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	tests.AssertEquals(t, 669060, evalaluateInputPart2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	tests.AssertEquals(t, 23340, evalaluateInputPart2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 112899558798666, part2(inputs))
}
