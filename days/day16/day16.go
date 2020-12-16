package main

import (
	"../../utils/files"
	"../../utils/sets"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const filename = "day16.txt"

var constraintRegex = regexp.MustCompile(`^(.*): (\d*)-(\d*) or (\d*)-(\d*)$`)

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	constraints, tickets := parseInputs(inputs)
	_, sumInvalids := findValidTickets(constraints, tickets[1:])
	return sumInvalids
}

func part2(inputs []string) int {
	constraints, tickets := parseInputs(inputs)
	nbFields := len(constraints)
	myTicket := tickets[0]

	ticketsValids, _ := findValidTickets(constraints, tickets[1:])

	fieldsWithPositions := determineFieldsWithPositionsPossibles(constraints, ticketsValids, nbFields)
	sort.Slice(fieldsWithPositions, func(i, j int) bool {
		return len(fieldsWithPositions[i].positions) < len(fieldsWithPositions[j].positions)
	})

	solution := make(map[string]int)
	if !findSolution(fieldsWithPositions, nbFields, solution) {
		panic("no solution found")
	}

	result := 1
	for field, position := range solution {
		if strings.HasPrefix(field, "departure") {
			result *= myTicket.values[position]
		}
	}
	return result
}

type constraint struct {
	field                  string
	min1, max1, min2, max2 int
}

type ticket struct {
	values []int
}

type fieldWithPositions struct {
	field     string
	positions []int
}

func (c constraint) isValid(v int) bool {
	return (v >= c.min1 && v <= c.max1) || (v >= c.min2 && v <= c.max2)
}

func parseInputs(inputs []string) ([]constraint, []ticket) {
	var constraints []constraint
	var tickets []ticket

	for _, line := range inputs {
		matches := constraintRegex.FindStringSubmatch(line)
		if matches != nil {
			min1, _ := strconv.Atoi(matches[2])
			max1, _ := strconv.Atoi(matches[3])
			min2, _ := strconv.Atoi(matches[4])
			max2, _ := strconv.Atoi(matches[5])
			c := constraint{field: matches[1], min1: min1, max1: max1, min2: min2, max2: max2}
			constraints = append(constraints, c)
		}

		if strings.Contains(line, ",") {
			t := ticket{}
			for _, v := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(v)
				t.values = append(t.values, value)
			}
			tickets = append(tickets, t)
		}
	}

	return constraints, tickets
}

func findValidTickets(constraints []constraint, tickets []ticket) ([]ticket, int) {
	var valids []ticket
	sumInvalid := 0

	for _, t := range tickets {
		ticketValid := true
		for _, v := range t.values {
			valueValid := false
			for _, c := range constraints {
				if c.isValid(v) {
					valueValid = true
					break
				}
			}
			if !valueValid {
				sumInvalid += v
				ticketValid = false
				break
			}
		}
		if ticketValid {
			valids = append(valids, t)
		}
	}

	return valids, sumInvalid
}

func determineFieldsWithPositionsPossibles(constraints []constraint, tickets []ticket, nbFields int) []fieldWithPositions {
	mapFieldsWithPositions := make(map[string][]int, nbFields)
	for position := 0; position < nbFields; position++ {
		for _, c := range constraints {
			validForAllTickets := true
			for _, t := range tickets {
				if !c.isValid(t.values[position]) {
					validForAllTickets = false
					break
				}
			}
			if validForAllTickets {
				mapFieldsWithPositions[c.field] = append(mapFieldsWithPositions[c.field], position)
			}
		}
	}

	fieldsWithPositions := make([]fieldWithPositions, nbFields)
	for field, positions := range mapFieldsWithPositions {
		fieldsWithPositions = append(fieldsWithPositions, fieldWithPositions{field, positions})
	}
	return fieldsWithPositions
}

func findSolution(fieldsWithPositions []fieldWithPositions, nbFields int, pendingSolution map[string]int) bool {
	positionsUsed := sets.NewSet()
	for _, pos := range pendingSolution {
		positionsUsed.Add(pos)
	}

	for _, fieldWithPositions := range fieldsWithPositions {
		_, fieldUsed := pendingSolution[fieldWithPositions.field]
		if !fieldUsed {
			for _, position := range fieldWithPositions.positions {
				if !positionsUsed.Contains(position) {
					pendingSolution[fieldWithPositions.field] = position
					if len(pendingSolution) == nbFields || findSolution(fieldsWithPositions, nbFields, pendingSolution) {
						return true
					}
					delete(pendingSolution, fieldWithPositions.field)
				}
			}
		}
	}
	return false
}
