package main

import (
	"../../utils/files"
	"../../utils/str"
	"fmt"
	"strings"
)

const filename = "day22.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	deckP1, deckP2 := parseInputs(inputs)
	winnerDeck := classicCombat(deckP1, deckP2)
	return calculateScore(winnerDeck)
}

func part2(inputs []string) int {
	deckP1, deckP2 := parseInputs(inputs)
	_, winnerDeck := recursiceCombat(deckP1, deckP2)
	return calculateScore(winnerDeck)
}

func classicCombat(deckP1 []int, deckP2 []int) []int {
	for {
		var cardP1, cardP2 int
		cardP1, deckP1 = popFirst(deckP1)
		cardP2, deckP2 = popFirst(deckP2)

		if cardP1 > cardP2 {
			deckP1 = append(deckP1, cardP1, cardP2)
		} else {
			deckP2 = append(deckP2, cardP2, cardP1)
		}

		if len(deckP1) == 0 {
			return deckP2
		}
		if len(deckP2) == 0 {
			return deckP1
		}
	}
}

func recursiceCombat(initialDeckP1 []int, initialDeckP2 []int) (int, []int) {
	deckP1 := make([]int, len(initialDeckP1))
	deckP2 := make([]int, len(initialDeckP2))
	copy(deckP1, initialDeckP1)
	copy(deckP2, initialDeckP2)
	var playedRounds []Round

	for {
		round := Round{deckP1, deckP2}
		for _, played := range playedRounds {
			if round.isEqualTo(played) {
				return 1, deckP1
			}
		}
		playedRounds = append(playedRounds, round)

		var cardP1, cardP2 int
		cardP1, deckP1 = popFirst(deckP1)
		cardP2, deckP2 = popFirst(deckP2)

		var winner int
		if len(deckP1) >= cardP1 && len(deckP2) >= cardP2 {
			winner, _ = recursiceCombat(deckP1[:cardP1], deckP2[:cardP2])
		} else if cardP1 > cardP2 {
			winner = 1
		} else {
			winner = 2
		}

		if winner == 1 {
			deckP1 = append(deckP1, cardP1, cardP2)
		} else {
			deckP2 = append(deckP2, cardP2, cardP1)
		}

		if len(deckP1) == 0 {
			return 2, deckP2
		}
		if len(deckP2) == 0 {
			return 1, deckP1
		}
	}
}

type Round struct {
	deckP1, deckP2 []int
}

func (r Round) isEqualTo(other Round) bool {
	if len(r.deckP1) != len(other.deckP1) || len(r.deckP2) != len(other.deckP2) {
		return false
	}
	for i, cardP1 := range r.deckP1 {
		if other.deckP1[i] != cardP1 {
			return false
		}
	}
	for i, cardP2 := range r.deckP2 {
		if other.deckP2[i] != cardP2 {
			return false
		}
	}
	return true
}

func parseInputs(inputs []string) ([]int, []int) {
	var deckP1 []int
	var deckP2 []int

	currentPlayer := 0
	for _, line := range inputs {
		if strings.Contains(line, "Player") {
			currentPlayer++
			continue
		}
		if line == "" {
			continue
		}
		card := str.Atoi(line)
		if currentPlayer == 1 {
			deckP1 = append(deckP1, card)
		} else {
			deckP2 = append(deckP2, card)
		}
	}

	return deckP1, deckP2
}

func popFirst(deck []int) (int, []int) {
	return deck[0], deck[1:]
}

func calculateScore(deck []int) int {
	score := 0
	for i, card := range deck {
		score += card * (len(deck) - i)
	}
	return score
}
