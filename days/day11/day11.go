package main

import (
	"../../utils/files"
	"fmt"
)

const filename = "day11.txt"
const floor = '.'
const emptySeat = 'L'
const occupiedSeat = '#'

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	grid := newGrid(inputs)

	for {
		changed, nbOccupied := grid.nextRoundPart1()
		if !changed {
			return nbOccupied
		}
	}
}

func part2(inputs []string) int {
	grid := newGrid(inputs)

	for {
		changed, nbOccupied := grid.nextRoundPart2()
		if !changed {
			return nbOccupied
		}
	}
}

type Grid struct {
	seats, nextSeats [][]rune
	width, height    int
}

func newGrid(inputs []string) Grid {
	var seats [][]rune
	for _, row := range inputs {
		var rowSeats []rune
		for _, seat := range row {
			rowSeats = append(rowSeats, seat)
		}
		seats = append(seats, rowSeats)
	}

	return Grid{seats: seats, width: len(seats[0]), height: len(seats)}
}

func (grid *Grid) resetNextSeats() {
	grid.nextSeats = make([][]rune, grid.height)
	for i := range grid.nextSeats {
		grid.nextSeats[i] = make([]rune, grid.width)
	}
}

func (grid *Grid) nextRoundPart1() (bool, int) {
	//grid.display()
	grid.resetNextSeats()
	changed := false
	nbOccupied := 0

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			seat := grid.seats[y][x]
			occupiedAdjacents := grid.countOccupiedAdjacents(x, y)
			switch seat {
			case floor:
				grid.nextSeats[y][x] = floor
			case emptySeat:
				if occupiedAdjacents == 0 {
					grid.nextSeats[y][x] = occupiedSeat
					changed = true
					nbOccupied++
				} else {
					grid.nextSeats[y][x] = emptySeat
				}
			case occupiedSeat:
				if occupiedAdjacents >= 4 {
					grid.nextSeats[y][x] = emptySeat
					changed = true
				} else {
					grid.nextSeats[y][x] = occupiedSeat
					nbOccupied++
				}
			}
		}
	}
	grid.seats = grid.nextSeats

	return changed, nbOccupied
}

func (grid *Grid) nextRoundPart2() (bool, int) {
	//grid.display()
	grid.resetNextSeats()
	changed := false
	nbOccupied := 0

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			seat := grid.seats[y][x]
			occupiedVisibles := grid.countOccupiedVisibles(x, y)
			switch seat {
			case floor:
				grid.nextSeats[y][x] = floor
			case emptySeat:
				if occupiedVisibles == 0 {
					grid.nextSeats[y][x] = occupiedSeat
					changed = true
					nbOccupied++
				} else {
					grid.nextSeats[y][x] = emptySeat
				}
			case occupiedSeat:
				if occupiedVisibles >= 5 {
					grid.nextSeats[y][x] = emptySeat
					changed = true
				} else {
					grid.nextSeats[y][x] = occupiedSeat
					nbOccupied++
				}
			}
		}
	}
	grid.seats = grid.nextSeats

	return changed, nbOccupied
}

var offsets = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (grid *Grid) countOccupiedAdjacents(x int, y int) int {
	count := 0

	for _, offset := range offsets {
		seatX := x + offset[0]
		seatY := y + offset[1]

		if seatX < 0 || seatX >= grid.width ||
			seatY < 0 || seatY >= grid.height {
			continue
		}
		if grid.seats[seatY][seatX] == occupiedSeat {
			count++
		}
	}

	return count
}

func (grid *Grid) countOccupiedVisibles(x int, y int) int {
	count := 0

	for _, offset := range offsets {
		seatX := x
		seatY := y
		for {
			seatX = seatX + offset[0]
			seatY = seatY + offset[1]
			if seatX < 0 || seatX >= grid.width ||
				seatY < 0 || seatY >= grid.height {
				break
			}
			seat := grid.seats[seatY][seatX]
			if seat == floor {
				continue
			}
			if seat == occupiedSeat {
				count++
			}
			break
		}
	}

	return count
}

func (grid *Grid) display() {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			fmt.Print(string(grid.seats[y][x]))
		}
		fmt.Println()
	}
	fmt.Println()
}
