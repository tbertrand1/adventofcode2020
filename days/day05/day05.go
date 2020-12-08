package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../../utils/files"
)

const filename = "day05.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)

	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	var maxSeatID int
	for _, input := range inputs {
		seatID := calculateSeatID(input)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func part2(inputs []string) int {
	var seatsID []int
	for _, input := range inputs {
		seatsID = append(seatsID, calculateSeatID(input))
	}

	sort.Ints(seatsID)
	previousSeatID := 0
	for _, seatID := range seatsID {
		if previousSeatID != 0 {
			if seatID-previousSeatID == 2 {
				return previousSeatID + 1
			}
		}
		previousSeatID = seatID
	}

	panic("Not found")
}

func calculateSeatID(input string) int {
	row := binaryToInt(input[:7])
	column := binaryToInt(input[7:])
	return row*8 + column
}

func binaryToInt(input string) int {
	binary := strings.ReplaceAll(input, "B", "1")
	binary = strings.ReplaceAll(binary, "R", "1")
	binary = strings.ReplaceAll(binary, "F", "0")
	binary = strings.ReplaceAll(binary, "L", "0")
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic("Cannot convert binary to int")
	}
	return int(i)
}
