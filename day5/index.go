package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	topCrates := getCratesOnTopOfEachStack("input.txt")
	fmt.Printf("Top Crates: %v\n", topCrates)
	topCrates = getCratesOnTopOfEachStackV2("input.txt")
	fmt.Printf("Top Crates(m): %v\n", topCrates)
}

func getCratesOnTopOfEachStack(filePath string) string {
	contents := getFileContents(filePath)
	stacks, procedure := getCratesAndProcedure(string(contents))
	stacks = processProcedures(procedure, stacks)
	return getTopCrates(stacks)
}

func getCratesOnTopOfEachStackV2(filePath string) string {
	contents := getFileContents(filePath)
	stacks, procedure := getCratesAndProcedure(string(contents))
	stacks = processProceduresV2(procedure, stacks)
	return getTopCrates(stacks)
}

func getFileContents(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func getCratesAndProcedure(data string) ([]stack, [][]int) {
	dataArray := strings.Split(data, "\n\n")
	length := len(dataArray)
	if length != 2 {
		panic("expects 2 parts separated by \\n\\n len:" + strconv.Itoa(length))
	}
	return createCrates(dataArray[0]), getProcedure(dataArray[1])
}

func createCrates(data string) []stack {
	var stacks []stack
	rows := strings.Split(data, "\n")
	lastRow := rows[len(rows)-1]

	for j := 1; j <= len(lastRow)-1; j += 4 {
		stack := stack{}
		for i := len(rows) - 2; i >= 0; i-- {
			crate := string(rows[i][j])
			if crate == " " {
				break
			}
			stack.crates = append(stack.crates, crate)
		}
		stacks = append(stacks, stack)
	}
	return stacks
}

func getProcedure(data string) [][]int {
	var procedure [][]int
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		formatted := strings.ReplaceAll(row, "move", "")
		formatted = strings.ReplaceAll(formatted, "from", "")
		formatted = strings.ReplaceAll(formatted, "to", "")
		formatted = strings.TrimSpace(formatted)
		formatted = removeDoubleSpace(formatted)
		columns := strings.Split(formatted, " ")
		columnLength := len(columns)
		if columnLength != 3 {
			panic("expects 3 parts separated by space. length:" + strconv.Itoa(columnLength) + " original:" + row + " formatted:" + formatted)
		}
		var columnInt []int
		for _, v := range columns {
			j, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			columnInt = append(columnInt, j)
		}
		procedure = append(procedure, columnInt)
	}
	return procedure
}

func removeDoubleSpace(data string) string {
	formatted := strings.ReplaceAll(data, "  ", " ")
	if data == formatted {
		return formatted
	}
	return removeDoubleSpace(formatted)
}

func processProcedures(procedure [][]int, stacks []stack) []stack {
	for _, row := range procedure {
		rowLength := len(row)
		if rowLength != 3 {
			panic("procedure expects 3 parts. parts:" + strconv.Itoa(rowLength))
		}
		stacks = process1(row[0], row[1], row[2], stacks)
	}
	return stacks
}

func processProceduresV2(procedure [][]int, stacks []stack) []stack {
	for _, row := range procedure {
		rowLength := len(row)
		if rowLength != 3 {
			panic("procedure expects 3 parts. parts:" + strconv.Itoa(rowLength))
		}
		stacks = process1V2(row[0], row[1], row[2], stacks)
	}
	return stacks
}

func getTopCrates(stacks []stack) string {
	var topCrates string
	for _, stack := range stacks {
		top := stack.crates[len(stack.crates)-1]
		topCrates = topCrates + top
	}
	return topCrates
}

func process1(cratesToMoveCount int, stackToMoveFrom int, stackToMoveTo int, stacks []stack) []stack {
	for i := cratesToMoveCount; i > 0; i-- {
		crateToMove := stacks[stackToMoveFrom-1].crates[len(stacks[stackToMoveFrom-1].crates)-1]
		stacks[stackToMoveFrom-1].crates = stacks[stackToMoveFrom-1].crates[0 : len(stacks[stackToMoveFrom-1].crates)-1]
		stacks[stackToMoveTo-1].crates = append(stacks[stackToMoveTo-1].crates, crateToMove)
	}
	return stacks
}

func process1V2(cratesToMoveCount int, stackToMoveFrom int, stackToMoveTo int, stacks []stack) []stack {
	var cratesToMove []string
	for i := cratesToMoveCount; i > 0; i-- {
		cratesToMove = append(cratesToMove, stacks[stackToMoveFrom-1].crates[len(stacks[stackToMoveFrom-1].crates)-1])
		stacks[stackToMoveFrom-1].crates = stacks[stackToMoveFrom-1].crates[0 : len(stacks[stackToMoveFrom-1].crates)-1]
	}
	for i := len(cratesToMove) - 1; i >= 0; i-- {
		stacks[stackToMoveTo-1].crates = append(stacks[stackToMoveTo-1].crates, cratesToMove[i])
	}
	return stacks
}

type stack struct {
	crates []string
}
