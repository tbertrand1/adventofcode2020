package main

import (
	"../../utils/files"
	"fmt"
)

const filename = "day09.txt"

func main() {
	inputs := files.ReadLinesAsIntsOfFile(filename)

	invalidNumber, index := part1(inputs, 25)
	fmt.Printf("Part 1: %d\n", invalidNumber)
	fmt.Printf("Part 2: %d\n", part2(inputs, invalidNumber, index))
}

func part1(inputs []int, preamble int) (int, int) {
	for i := preamble; i < len(inputs); i++ {
		if !isValid(inputs, preamble, i) {
			return inputs[i], i
		}
	}
	panic("Not found")
}

func isValid(inputs []int, preamble int, index int) bool {
	for i := index - preamble; i < index-1; i++ {
		for j := i + 1; j < index; j++ {
			if inputs[i]+inputs[j] == inputs[index] {
				return true
			}
		}
	}
	return false
}

func part2(inputs []int, invalidNumber int, index int) int {
	for i := 0; i < index; i++ {
		numberI := inputs[i]
		min, max, sum := numberI, numberI, numberI

		for j := i + 1; j < index && sum < invalidNumber; j++ {
			numberJ := inputs[j]
			sum += numberJ
			if sum == invalidNumber {
				return min + max
			}

			if numberJ > max {
				max = numberJ
			}
			if numberJ < min {
				min = numberJ
			}
		}
	}
	panic("Not found")
}
