package main

import (
	"../../utils/str"
	"fmt"
)

const puzzleInput = "716892543"

type Cup struct {
	value int
	next  *Cup
}

func main() {
	fmt.Printf("Part 1: %d\n", part1(puzzleInput))
	fmt.Printf("Part 2: %d\n", part2(puzzleInput))
}

func part1(input string) int {
	cups, firstCup := parseInputPart1(input)
	moveCups(cups, firstCup, 100)

	result := ""
	for cup := cups[1].next; cup.value != 1; cup = cup.next {
		result += fmt.Sprint(cup.value)
	}
	return str.Atoi(result)
}

func part2(input string) int {
	cups, firstCup := parseInputPart2(input)
	moveCups(cups, firstCup, 10000000)

	return cups[1].next.value * cups[1].next.next.value
}

func parseInputPart1(input string) (map[int]*Cup, *Cup) {
	cups := make(map[int]*Cup)
	var firstCup, lastCup *Cup

	for _, label := range input {
		cup := &Cup{value: int(label - '0')}
		cups[cup.value] = cup

		if lastCup == nil {
			firstCup = cup
		} else {
			lastCup.next = cup
		}

		lastCup = cup
	}

	lastCup.next = firstCup

	return cups, firstCup
}

func parseInputPart2(input string) (map[int]*Cup, *Cup) {
	cups := make(map[int]*Cup)
	var firstCup, lastCup *Cup

	for i := 0; i < 1000000; i++ {
		var cup *Cup
		if i < len(input) {
			cup = &Cup{value: int(input[i] - '0')}
		} else {
			cup = &Cup{value: i + 1}
		}
		cups[cup.value] = cup

		if lastCup == nil {
			firstCup = cup
		} else {
			lastCup.next = cup
		}

		lastCup = cup
	}

	lastCup.next = firstCup

	return cups, firstCup
}

func moveCups(cups map[int]*Cup, firstCup *Cup, nbMoves int) {
	currentCup := firstCup

	for n := 0; n < nbMoves; n++ {
		cup1 := currentCup.next
		cup2 := cup1.next
		cup3 := cup2.next

		currentCup.next = cup3.next

		destinationValue := determineDestinationValue(currentCup.value, cup1.value, cup2.value, cup3.value, len(cups))
		destinationCup := cups[destinationValue]
		if destinationCup == nil {
			panic(fmt.Errorf("destination not found for value: %d", destinationValue))
		}

		cup3.next = destinationCup.next
		destinationCup.next = cup1

		currentCup = currentCup.next
	}
}

func determineDestinationValue(currentCupValue, cup1Value, cup2Value, cup3Value, nbCups int) int {
	destinationValue := currentCupValue
	if destinationValue == 1 {
		destinationValue = nbCups
	} else {
		destinationValue = destinationValue - 1
	}

	for destinationValue == cup1Value || destinationValue == cup2Value || destinationValue == cup3Value {
		if destinationValue == 1 {
			destinationValue = nbCups
		} else {
			destinationValue = destinationValue - 1
		}
	}

	return destinationValue
}
