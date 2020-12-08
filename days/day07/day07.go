package main

import (
	"fmt"
	"regexp"
	"strconv"

	"../../utils/files"
)

const filename = "day07.txt"
const shinyGold = "shiny gold"

type Color string

type Rule struct {
	color  Color
	number int
}

type Graph map[Color][]Rule

func main() {
	inputs := files.ReadLinesOfFile(filename)

	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	graph := buildGraph(inputs)
	nbBags := 0
	for color := range graph {
		if graph.CanContains(color, shinyGold) {
			nbBags++
		}
	}
	return nbBags
}

func part2(inputs []string) int {
	graph := buildGraph(inputs)
	return graph.CountBags(shinyGold)
}

func buildGraph(inputs []string) Graph {
	graph := make(Graph)
	regexLine := regexp.MustCompile("^([a-z]+ [a-z]+) bags contain (.*)\\.$")
	regexRule := regexp.MustCompile("([0-9]) ([a-z]+ [a-z]+) bag")
	for _, input := range inputs {
		lineMatches := regexLine.FindStringSubmatch(input)
		lineColor := Color(lineMatches[1])
		ruleMatches := regexRule.FindAllStringSubmatch(lineMatches[2], -1)
		for _, rule := range ruleMatches {
			ruleColor := Color(rule[2])
			ruleNumber, _ := strconv.Atoi(rule[1])
			graph[lineColor] = append(graph[lineColor], Rule{color: ruleColor, number: ruleNumber})
		}
	}
	return graph
}

func (graph Graph) CanContains(color Color, expected Color) bool {
	for _, rule := range graph[color] {
		if rule.color == expected {
			return true
		}
		if graph.CanContains(rule.color, expected) {
			return true
		}
	}
	return false
}

func (graph Graph) CountBags(color Color) int {
	count := 0
	for _, rule := range graph[color] {
		count += rule.number * (1 + graph.CountBags(rule.color))
	}
	return count
}
