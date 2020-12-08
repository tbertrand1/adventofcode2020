package main

import (
	"fmt"
	"regexp"
	"strconv"

	"../../utils/files"
)

const filename = "day02.txt"
const inputRegex = "(\\d+)-(\\d+) ([a-z]): ([a-z]+)"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	valids := 0
	regex := regexp.MustCompile(inputRegex)

	for _, input := range inputs {
		matches := regex.FindStringSubmatch(input)
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if isPasswordValidForMinMax(matches[4], min, max, matches[3][0]) {
			valids++
		}
	}

	return valids
}

func part2(inputs []string) int {
	valids := 0
	regex := regexp.MustCompile(inputRegex)

	for _, input := range inputs {
		matches := regex.FindStringSubmatch(input)
		position1, _ := strconv.Atoi(matches[1])
		position2, _ := strconv.Atoi(matches[2])
		if isPasswordValidForExactPositions(matches[4], position1, position2, matches[3][0]) {
			valids++
		}
	}

	return valids
}

func isPasswordValidForMinMax(password string, min int, max int, policyLetter byte) bool {
	countLetter := 0
	for i := 0; i < len(password); i++ {
		if password[i] == policyLetter {
			countLetter++
		}
	}
	return countLetter >= min && countLetter <= max
}

func isPasswordValidForExactPositions(password string, position1 int, position2 int, policyLetter byte) bool {
	position1ContainsLetter := password[position1-1] == policyLetter
	position2ContainsLetter := password[position2-1] == policyLetter
	return (position1ContainsLetter && !position2ContainsLetter) ||
		(!position1ContainsLetter && position2ContainsLetter)
}
