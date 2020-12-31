package main

import (
	"../../utils/files"
	"fmt"
	"sort"
	"strings"
)

const filename = "day21.txt"

func main() {
	inputs := files.ReadLinesOfFile(filename)
	fmt.Printf("Part 1: %d\n", part1(inputs))
	fmt.Printf("Part 2: %s\n", part2(inputs))
}

func part1(inputs []string) int {
	ingredientsDict, allergensDict := parseInputs(inputs)
	ingredientsWithoutAllergens := findIngredientsWithoutAllergens(ingredientsDict, allergensDict)

	result := 0
	for _, ingredient := range ingredientsWithoutAllergens {
		result += ingredient.appearance
	}
	return result
}

func part2(inputs []string) string {
	_, allergensDict := parseInputs(inputs)
	ingredientsNamesByAllergenName := computeIngredientsNamesByAllergenName(allergensDict)

	var sortedAllergensNames []string
	for allergenName := range allergensDict {
		sortedAllergensNames = append(sortedAllergensNames, allergenName)
	}
	sort.Strings(sortedAllergensNames)

	var ingredientsNames []string
	for _, allergenName := range sortedAllergensNames {
		ingredientsNames = append(ingredientsNames, ingredientsNamesByAllergenName[allergenName])
	}

	return strings.Join(ingredientsNames, ",")
}

type Ingredient struct {
	name       string
	appearance int
}

type Allergen struct {
	name        string
	ingredients []string
}

func parseInputs(inputs []string) (map[string]*Ingredient, map[string]*Allergen) {
	ingredientsDict := make(map[string]*Ingredient)
	allergensDict := make(map[string]*Allergen)

	for _, food := range inputs {
		splitted := strings.Split(food, " (contains ")
		allegensList := strings.ReplaceAll(splitted[1], ")", "")

		ingredients := strings.Split(splitted[0], " ")
		allergens := strings.Split(allegensList, ", ")

		for _, ingredientName := range ingredients {
			ingredient, found := ingredientsDict[ingredientName]
			if found {
				ingredient.appearance += 1
			} else {
				ingredient = &Ingredient{ingredientName, 1}
				ingredientsDict[ingredientName] = ingredient
			}
		}
		for _, allergenName := range allergens {
			allergen, found := allergensDict[allergenName]
			if found {
				allergen.ingredients = ingredientsIntersection(allergen.ingredients, ingredients)
			} else {
				allergen = &Allergen{allergenName, ingredients}
				allergensDict[allergenName] = allergen
			}
		}
	}

	return ingredientsDict, allergensDict
}

func findIngredientsWithoutAllergens(ingredientsDict map[string]*Ingredient, allergensDict map[string]*Allergen) []*Ingredient {
	var ingredients []*Ingredient
	for ingredientName, ingredient := range ingredientsDict {
		containsAllergen := false
		for _, allergen := range allergensDict {
			if allergen.containsIngredient(ingredientName) {
				containsAllergen = true
				break
			}
		}
		if !containsAllergen {
			ingredients = append(ingredients, ingredient)
		}
	}
	return ingredients
}

func computeIngredientsNamesByAllergenName(allergensDict map[string]*Allergen) map[string]string {
	ingredientsFound := make(map[string]bool)
	ingredientsNamesByAllergenName := make(map[string]string)

	for {
		solutionFound := true
		for allergenName, allergen := range allergensDict {
			var updatedIngredients []string
			for _, ingredient := range allergen.ingredients {
				_, found := ingredientsFound[ingredient]
				if !found {
					updatedIngredients = append(updatedIngredients, ingredient)
				}
			}
			allergen.ingredients = updatedIngredients

			if len(updatedIngredients) == 1 {
				ingredientName := updatedIngredients[0]
				ingredientsFound[ingredientName] = true
				ingredientsNamesByAllergenName[allergenName] = ingredientName
			}
			if len(updatedIngredients) > 1 {
				solutionFound = false
			}
		}
		if solutionFound {
			break
		}
	}
	return ingredientsNamesByAllergenName
}

func ingredientsIntersection(ingredients1 []string, ingredients2 []string) []string {
	var intersection []string

	for _, ingredient1 := range ingredients1 {
		for _, ingredient2 := range ingredients2 {
			if ingredient1 == ingredient2 {
				intersection = append(intersection, ingredient1)
			}
		}
	}

	return intersection
}

func (a *Allergen) containsIngredient(ingredientName string) bool {
	for _, name := range a.ingredients {
		if name == ingredientName {
			return true
		}
	}
	return false
}
