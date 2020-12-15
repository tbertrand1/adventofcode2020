package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "17,1,3,16,19,0"
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	return findNumber(input, 2020)
}

func part2(input string) int {
	return findNumber(input, 30000000)
}

func findNumber(input string, end int) int {
	// Array of number spoken, with number as index and position as value
	spoken := make([]int, end)

	var lastNumber int
	position := 1
	for _, n := range strings.Split(input, ",") {
		lastNumber, _ = strconv.Atoi(n)
		spoken[lastNumber] = position
		position++
	}

	for ; position <= end; position++ {
		currentNumber := 0
		previousPosition := position - 1
		if spoken[lastNumber] != 0 {
			// Number already spoken
			currentNumber = previousPosition - spoken[lastNumber]
		}
		spoken[lastNumber] = previousPosition
		lastNumber = currentNumber
	}

	return lastNumber
}
