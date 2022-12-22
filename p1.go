package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func getFileContents(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func convertToArray(data string) []string {
	dataArray := strings.Split(data, "\n")
	return dataArray
}

func getTotals(data []string) []int {
	var totals []int
	currentTotal := 0
	for _, s := range data {
		if s == "" {
			totals = append(totals, currentTotal)
			currentTotal = 0
			continue
		}
		value, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		currentTotal += value
	}
	totals = append(totals, currentTotal)
	sort.Ints(totals)
	return totals
}

func getMax(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	totals := getTotals(dataArray)
	return totals[len(totals)-1]
}

func getSumTop3(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	totals := getTotals(dataArray)
	top3 := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]
	return top3
}

func getResults(data []string) []int {
	var totals []int
	for _, s := range data {
		firstChar := s[0:1]
		lastChar := s[len(s)-1:]
		playScore := getPlayScore(lastChar)
		playResultScore := getPlayResultScore(lastChar, firstChar)
		totals = append(totals, playScore+playResultScore)
	}
	return totals
}

func getResultsV2(data []string) []int {
	var totals []int
	for _, s := range data {
		firstChar := s[0:1]
		lastChar := s[len(s)-1:]
		playResultScore := getPlayResultScoreV2(lastChar)
		playScore := getPlayScoreV2(firstChar, lastChar)
		totals = append(totals, playScore+playResultScore)
	}
	return totals
}

func getPlayScore(data string) int {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	if data == "X" {
		return 1
	}
	if data == "Y" {
		return 2
	}
	if data == "Z" {
		return 3
	}
	return 0
}

func getPlayScoreV2(opponent string, result string) int {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// A for Rock, B for Paper, and C for Scissors
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	if opponent == "A" {
		if result == "X" {
			return 3
		}
		if result == "Y" {
			return 1
		}
		if result == "Z" {
			return 2
		}
	}
	if opponent == "B" {
		if result == "X" {
			return 1
		}
		if result == "Y" {
			return 2
		}
		if result == "Z" {
			return 3
		}
	}
	if opponent == "C" {
		if result == "X" {
			return 2
		}
		if result == "Y" {
			return 3
		}
		if result == "Z" {
			return 1
		}
	}
	return 0
}

func getPlayResultScore(you string, opponent string) int {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors
	if you == "X" {
		if opponent == "A" {
			return 3
		}
		if opponent == "B" {
			return 0
		}
		if opponent == "C" {
			return 6
		}
	}
	if you == "Y" {
		if opponent == "A" {
			return 6
		}
		if opponent == "B" {
			return 3
		}
		if opponent == "C" {
			return 0
		}
	}
	if you == "Z" {
		if opponent == "A" {
			return 0
		}
		if opponent == "B" {
			return 6
		}
		if opponent == "C" {
			return 3
		}
	}
	return 0
}

func getPlayResultScoreV2(result string) int {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	if result == "X" {
		return 0
	}
	if result == "Y" {
		return 3
	}
	if result == "Z" {
		return 6
	}
	return 0
}

func getScore(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	results := getResults(dataArray)
	var total int
	for _, s := range results {
		total += s
	}
	return total
}

func getScoreV2(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	results := getResultsV2(dataArray)
	var total int
	for _, s := range results {
		total += s
	}
	return total
}

func getItemTypes(data []string) []rune {
	var itemTypes []rune
	for _, rucksack := range data {
		rucksackLength := len(rucksack)
		if rucksackLength%2 != 0 {
			panic("rucksack length is not even. length:" + strconv.Itoa(rucksackLength) + " rucksack:" + rucksack)
		}
		compartmentLength := rucksackLength / 2
		compartment1 := rucksack[0:compartmentLength]
		compartment2 := rucksack[compartmentLength:]
		itemType := getItemType(compartment1, compartment2)
		itemTypes = append(itemTypes, itemType)
	}
	return itemTypes
}

