package main

import (
	"../../utils/files"
	"fmt"
	"regexp"
	"strconv"
)

const filename = "day14.txt"

var maskRegex = regexp.MustCompile(`^mask = ([01X]*)$`)
var memRegex = regexp.MustCompile(`^mem\[(\d*)] = (\d*)$`)

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", execute(inputs, true))
	fmt.Printf("Part 2: %d\n", execute(inputs, false))
}

func execute(inputs []string, part1 bool) int {
	memory := newMemory()

	for _, input := range inputs {
		matches := maskRegex.FindStringSubmatch(input)
		if matches != nil {
			memory.mask = matches[1]
		} else {
			matches = memRegex.FindStringSubmatch(input)
			if matches == nil {
				panic("Unsupported line")
			}
			address, _ := strconv.Atoi(matches[1])
			value, _ := strconv.Atoi(matches[2])
			if part1 {
				memory.setValuePart1(address, value)
			} else {
				memory.setValuePart2(address, value)
			}
		}
	}

	return memory.sum()
}

type memory struct {
	mask   string
	values map[int]int
}

func newMemory() *memory {
	m := &memory{}
	m.values = make(map[int]int)
	return m
}

func (m *memory) sum() int {
	sum := 0
	for _, v := range m.values {
		sum += v
	}
	return sum
}

func (m *memory) setValuePart1(address int, value int) {
	original := intToBinary(value, len(m.mask))
	binary := ""

	for i, b := range m.mask {
		if b == 'X' {
			binary += string(original[i])
		} else {
			binary += string(b)
		}
	}

	m.values[address] = binaryToInt(binary)
}

func (m *memory) setValuePart2(address int, value int) {
	original := intToBinary(address, len(m.mask))
	binaries := []string{""}

	for i, b := range m.mask {
		var nextBinaries []string

		for _, binary := range binaries {
			if b == 'X' {
				nextBinaries = append(nextBinaries, binary+"0")
				nextBinaries = append(nextBinaries, binary+"1")
			} else if b == '1' {
				nextBinaries = append(nextBinaries, binary+"1")
			} else {
				nextBinaries = append(nextBinaries, binary+string(original[i]))
			}
		}

		binaries = nextBinaries
	}

	for _, binary := range binaries {
		m.values[binaryToInt(binary)] = value
	}
}

func intToBinary(value int, size int) string {
	binary := strconv.FormatInt(int64(value), 2)
	for len(binary) < size {
		binary = "0" + binary
	}
	return binary
}

func binaryToInt(value string) int {
	i, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		panic("Cannot convert binary to int")
	}
	return int(i)
}
