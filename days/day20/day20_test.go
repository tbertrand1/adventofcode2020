package main

import (
	"../../utils/files"
	"../../utils/tests"
	"testing"
)

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, 111936085519519, part1(inputs))
}
