package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := readFileIntoStringArray()
	partOne(lines)
	partTwo(lines)
}

type BagType string

type Rule struct {
	TypeOfBag BagType
	Amount    int
}

func partOne(lines []string) {
	bags := getBags(lines)

	mapOfBagsThatLeadToGold := make(map[BagType]int)
	mapOfBagsThatLeadToGold["shiny gold"] = 1

	pathsRemain := true
	for pathsRemain == true {
		pathsRemain = false
	bagIteration:
		for bagType, rules := range bags {
			_, seenPath := mapOfBagsThatLeadToGold[bagType]
			if !seenPath {
				pathsRemain = true

				for _, rule := range rules {
					bagTypeLeadsToGold, seenPath := mapOfBagsThatLeadToGold[rule.TypeOfBag]
					if !seenPath {
						continue bagIteration
					}

					if bagTypeLeadsToGold == 1 {
						mapOfBagsThatLeadToGold[bagType] = 1
						continue bagIteration
					}
				}

				mapOfBagsThatLeadToGold[bagType] = 0
			}
		}
	}

	count := 0
	for _, val := range mapOfBagsThatLeadToGold {
		if val == 1 {
			count++
		}
	}

	// Subtract 1 from result for the shiny gold bag itself
	fmt.Println("Part One: ", (count - 1))
}

func partTwo(lines []string) {
	bags := getBags(lines)

	fmt.Println("Part Two: ", getCountOfBags("shiny gold", bags))
}

func getCountOfBags(bagType BagType, bagTypeToRules map[BagType][]Rule) int {
	rules := bagTypeToRules[bagType]

	if len(rules) == 0 {
		return 0
	}

	count := 0

	for _, rule := range rules {
		count += rule.Amount + rule.Amount*getCountOfBags(rule.TypeOfBag, bagTypeToRules)
	}

	return count
}

func getBags(lines []string) map[BagType][]Rule {
	bags := make(map[BagType][]Rule)
	for _, line := range lines {
		var rules []Rule
		result := strings.Split(line, "contain")

		bagTypeSplitBySpaces := strings.Split(result[0], " ")
		bagType := BagType(bagTypeSplitBySpaces[0] + " " + bagTypeSplitBySpaces[1])

		if !strings.Contains(result[1], "no other bags") {
			containedBags := strings.Split(strings.TrimSpace(result[1]), ", ")
			for _, containedBag := range containedBags {
				bagSplitBySpaces := strings.Split(containedBag, " ")

				amountOfBags, _ := strconv.Atoi(bagSplitBySpaces[0])

				rules = append(rules, Rule{
					Amount:    amountOfBags,
					TypeOfBag: BagType(bagSplitBySpaces[1] + " " + bagSplitBySpaces[2]),
				})
			}
		}

		bags[bagType] = rules
	}

	return bags
}

func readFileIntoStringArray() []string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
