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
	image := Image{allTiles: tiles}
	image.assembleImage()

	result := 1
	for _, corner := range image.getCorners() {
		result *= corner.id
	}
	return result
}

func part2(inputs []string) int {
	tiles := ParseTiles(inputs)
	image := Image{allTiles: tiles}
	image.assembleImage()

	//image.printTiles()
	//image.printImage()

	monster := NewMonsterFromData([]string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	})

	nbMonsters := image.countMonster(monster)
	return image.countRune('#') - nbMonsters*monster.countRune('#')
}
