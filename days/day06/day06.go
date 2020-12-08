package main

import (
	"fmt"

	"../../utils/files"
	"../../utils/sets"
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
	runes := sets.NewSet()

	for index, line := range inputs {
		if len(line) != 0 {
			for _, r := range line {
				runes.Add(r)
			}
		}

		if len(line) == 0 || index == lastIndex {
			sum += runes.Length()
			runes = sets.NewSet()
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
			for _, r := range line {
				runes[r]++
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
