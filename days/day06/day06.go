package main

import (
	"fmt"

	files "../../utils/files"
)

const filename = "day06.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)

	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	sum := 0

	lastIndex := len(inputs) - 1
	runes := make(map[rune]bool)

	for index, line := range inputs {
		if len(line) != 0 {
			for _, rune := range line {
				runes[rune] = true
			}
		}

		if len(line) == 0 || index == lastIndex {
			sum += len(runes)
			runes = make(map[rune]bool)
		}
	}

	return sum
}

func part2(inputs []string) int {
	sum := 0

	lastIndex := len(inputs) - 1
	runes := make(map[rune]int)
	personsInGroup := 0

	for index, line := range inputs {
		if len(line) != 0 {
			personsInGroup++
			for _, rune := range line {
				runes[rune]++
			}
		}

		if len(line) == 0 || index == lastIndex {
			for _, count := range runes {
				if count == personsInGroup {
					sum++
				}
			}
			runes = make(map[rune]int)
			personsInGroup = 0
		}
	}

	return sum
}
