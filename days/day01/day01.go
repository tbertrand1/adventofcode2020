package main

import (
	"fmt"

	files "../../utils/files"
)

const filename = "day01.txt"

func main() {
	values := files.ReadLinesAsIntsOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(values))
	fmt.Printf("Part 2: %d\n", part2(values))
}

func part1(values []int) int {
	for index1, value1 := range values {
		for _, value2 := range values[:index1] {
			if value1+value2 == 2020 {
				return value1 * value2
			}
		}
	}
	panic("Not found")
}

func part2(values []int) int {
	for index1, value1 := range values {
		remainingValues := values[:index1]
		for index2, value2 := range remainingValues {
			for _, value3 := range remainingValues[:index2] {
				if value1+value2+value3 == 2020 {
					return value1 * value2 * value3
				}
			}
		}
	}
	panic("Not found")
}
