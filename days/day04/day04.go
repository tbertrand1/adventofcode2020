package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../../utils/files"
)

const filename = "day04.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", checkPassports(inputs, 1))
	fmt.Printf("Part 2: %d\n", checkPassports(inputs, 2))
}

func checkPassports(inputs []string, part int) int {
	valids := 0

	currentPassport := make(map[string]string)

	for _, input := range inputs {
		fields := strings.Fields(input)
		if len(fields) == 0 {
			if isPassportValid(currentPassport, part) {
				valids++
			}
			currentPassport = make(map[string]string)
		} else {
			regex := regexp.MustCompile("([a-z]+):([^ ]+)")
			for _, field := range fields {
				matches := regex.FindStringSubmatch(field)
				currentPassport[matches[1]] = matches[2]
			}
		}
	}
	if isPassportValid(currentPassport, part) {
		valids++
	}

	return valids
}

func isPassportValid(passport map[string]string, part int) bool {
	if part == 1 {
		return isPassportValidPart1(passport)
	}
	return isPassportValidPart2(passport)
}

func isPassportValidPart1(passport map[string]string) bool {
	_, hasByr := passport["byr"]
	_, hasIyr := passport["iyr"]
	_, hasEyr := passport["eyr"]
	_, hasHgt := passport["hgt"]
	_, hasHcl := passport["hcl"]
	_, hasEcl := passport["ecl"]
	_, hasPid := passport["pid"]
	return hasByr && hasIyr && hasEyr && hasHgt &&
		hasHcl && hasEcl && hasPid
}

func isPassportValidPart2(passport map[string]string) bool {
	byr, hasByr := passport["byr"]
	iyr, hasIyr := passport["iyr"]
	eyr, hasEyr := passport["eyr"]
	hgt, hasHgt := passport["hgt"]
	hcl, hasHcl := passport["hcl"]
	ecl, hasEcl := passport["ecl"]
	pid, hasPid := passport["pid"]

	if !hasByr || !hasIyr || !hasEyr || !hasHgt ||
		!hasHcl || !hasEcl || !hasPid {
		return false
	}

	return isIntBetween(byr, 1920, 2002) &&
		isIntBetween(iyr, 2010, 2020) &&
		isIntBetween(eyr, 2020, 2030) &&
		isHeightValid(hgt) &&
		isHairColor(hcl) &&
		isEyeColor(ecl) &&
		isPid(pid)
}

func isIntBetween(value string, min int, max int) bool {
	intValue, err := strconv.Atoi(value)
	return err == nil && intValue >= min && intValue <= max
}

func isHeightValid(value string) bool {
	unit := value[len(value)-2:]
	height := value[:len(value)-2]
	if unit == "cm" {
		return isIntBetween(height, 150, 193)
	} else if unit == "in" {
		return isIntBetween(height, 59, 76)
	}

	return false
}

func isHairColor(value string) bool {
	return regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(value)
}

func isEyeColor(value string) bool {
	return value == "amb" || value == "blu" || value == "brn" ||
		value == "gry" || value == "grn" || value == "hzl" || value == "oth"
}

func isPid(value string) bool {
	return regexp.MustCompile("^[0-9]{9}$").MatchString(value)
}
