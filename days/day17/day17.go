package main

import (
	"../../utils/files"
	"../../utils/sets"
	"fmt"
)

const filename = "day17.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	return countActivesAfterNbCycles(inputs, 6, 3)
}

func part2(inputs []string) int {
	return countActivesAfterNbCycles(inputs, 6, 4)
}

func countActivesAfterNbCycles(inputs []string, nbCycles, dimensions int) int {
	actives := sets.NewSet()
	for y, line := range inputs {
		for x, cube := range line {
			if cube == '#' {
				actives.Add(coordinate{x, y, 0, 0})
			}
		}
	}

	for cycle := 0; cycle < nbCycles; cycle++ {
		neighbors := make(map[coordinate]int)

		for _, c := range actives.Keys() {
			coord := c.(coordinate)

			for x := coord.x - 1; x <= coord.x+1; x++ {
				for y := coord.y - 1; y <= coord.y+1; y++ {
					for z := coord.z - 1; z <= coord.z+1; z++ {
						if dimensions == 3 {
							neighbor := coordinate{x, y, z, 0}
							if coord != neighbor {
								neighbors[neighbor]++
							}
						} else if dimensions == 4 {
							for w := coord.w - 1; w <= coord.w+1; w++ {
								neighbor := coordinate{x, y, z, w}
								if coord != neighbor {
									neighbors[neighbor]++
								}
							}
						} else {
							panic("Dimensions not supported")
						}
					}
				}
			}
		}

		newActives := sets.NewSet()
		for coord, count := range neighbors {
			wasActive := actives.Contains(coord)
			if wasActive && (count == 2 || count == 3) {
				newActives.Add(coord)
			}
			if !wasActive && count == 3 {
				newActives.Add(coord)
			}
		}

		actives = newActives
	}

	return actives.Length()
}

type coordinate struct {
	x, y, z, w int
}
