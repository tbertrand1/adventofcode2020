package main

import (
	"testing"

	files "../../utils/files"
	tests "../../utils/tests"
)

func testInputs() []string {
	return []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
}

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part1()", 7, part1(testInputs()))
}

func TestPart1(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part1()", 284, part1(inputs))
}

func TestPart2WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, "part2()", 336, part2(testInputs()))
}

func TestPart2(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part2()", 3510149120, part2(inputs))
}

func TestPart2WithGoroutines(t *testing.T) {
	inputs := files.ReadLinesOfFile(filename)
	tests.AssertEquals(t, "part2WithGoroutines()", 3510149120, part2WithGoroutines(inputs))
}

func BenchmarkPart2(b *testing.B) {
	inputs := files.ReadLinesOfFile(filename)
	for i := 0; i < b.N; i++ {
		part2(inputs)
	}
}

func BenchmarkPart2WithGoroutines(b *testing.B) {
	inputs := files.ReadLinesOfFile(filename)
	for i := 0; i < b.N; i++ {
		part2WithGoroutines(inputs)
	}
}
