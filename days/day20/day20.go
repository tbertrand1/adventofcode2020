package main

import (
	"../../utils/files"
	"fmt"
)

const filename = "day20.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	tiles := ParseTiles(inputs)
	corners := findCorners(tiles)
	result := 1
	for _, corner := range corners {
		result *= corner.id
	}
	return result
}

func part2(inputs []string) int {
	tiles := ParseTiles(inputs)
	return len(tiles)
}

func findCorners(tiles map[int]*Tile) []*Tile {
	var corners []*Tile
	for _, tile := range tiles {
		bordersUniques := 0
		for _, border := range tile.borders {
			if findOtherTileWithSameBorder(tile, border, tiles) == nil {
				bordersUniques++
			}
		}
		if bordersUniques == 2 {
			corners = append(corners, tile)
		}
	}
	if len(corners) != 4 {
		panic(fmt.Errorf("invalid number of corners: %d", len(corners)))
	}
	return corners
}

func findOtherTileWithSameBorder(currentTile *Tile, currentBorder TileLine, tiles map[int]*Tile) *Tile {
	for id, tile := range tiles {
		if id == currentTile.id {
			continue
		}

		for _, border := range tile.borders {
			if border.isEqualTo(currentBorder) || border.reverse().isEqualTo(currentBorder) {
				return tile
			}
		}
	}
	return nil
}
