package main

import (
	"../../utils/files"
	"../../utils/slices"
	"fmt"
	"sort"
)

const filename = "day10.txt"

func main() {
	inputs := files.ReadLinesAsIntsOfFile(filename)
	inputs = completeInputs(inputs)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func completeInputs(inputs []int) []int {
	sort.Ints(inputs)
	var completed []int
	completed = append(completed, 0)
	completed = append(completed, inputs...)
	completed = append(completed, slices.Last(inputs)+3)
	return completed
}

func part1(inputs []int) int {
	differencesOf1, differencesOf3 := 0, 0

	previousInput := inputs[0]
	for _, input := range inputs {
		difference := input - previousInput
		if difference == 1 {
			differencesOf1++
		}
		if difference == 3 {
			differencesOf3++
		}
		previousInput = input
	}

	return differencesOf1 * differencesOf3
}

func part2(inputs []int) int {
	nbWays := make(map[int]int)
	nbWays[0] = 1

	for _, input := range inputs[1:] {
		nbWays[input] = nbWays[input-1] + nbWays[input-2] + nbWays[input-3]
	}

	return nbWays[slices.Last(inputs)]
}
