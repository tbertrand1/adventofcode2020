package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var tileSize = 10
var tileHeaderRegex = regexp.MustCompile(`^Tile (\d{4}):$`)

type Tile struct {
	id      int
	content []TileLine
	borders []TileLine
}

func NewTileFromData(data []string) *Tile {
	match := tileHeaderRegex.FindStringSubmatch(data[0])
	if match == nil {
		panic(fmt.Errorf("invalid tile header: %s", data[0]))
	}

	id, _ := strconv.Atoi(match[1])

	contentData := data[1:]
	var content []TileLine
	for _, row := range contentData {
		var line TileLine
		for _, letter := range row {
			line = append(line, letter)
		}
		if len(line) != tileSize {
			panic("wrong size")
		}
		content = append(content, line)
	}
	if len(content) != tileSize {
		panic("wrong size")
	}

	tile := Tile{id: id, content: content}
	tile.borders = computeBorders(content)
	return &tile
}

func ParseTiles(inputs []string) map[int]*Tile {
	tiles := make(map[int]*Tile)

	var tileData []string
	for i := 0; i < len(inputs); i++ {
		line := inputs[i]
		if line != "" {
			tileData = append(tileData, line)
		}
		if line == "" || i == len(inputs)-1 {
			tile := NewTileFromData(tileData)
			tiles[tile.id] = tile
			tileData = []string{}
		}
	}

	return tiles
}
