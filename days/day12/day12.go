package main

import (
	"../../utils/files"
	"fmt"
	"strconv"
)

const filename = "day12.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	ship := &ship{p: position{0, 0}, d: east}

	for _, input := range inputs {
		value, _ := strconv.Atoi(input[1:])
		switch input[0] {
		case 'N':
			ship.translate(north, value)
		case 'S':
			ship.translate(south, value)
		case 'E':
			ship.translate(east, value)
		case 'W':
			ship.translate(west, value)
		case 'L':
			ship.rotateRight(360 - value)
		case 'R':
			ship.rotateRight(value)
		case 'F':
			ship.translate(ship.d, value)
		}
	}

	return ship.manhattanDistance()
}

func part2(inputs []string) int {
	ship := &ship{p: position{0, 0}, d: east}
	waypoint := &waypoint{p: position{10, 1}}

	for _, input := range inputs {
		value, _ := strconv.Atoi(input[1:])
		switch input[0] {
		case 'N':
			waypoint.translate(north, value)
		case 'S':
			waypoint.translate(south, value)
		case 'E':
			waypoint.translate(east, value)
		case 'W':
			waypoint.translate(west, value)
		case 'L':
			ship.rotateWaypointRight(waypoint, 360-value)
		case 'R':
			ship.rotateWaypointRight(waypoint, value)
		case 'F':
			ship.translateToWaypoint(waypoint, value)
		}
	}

	return ship.manhattanDistance()
}

type position struct {
	x, y int
}

type ship struct {
	p position
	d direction
}

type waypoint struct {
	p position
}

type direction struct {
	dx, dy, index int
}

var north = direction{0, 1, 0}
var east = direction{1, 0, 1}
var south = direction{0, -1, 2}
var west = direction{-1, 0, 3}
var directions = []direction{
	north, east, south, west,
}

func (s *ship) translate(d direction, v int) {
	s.p.x += v * d.dx
	s.p.y += v * d.dy
}

func (s *ship) rotateRight(v int) {
	indexNewDirection := (s.d.index + v/90) % 4
	s.d = directions[indexNewDirection]
}

func (s *ship) translateToWaypoint(w *waypoint, v int) {
	s.p.x += v * w.p.x
	s.p.y += v * w.p.y
}

func (s *ship) rotateWaypointRight(w *waypoint, v int) {
	if v == 90 {
		w.p.x, w.p.y = w.p.y, -w.p.x
	} else if v == 180 {
		w.p.x, w.p.y = -w.p.x, -w.p.y
	} else if v == 270 {
		w.p.x, w.p.y = -w.p.y, w.p.x
	} else {
		panic("Angle not supported")
	}
}

func (w *waypoint) translate(d direction, v int) {
	w.p.x += v * d.dx
	w.p.y += v * d.dy
}

func (s *ship) manhattanDistance() int {
	return abs(s.p.x) + abs(s.p.y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
