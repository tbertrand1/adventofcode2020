package main

import (
	"fmt"

	utils "../../utils"
)

const filename = "day03.txt"
const treeChar = '#'

func main() {
	values := utils.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(values))
	fmt.Printf("Part 2: %d\n", part2(values))
}

func part1(values []string) int {
	return countTrees(values, 3, 1)
}

func part2(input []string) int {
	return countTrees(input, 1, 1) *
		countTrees(input, 3, 1) *
		countTrees(input, 5, 1) *
		countTrees(input, 7, 1) *
		countTrees(input, 1, 2)
}

func countTrees(values []string, rightDelta int, downDelta int) int {
	nbTrees := 0
	lineSize := len(values[0])
	top, left := 0, 0
	for top < len(values)-1 {
		top += downDelta
		left = (left + rightDelta) % lineSize
		if values[top][left] == treeChar {
			nbTrees++
		}
	}
	return nbTrees
}
