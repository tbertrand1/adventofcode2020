package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var tileSize = 10
var tileHeaderRegex = regexp.MustCompile(`^Tile (\d{4}):$`)

type Tile struct {
	id                                               int
	content                                          [][]rune
	borderTop, borderBottom, borderLeft, borderRight Border
	currentPosition                                  int
}

type Border []rune

func ParseTiles(inputs []string) []*Tile {
	var tiles []*Tile

	var tileData []string
	for i := 0; i < len(inputs); i++ {
		line := inputs[i]
		if line != "" {
			tileData = append(tileData, line)
		}
		if line == "" || i == len(inputs)-1 {
			tile := NewTileFromData(tileData)
			tiles = append(tiles, tile)
			tileData = []string{}
		}
	}

	return tiles
}

func NewTileFromData(data []string) *Tile {
	match := tileHeaderRegex.FindStringSubmatch(data[0])
	if match == nil {
		panic(fmt.Errorf("invalid tile header: %s", data[0]))
	}

	id, _ := strconv.Atoi(match[1])

	contentData := data[1:]
	var content [][]rune
	for _, row := range contentData {
		var line []rune
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

	tile := Tile{id: id, content: content, currentPosition: 0}
	tile.computeBorders()
	return &tile
}

func (t *Tile) computeBorders() {
	t.borderTop = make(Border, tileSize)
	t.borderBottom = make(Border, tileSize)
	t.borderLeft = make(Border, tileSize)
	t.borderRight = make(Border, tileSize)
	for i := 0; i < tileSize; i++ {
		t.borderTop[i] = t.content[0][i]
		t.borderBottom[i] = t.content[tileSize-1][i]
		t.borderLeft[i] = t.content[i][0]
		t.borderRight[i] = t.content[i][tileSize-1]
	}
}

func (t *Tile) getBorders() []Border {
	return []Border{t.borderTop, t.borderBottom, t.borderLeft, t.borderRight}
}

func (t *Tile) toString() string {
	str := fmt.Sprintf("-- Tile %d\n", t.id)
	for _, row := range t.content {
		for _, col := range row {
			str += string(col)
		}
		str += "\n"
	}
	return str
}

func initialEmptyContent(content [][]rune) [][]rune {
	emptyContent := make([][]rune, len(content))
	for row := range content {
		emptyContent[row] = make([]rune, len(content[row]))
	}
	return emptyContent
}

func (t *Tile) rotate() *Tile {
	newContent := initialEmptyContent(t.content)
	for rowIndex := range t.content {
		for colIndex := range t.content[rowIndex] {
			newRowIndex := colIndex
			newColIndex := len(t.content) - 1 - rowIndex
			newContent[newRowIndex][newColIndex] = t.content[rowIndex][colIndex]
		}
	}

	newTile := Tile{id: t.id, content: newContent, currentPosition: t.currentPosition + 1}
	newTile.computeBorders()
	return &newTile
}

func (t *Tile) flip() *Tile {
	newContent := initialEmptyContent(t.content)
	for rowIndex := range t.content {
		for colIndex := range t.content[rowIndex] {
			newRowIndex := rowIndex
			newColIndex := len(t.content[rowIndex]) - 1 - colIndex
			newContent[newRowIndex][newColIndex] = t.content[rowIndex][colIndex]
		}
	}

	newTile := Tile{id: t.id, content: newContent, currentPosition: t.currentPosition + 1}
	newTile.computeBorders()
	return &newTile
}

func (t *Tile) getAllPositions() []*Tile {
	rotate1 := t.rotate()
	rotate2 := rotate1.rotate()
	rotate3 := rotate2.rotate()
	flip1 := rotate3.flip()
	flipRotate1 := flip1.rotate()
	flipRotate2 := flipRotate1.rotate()
	flipRotate3 := flipRotate2.rotate()
	return []*Tile{t, rotate1, rotate2, rotate3, flip1, flipRotate1, flipRotate2, flipRotate3}
}

func (b Border) reverse() Border {
	reversed := make(Border, tileSize)
	for i := tileSize - 1; i >= 0; i-- {
		reversed[i] = b[tileSize-1-i]
	}
	return reversed
}

func (b Border) isEqualTo(other Border) bool {
	return reflect.DeepEqual(b, other)
}
