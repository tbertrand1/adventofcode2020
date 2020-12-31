package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func testInputs() []string {
	return []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 5, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 1685, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	actual := part2(testInputs())
	expected := "mxmxvkd,sqjhc,fvjkl"
	if expected != actual {
		t.Errorf("Test failed; expected %s; actual %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	actual := part2(inputs)
	expected := "ntft,nhx,kfxr,xmhsbd,rrjb,xzhxj,chbtp,cqvc"
	if expected != actual {
		t.Errorf("Test failed; expected %s; actual %s", expected, actual)
	}
}
