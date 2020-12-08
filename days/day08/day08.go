package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../utils/files"
	"../../utils/sets"
)

const filename = "day08.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)

	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	handheldConsole := buildHandheldConsole(inputs)
	accumulator, status := handheldConsole.execute()
	if status == loop {
		return accumulator
	}
	panic("Wrong status")
}

func part2(inputs []string) int {
	handheldConsole := buildHandheldConsole(inputs)

	for index, instruction := range handheldConsole.instructions {
		if instruction.operation == jumpOperation || instruction.operation == noOperation {
			originalOperation := instruction.operation
			if originalOperation == jumpOperation {
				handheldConsole.updateInstructionOperation(index, noOperation)
			} else {
				handheldConsole.updateInstructionOperation(index, jumpOperation)
			}

			accumulator, end := handheldConsole.execute()
			if end == success {
				return accumulator
			}

			handheldConsole.updateInstructionOperation(index, originalOperation)
		}
	}
	panic("Cannot found a solution")
}

type handheldConsole struct {
	instructions map[int]*instruction
}

type instruction struct {
	operation operation
	argument  int
}

type operation string

const accumulatorOperation = operation("acc")
const jumpOperation = operation("jmp")
const noOperation = operation("nop")

type executionStatus string

const loop = executionStatus("LOOP")
const success = executionStatus("SUCCESS")

func buildHandheldConsole(inputs []string) handheldConsole {
	handheldConsole := handheldConsole{}
	handheldConsole.instructions = make(map[int]*instruction)

	for index, input := range inputs {
		fields := strings.Fields(input)
		operation := operation(fields[0])
		argument, _ := strconv.Atoi(fields[1])
		handheldConsole.instructions[index] = &instruction{operation: operation, argument: argument}
	}

	return handheldConsole
}

func (hc *handheldConsole) execute() (int, executionStatus) {
	index := 0
	accumulator := 0
	indexExecuted := sets.NewSet()

	for {
		currentIndex := index
		if currentIndex == len(hc.instructions) {
			return accumulator, success
		}

		instruction := hc.instructions[currentIndex]
		if indexExecuted.Contains(currentIndex) {
			return accumulator, loop
		}

		switch instruction.operation {
		case accumulatorOperation:
			accumulator += instruction.argument
			index++
		case jumpOperation:
			index += instruction.argument
		case noOperation:
			index++
		default:
			panic("Operation not supported")
		}

		indexExecuted.Add(currentIndex)
	}
}

func (hc *handheldConsole) updateInstructionOperation(index int, operation operation) {
	hc.instructions[index].operation = operation
}
