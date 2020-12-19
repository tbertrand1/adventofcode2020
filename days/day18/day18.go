package main

import (
	"../../utils/files"
	"fmt"
	"strings"
)

const filename = "day18.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += evalaluateInputPart1(input)
	}
	return sum
}

func part2(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += evalaluateInputPart2(input)
	}
	return sum
}

func evalaluateInputPart1(input string) int {
	result, _ := evaluateValueFromIndex(strings.ReplaceAll(input, " ", ""), 0, false)
	return result
}

func evalaluateInputPart2(input string) int {
	result, _ := evaluateValueFromIndex(strings.ReplaceAll(input, " ", ""), 0, true)
	return result
}

type operator string

const (
	empty   operator = ""
	sum              = "+"
	product          = "*"
)

type operation struct {
	operator operator
	value    int
}

func evaluateValueFromIndex(input string, index int, sumBeforeProduct bool) (int, int) {
	var operations []operation
	op := operation{operator: empty}

	for i := index; i < len(input); i++ {
		char := input[i]
		switch char {
		case '(':
			op.value, i = evaluateValueFromIndex(input, i+1, sumBeforeProduct)
			operations = append(operations, op)
			op = operation{operator: empty}
		case ')':
			return calculateValue(operations, sumBeforeProduct), i
		case '+':
			op.operator = sum
		case '*':
			op.operator = product
		default:
			op.value = int(char - '0')
			operations = append(operations, op)
			op = operation{operator: empty}
		}
	}

	return calculateValue(operations, sumBeforeProduct), len(input) - 1
}

func calculateValue(operations []operation, sumBeforeProduct bool) int {
	if len(operations) == 0 {
		panic("illegal state")
	}
	if sumBeforeProduct {
		return calculateValueWithSumBeforeProduct(operations)
	} else {
		return calculateValueFromLeftToRight(operations)
	}
}

func calculateValueFromLeftToRight(operations []operation) int {
	value := operations[0].value
	for _, op := range operations[1:] {
		if op.operator == sum {
			value += op.value
		} else if op.operator == product {
			value *= op.value
		} else {
			panic("operator not supported")
		}
	}
	return value
}

func calculateValueWithSumBeforeProduct(operations []operation) int {
	var remainingOperations []operation

	// Apply sum operator first
	for _, op := range operations {
		if op.operator == sum {
			if len(remainingOperations) == 0 {
				panic("illegal state")
			}
			previousPosition := len(remainingOperations) - 1
			previousOp, remainningOperations := remainingOperations[previousPosition], remainingOperations[:previousPosition]
			newOp := operation{operator: previousOp.operator, value: previousOp.value + op.value}
			remainningOperations = append(remainningOperations, newOp)
		} else {
			remainingOperations = append(remainingOperations, op)
		}
	}

	// Then apply remaining operations : product
	value := remainingOperations[0].value
	for _, op := range remainingOperations[1:] {
		if op.operator == product {
			value *= op.value
		} else {
			panic("operator not supported")
		}
	}
	return value
}
