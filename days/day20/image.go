package main

import "fmt"

var nbTilesSide = 12

type Image struct {
	allTiles       []*Tile
	assembledTiles [][]*Tile
	content        [][]rune
}

func (i *Image) findOneCorder() *Tile {
	for _, tile := range i.allTiles {
		bordersUniques := 0
		for _, border := range tile.getBorders() {
			if i.findOtherTileWithSameBorder(tile, border) == nil {
				bordersUniques++
			}
		}
		if bordersUniques == 2 {
			return tile
		}
	}
	panic("no corner found")
}

func (i *Image) findOtherTileWithSameBorder(currentTile *Tile, currentBorder Border) *Tile {
	for _, tile := range i.allTiles {
		if tile.id == currentTile.id {
			continue
		}
		for _, border := range tile.getBorders() {
			if border.isEqualTo(currentBorder) || border.reverse().isEqualTo(currentBorder) {
				return tile
			}
		}
	}
	return nil
}

func (i *Image) assembleImage() {
	topLeftTile := i.findOneCorder()

	for _, topLeftTilePosition := range topLeftTile.getAllPositions() {
		i.resetAssembledTiles()
		i.assembledTiles[0][0] = topLeftTilePosition
		assembledTilesId := map[int]bool{topLeftTile.id: true}

		rowIndex := 0
		colIndex := 1

		noTileAssemble := false
		for !noTileAssemble {
			noTileAssemble = true

			for _, tile := range i.allTiles {
				_, alreadyAssemble := assembledTilesId[tile.id]
				if alreadyAssemble {
					continue
				}

				for _, tilePosition := range tile.getAllPositions() {
					valid := true
					if rowIndex > 0 {
						// Check with tile on previous row
						topTile := i.assembledTiles[rowIndex-1][colIndex]
						valid = valid && topTile.borderBottom.isEqualTo(tilePosition.borderTop)
					}
					if colIndex > 0 {
						// Check with tile on previous column
						leftTile := i.assembledTiles[rowIndex][colIndex-1]
						valid = valid && leftTile.borderRight.isEqualTo(tilePosition.borderLeft)
					}

					if valid {
						// Assemble the tile
						noTileAssemble = false
						i.assembledTiles[rowIndex][colIndex] = tilePosition
						assembledTilesId[tilePosition.id] = true
						if colIndex == nbTilesSide-1 {
							rowIndex += 1
							colIndex = 0
						} else {
							colIndex += 1
						}
						break
					}
				}
			}

			if len(assembledTilesId) == len(i.allTiles) {
				// All tiles assembled
				i.updateImageContent()
				return
			}
		}
	}

	panic("no solution found")
}

func (i *Image) resetAssembledTiles() {
	i.assembledTiles = make([][]*Tile, nbTilesSide)
	for r := range i.assembledTiles {
		i.assembledTiles[r] = make([]*Tile, nbTilesSide)
	}
}

func (i *Image) getCorners() []*Tile {
	return []*Tile{
		i.assembledTiles[0][0],
		i.assembledTiles[0][nbTilesSide-1],
		i.assembledTiles[nbTilesSide-1][0],
		i.assembledTiles[nbTilesSide-1][nbTilesSide-1],
	}
}

func (i *Image) updateImageContent() {
	imageSize := nbTilesSide * (tileSize - 2)
	i.content = make([][]rune, imageSize)
	for r := range i.content {
		i.content[r] = make([]rune, imageSize)
	}

	for rowIndex, row := range i.assembledTiles {
		for colIndex, tile := range row {
			for tileRowIndex := 0; tileRowIndex < tileSize-2; tileRowIndex++ {
				for tileColIndex := 0; tileColIndex < tileSize-2; tileColIndex++ {
					contentRow := rowIndex*(tileSize-2) + tileRowIndex
					contentCol := colIndex*(tileSize-2) + tileColIndex
					i.content[contentRow][contentCol] = tile.content[tileRowIndex+1][tileColIndex+1]
				}
			}
		}
	}
}

func (i *Image) countMonster(monster *Monster) int {
	for _, monsterPosition := range monster.getAllPositions() {
		count := 0

		for rowIndex := 0; rowIndex < len(i.content)-monsterPosition.height; rowIndex++ {
			for colIndex := 0; colIndex < len(i.content[0])-monsterPosition.width; colIndex++ {
				monsterFound := true
				for monsterRowIndex := 0; monsterRowIndex < monsterPosition.height; monsterRowIndex++ {
					for monsterColIndex := 0; monsterColIndex < monsterPosition.width; monsterColIndex++ {
						if monsterPosition.content[monsterRowIndex][monsterColIndex] != '#' {
							continue
						}
						if i.content[rowIndex+monsterRowIndex][colIndex+monsterColIndex] != '#' {
							monsterFound = false
							break
						}
					}
					if !monsterFound {
						break
					}
				}
				if monsterFound {
					count++
				}
			}
		}

		if count != 0 {
			return count
		}
	}
	return 0
}

func (i *Image) printTiles() {
	for rowIndex, row := range i.assembledTiles {
		fmt.Printf("---- Row %d\n", rowIndex)
		for _, tile := range row {
			fmt.Print(tile.toString())
		}
		fmt.Println()
	}
}

func (i *Image) printImage() {
	for _, row := range i.content {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func (i *Image) countRune(r rune) int {
	count := 0
	for _, row := range i.content {
		for _, col := range row {
			if col == r {
				count += 1
			}
		}
	}
	return count
}
