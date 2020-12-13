package main

import (
	"../../utils/files"
	"fmt"
	"strconv"
	"strings"
)

const filename = "day13.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	timestamp, buses := parseInputs(inputs)

	minWait := timestamp
	var firstBus *bus

	for _, b := range buses {
		busWait := (timestamp/b.frequency+1)*b.frequency - timestamp
		if busWait < minWait {
			minWait = busWait
			firstBus = b
		}
	}

	if firstBus == nil {
		panic("No bus found")
	}

	return minWait * firstBus.id
}

func part2(inputs []string) int {
	_, buses := parseInputs(inputs)
	timestamp := 0
	loopTime := buses[0].frequency // time to loop to respect constraints

	for _, b := range buses[1:] {
		// Increase timestamp with loop time until constraint is respected for current bus
		for (timestamp+b.offset)%b.frequency != 0 {
			timestamp += loopTime
		}
		loopTime *= b.frequency
	}

	return timestamp
}

type bus struct {
	id, frequency, offset int
}

func parseInputs(inputs []string) (int, []*bus) {
	timestamp, _ := strconv.Atoi(inputs[0])
	var buses []*bus
	for idx, v := range strings.Split(inputs[1], ",") {
		if v == "x" {
			continue
		}
		id, _ := strconv.Atoi(v)
		buses = append(buses, &bus{id, id, idx})
	}
	return timestamp, buses
}
