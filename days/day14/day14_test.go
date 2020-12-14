package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func TestPart1WithTestInputs(t *testing.T) {
	testInputs := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	tests.AssertEquals(t, 165, execute(testInputs, true))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 10717676595607, execute(inputs, true))
}

func TestPart2WithTestInputs(t *testing.T) {
	testInputs := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	tests.AssertEquals(t, 208, execute(testInputs, false))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 3974538275659, execute(inputs, false))
}
