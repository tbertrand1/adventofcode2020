package main

import (
	"../../utils/files"
	"../../utils/str"
	"fmt"
	"strings"
)

const filename = "day19.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %d\n", part2(inputs))
}

func part1(inputs []string) int {
	rules, messages := parseInputs(inputs)

	valids := 0
	for _, message := range messages {
		if isValid(message, []int{0}, rules) {
			valids++
		}
	}
	return valids
}

func part2(inputs []string) int {
	rules, messages := parseInputs(inputs)

	rule8 := rules[8]
	rule8.orRulesIds = [][]int{{42}, {42, 8}}
	rules[8] = rule8
	rule11 := rules[11]
	rule11.orRulesIds = [][]int{{42, 31}, {42, 11, 31}}
	rules[11] = rule11

	valids := 0
	for _, message := range messages {
		if isValid(message, []int{0}, rules) {
			valids++
		}
	}
	return valids
}

func parseInputs(inputs []string) (dictRules, []string) {
	rules := make(dictRules)
	var messages []string

	for _, input := range inputs {
		ruleSplit := strings.Split(input, ": ")
		if len(ruleSplit) == 2 {
			rule := parseRule(str.Atoi(ruleSplit[0]), ruleSplit[1])
			rules[rule.id] = rule
		} else {
			if input != "" {
				messages = append(messages, input)
			}
		}
	}

	return rules, messages
}

func parseRule(id int, content string) rule {
	rule := rule{id: id}
	if strings.Contains(content, "\"") {
		rule.literal = strings.ReplaceAll(content, "\"", "")
	} else {
		for _, orContent := range strings.Split(content, " | ") {
			var orRulesIds []int
			for _, ruleId := range strings.Split(orContent, " ") {
				orRulesIds = append(orRulesIds, str.Atoi(ruleId))
			}
			rule.orRulesIds = append(rule.orRulesIds, orRulesIds)
		}
	}
	return rule
}

type rule struct {
	id         int
	literal    string
	orRulesIds [][]int
}

type dictRules map[int]rule

func isValid(message string, stack []int, rules dictRules) bool {
	if len(stack) == 0 || len(message) == 0 {
		return len(stack) == 0 && len(message) == 0
	}

	id, newStack := pop(stack)
	currentRule := rules[id]

	if currentRule.literal == string(message[0]) {
		return isValid(message[1:], newStack, rules)
	} else {
		for _, orRulesIds := range currentRule.orRulesIds {
			var newStackWithRules []int
			newStackWithRules = append(newStackWithRules, newStack...)
			newStackWithRules = append(newStackWithRules, reverse(orRulesIds)...)
			if isValid(message, newStackWithRules, rules) {
				return true
			}
		}
	}
	return false
}

func pop(rules []int) (int, []int) {
	return rules[len(rules)-1], rules[:len(rules)-1]
}

func reverse(rules []int) []int {
	size := len(rules)
	reversed := make([]int, size)
	for i := size - 1; i >= 0; i-- {
		reversed[i] = rules[size-1-i]
	}
	return reversed
}
