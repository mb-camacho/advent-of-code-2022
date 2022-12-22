package main

import (
	"fmt"
)

func main() {
	// day 1
	maxTotal := getMax("input.txt")
	fmt.Printf("Max Total: %d\n", maxTotal)
	top3 := getSumTop3("input.txt")
	fmt.Printf("Max Total: %d\n", top3)

	// day 2
	score := getScore("input-day2.txt")
	fmt.Printf("Total Score: %d\n", score)
	score2 := getScoreV2("input-day2.txt")
	fmt.Printf("Total Score V2: %d\n", score2)

	// day 3
	sum := getSumOfPriorities("input-day3.txt")
	fmt.Printf("Sum of Priorities: %d\n", sum)
	sumGroup := getSumOfGroupPriorities("input-day3.txt")
	fmt.Printf("Sum of Group Priorities: %d\n", sumGroup)

	// day 4
	count := getCountOfAssignmentPairFullyContainingTheOther("input-day4.txt")
	fmt.Printf("Count of Assignement Pairs fully containing the other: %d\n", count)
	count2 := getCountOfAssignmentPairOverlapping("input-day4.txt")
	fmt.Printf("Count of Assignement Pairs partially containing the other: %d\n", count2)
}
