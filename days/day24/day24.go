package main

import (
	"../../utils/files"
	"fmt"
)

const filename = "day24.txt"

func main() {
	lines := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	blackTiles := parseInputs(lines)
	return len(blackTiles)
}

func part2(lines []string) int {
	blackTiles := parseInputs(lines)

	for n := 0; n < 100; n++ {
		nbAdjacentsByCoord := make(map[Coord]int)

		for coord := range blackTiles {
			for _, delta := range directions {
				adjacent := Coord{coord.x + delta.x, coord.y + delta.y}
				nbAdjacentsByCoord[adjacent] += 1
			}
		}

		newBlackTiles := make(map[Coord]bool)
		for coord, nbAdjacents := range nbAdjacentsByCoord {
			_, wasBlack := blackTiles[coord]
			if wasBlack && (nbAdjacents == 1 || nbAdjacents == 2) {
				newBlackTiles[coord] = true
			}
			if !wasBlack && nbAdjacents == 2 {
				newBlackTiles[coord] = true
			}
		}

		blackTiles = newBlackTiles
	}

	return len(blackTiles)
}

type Coord struct {
	x, y float64
}

var directions = map[string]Coord{
	"ne": {0.5, 1},
	"e":  {1, 0},
	"se": {0.5, -1},
	"nw": {-0.5, 1},
	"w":  {-1, 0},
	"sw": {-0.5, -1},
}

func parseInputs(lines []string) map[Coord]bool {
	tiles := make(map[Coord]int)

	for _, line := range lines {
		x := 0.
		y := 0.

		n := 0
		buffer := ""
		for n < len(line) {
			char := string(line[n])
			if char == "w" || char == "e" {
				delta := directions[buffer+char]
				x += delta.x
				y += delta.y
				buffer = ""
			} else {
				buffer = char
			}
			n++
		}

		tiles[Coord{x, y}] += 1
	}

	blackTiles := make(map[Coord]bool)
	for coord, nbFlipped := range tiles {
		if nbFlipped%2 == 1 {
			blackTiles[coord] = true
		}
	}
	return blackTiles
}
