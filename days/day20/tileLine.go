package main

import "reflect"

type TileLine []rune

func computeBorders(tileContent []TileLine) []TileLine {
	borderTop := make(TileLine, tileSize)
	borderBottom := make(TileLine, tileSize)
	borderLeft := make(TileLine, tileSize)
	borderRight := make(TileLine, tileSize)
	for i := 0; i < tileSize; i++ {
		borderTop[i] = tileContent[0][i]
		borderBottom[i] = tileContent[tileSize-1][i]
		borderLeft[i] = tileContent[i][0]
		borderRight[i] = tileContent[i][tileSize-1]
	}

	return []TileLine{borderTop, borderBottom, borderLeft, borderRight}
}

func (tl TileLine) reverse() TileLine {
	reversed := make(TileLine, tileSize)
	for i := tileSize - 1; i >= 0; i-- {
		reversed[i] = tl[tileSize-1-i]
	}
	return reversed
}

func (tl TileLine) isEqualTo(other TileLine) bool {
	return reflect.DeepEqual(tl, other)
}