func getGroupItemTypes(rucksacks []string) []rune {
	rucksacksLength := len(rucksacks)
	if rucksacksLength%3 != 0 {
		panic("number of rocksacks is not diivisible by 3. count:" + strconv.Itoa(rucksacksLength))
	}
	var itemTypes []rune
	for i := 0; i < rucksacksLength-1; i += 3 {
		group1 := rucksacks[i]
		group2 := rucksacks[i+1]
		group3 := rucksacks[i+2]
		itemType := getGroupItemType(group1, group2, group3)
		itemTypes = append(itemTypes, itemType)
	}
	return itemTypes
}

func getItemType(compartment1 string, compartment2 string) rune {
	for _, char1 := range compartment1 {
		for _, char2 := range compartment2 {
			if char1 == char2 {
				return char1
			}
		}
	}
	panic("no similar item found: " + compartment1 + compartment2)
}

func getGroupItemType(group1 string, group2 string, group3 string) rune {
	for _, char1 := range group1 {
		for _, char2 := range group2 {
			if char1 != char2 {
				continue
			}
			for _, char3 := range group3 {
				if char2 == char3 {
					return char2
				}
			}
		}
	}
	panic("no similar item found: " + group1 + group2 + group3)
}

func getPriorities(itemTypes []rune) []int {
	var priorities []int
	for _, itemType := range itemTypes {
		priorities = append(priorities, getPriority(itemType))
	}
	return priorities
}

func getPriority(itemType rune) int {
	// Lowercase item types a through z have priorities 1 through 26. 97 - 122
	// Uppercase item types A through Z have priorities 27 through 52. 65 - 90
	ascii := int(itemType)
	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	}
	if ascii >= 65 && ascii <= 90 {
		return ascii - 38
	}
	panic("cannot get priority of itemType:" + string(itemType) + " ascii:" + strconv.Itoa(ascii))
}

func getSum(priorities []int) int {
	total := 0
	for _, priority := range priorities {
		total += priority
	}
	return total
}

func getSumOfPriorities(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	itemTypes := getItemTypes(dataArray)
	priorities := getPriorities(itemTypes)
	return getSum(priorities)
}

func getSumOfGroupPriorities(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	itemTypes := getGroupItemTypes(dataArray)
	priorities := getPriorities(itemTypes)
	return getSum(priorities)
}

func isFullyContained(section1 string, section2 string) bool {
	minMax1 := getMinAndMax(section1)
	minMax2 := getMinAndMax(section2)
	if minMax1[0] >= minMax2[0] && minMax1[1] <= minMax2[1] {
		return true
	}
	if minMax2[0] >= minMax1[0] && minMax2[1] <= minMax1[1] {
		return true
	}
	return false
}

func isPartiallyContained(section1 string, section2 string) bool {
	minMax1 := getMinAndMax(section1)
	minMax2 := getMinAndMax(section2)
	for i := minMax1[0]; i <= minMax1[1]; i++ {
		if i >= minMax2[0] && i <= minMax2[1] {
			return true
		}
	}
	return false
}

func getMinAndMax(section string) []int {
	minMax := strings.Split(section, "-")
	if len(minMax) != 2 {
		panic("expecting 2 parts separated by -. " + section)
	}

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		panic(err)
	}
	return []int{min, max}
}

func getCountOfAssignmentPairFullyContainingTheOther(filePath string) int {
	data := getFileContents(filePath)
	dataArray := convertToArray(string(data))
	count := 0
	for _, sectionPair := range dataArray {
		sectionPairArray := strings.Split(sectionPair, ",")
		if len(sectionPairArray) != 2 {
			panic("expecting 2 parts separated by ,. " + sectionPair)
		}
		section1 := sectionPairArray[0]
		section2 := sectionPairArray[1]
		if isFullyContained(section1, section2) {
			count++
		}
	}
	return count
}

func getCountOfAssignmentPairOverlapping(filePath string) int {
	data := getFileContents(filePath)
	secrionPairs := convertToArray(string(data))
	count := 0
	for _, sectionPair := range secrionPairs {
		sectionPairArray := strings.Split(sectionPair, ",")
		if len(sectionPairArray) != 2 {
			panic("expecting 2 parts separated by ,. " + sectionPair)
		}
		section1 := sectionPairArray[0]
		section2 := sectionPairArray[1]
		if isPartiallyContained(section1, section2) {
			count++
		}
	}
	return count
}
