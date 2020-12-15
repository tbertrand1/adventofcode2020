package main

import (
	"../../utils/tests"
	"testing"
)

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 436, part1("0,3,6"))
}

func TestPart1(t *testing.T) {
	tests.AssertEquals(t, 694, part1("17,1,3,16,19,0"))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 175594, part2("0,3,6"))
}

func TestPart2(t *testing.T) {
	tests.AssertEquals(t, 21768614, part2("17,1,3,16,19,0"))
}
