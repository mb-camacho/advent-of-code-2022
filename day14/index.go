package main

import (
	"os"
	"strconv"
	"strings"
)

const maxX = 700
const maxY = 510

func main() {
}

func getFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func countUnitsOfSandWithFloor(data string) int {
	caveMap := buildMap()
	_, highestY := addRocks(caveMap, data)
	addFloor(caveMap, highestY+2)
	sandCoord := [2]int{500, 0}
	countSandAdded := 0
	for {
		if !processFallingSandV2(caveMap, sandCoord) {
			break
		}
		countSandAdded++
	}
	return countSandAdded
}

func processFallingSandV2(caveMap [][]*rune, sandCoord [2]int) bool {
	currentX := sandCoord[0]
	currentY := sandCoord[1]
	for {
		if *caveMap[currentX][currentY] == 'o' {
			return false
		}

		// down
		if *caveMap[currentX][currentY+1] == '.' {
			currentY++
			continue
		}

		// down left
		if *caveMap[currentX-1][currentY+1] == '.' {
			currentX--
			currentY++
			continue
		}

		// down right
		if *caveMap[currentX+1][currentY+1] == '.' {
			currentX++
			currentY++
			continue
		}

		settledSand := 'o'
		caveMap[currentX][currentY] = &settledSand
		return true
	}
}

func addFloor(caveMap [][]*rune, floor int) {
	for x := 0; x <= maxX+1; x++ {
		rock := '#'
		caveMap[x][floor] = &rock
	}
}

func countUnitsOfSand(data string) int {
	caveMap := buildMap()
	addRocks(caveMap, data)
	sandCoord := [2]int{500, 0}
	countSandAdded := 0
	for {
		if !processFallingSand(caveMap, sandCoord) {
			break
		}
		countSandAdded++
	}
	return countSandAdded
}

func processFallingSand(caveMap [][]*rune, sandCoord [2]int) bool {
	currentX := sandCoord[0]
	currentY := sandCoord[1]
	for {
		if currentY+1 > maxY {
			return false
		}

		// down
		if *caveMap[currentX][currentY+1] == '.' {
			currentY++
			continue
		}

		// down left
		if *caveMap[currentX-1][currentY+1] == '.' {
			currentX--
			currentY++
			continue
		}

		// down right
		if *caveMap[currentX+1][currentY+1] == '.' {
			currentX++
			currentY++
			continue
		}

		settledSand := 'o'
		caveMap[currentX][currentY] = &settledSand
		return true
	}
}

func addRocks(caveMap [][]*rune, data string) (int, int) {
	rocksAdded := 0
	rows := strings.Split(data, "\n")
	highestY := 0
	for _, row := range rows {
		coordinatePairs := strings.Split(row, " -> ")
		var lastRock []int
		for _, coorcoordinatePair := range coordinatePairs {
			coordinates := strings.Split(coorcoordinatePair, ",")
			var coordinatesInt []int
			if len(coordinates) != 2 {
				panic("invalid coordinates: " + coorcoordinatePair)
			}
			for _, coordinate := range coordinates {
				intValue, error := strconv.Atoi(coordinate)
				if error != nil {
					panic(error)
				}
				coordinatesInt = append(coordinatesInt, intValue)
			}
			if coordinatesInt[0] > maxX {
				panic("maxX should be increased to " + coordinates[0])
			}
			if coordinatesInt[1] > maxY {
				panic("maxY should be increased to " + coordinates[1])
			}
			if coordinatesInt[1] > highestY {
				highestY = coordinatesInt[1]
			}
			if len(lastRock) == 0 {
				if *caveMap[coordinatesInt[0]][coordinatesInt[1]] != '#' {
					rock := '#'
					rocksAdded++
					caveMap[coordinatesInt[0]][coordinatesInt[1]] = &rock
				}
			} else {
				rocksAdded += addLineOfRocks(caveMap, lastRock, coordinatesInt)
			}
			lastRock = coordinatesInt
		}
	}
	return rocksAdded, highestY
}

func addLineOfRocks(caveMap [][]*rune, startCoord []int, endCoord []int) int {
	rocksAdded := 0
	lastRock := startCoord
	for {

		if *caveMap[lastRock[0]][lastRock[1]] != '#' {
			rock := '#'
			caveMap[lastRock[0]][lastRock[1]] = &rock
			rocksAdded++
		}

		if lastRock[0] == endCoord[0] && lastRock[1] == endCoord[1] {
			break
		}

		if lastRock[0] < endCoord[0] {
			lastRock[0]++
		} else if lastRock[0] > endCoord[0] {
			lastRock[0]--
		}
		if lastRock[1] < endCoord[1] {
			lastRock[1]++
		} else if lastRock[1] > endCoord[1] {
			lastRock[1]--
		}
	}
	return rocksAdded
}

func buildMap() [][]*rune {
	var caveMap [][]*rune
	for i := 0; i <= maxX+1; i++ {
		var caveMapRow []*rune
		for j := 0; j <= maxY+1; j++ {
			char := '.'
			caveMapRow = append(caveMapRow, &char)
		}
		caveMap = append(caveMap, caveMapRow)
	}
	return caveMap
}
