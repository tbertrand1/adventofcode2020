package main

import (
	"testing"
)

func initialTile() *Tile {
	return NewTileFromData([]string{
		"Tile 1234:",
		"0123456789",
		"a.........",
		"b#########",
		"c.........",
		"d#########",
		"e.........",
		"f#########",
		"g.........",
		"h#########",
		"i.........",
	})
}

func TestToString(t *testing.T) {
	expected :=
		`-- Tile 1234
0123456789
a.........
b#########
c.........
d#########
e.........
f#########
g.........
h#########
i.........
`

	tile := initialTile()
	if expected != tile.toString() {
		t.Error("Test failed")
	}
}

func TestRotate(t *testing.T) {
	expected :=
		`-- Tile 1234
ihgfedcba0
.#.#.#.#.1
.#.#.#.#.2
.#.#.#.#.3
.#.#.#.#.4
.#.#.#.#.5
.#.#.#.#.6
.#.#.#.#.7
.#.#.#.#.8
.#.#.#.#.9
`

	tile := initialTile().rotate()
	if expected != tile.toString() {
		t.Error("Test failed")
	}
}

func TestFlip(t *testing.T) {
	expected :=
		`-- Tile 1234
9876543210
.........a
#########b
.........c
#########d
.........e
#########f
.........g
#########h
.........i
`

	tile := initialTile().flip()
	if expected != tile.toString() {
		t.Error("Test failed")
	}
}
