package main

import (
	"fmt"
	"sync"

	"../../utils/files"
)

const filename = "day03.txt"
const treeChar = '#'

func main() {
	values := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(values))
	fmt.Printf("Part 2: %d\n", part2WithGoroutines(values))
}

func part1(inputs []string) int {
	slope := Slope{right: 3, down: 1}
	return slope.countTrees(inputs)
}

func part2(inputs []string) int {
	slopes := []Slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	total := 1
	for _, slope := range slopes {
		total *= slope.countTrees(inputs)
	}
	return total
}

func part2WithGoroutines(inputs []string) int {
	slopes := []Slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	results := make(chan int, len(slopes))
	wg := new(sync.WaitGroup)
	for _, slope := range slopes {
		wg.Add(1)
		go slope.countTreesAsync(inputs, wg, results)
	}

	wg.Wait()
	close(results)

	total := 1
	for result := range results {
		total *= result
	}
	return total
}

type Slope struct {
	right, down int
}

func (slope Slope) countTrees(inputs []string) int {
	nbTrees := 0
	lineSize := len(inputs[0])
	top, left := 0, 0
	for top < len(inputs)-1 {
		top += slope.down
		left = (left + slope.right) % lineSize
		if inputs[top][left] == treeChar {
			nbTrees++
		}
	}
	return nbTrees
}

func (slope Slope) countTreesAsync(inputs []string, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	results <- slope.countTrees(inputs)
}
