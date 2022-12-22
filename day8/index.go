package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
}

func getFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func countVisibleTrees(contents string) int {
	forest := convertToMatrix(contents)
	markVisible(forest)
	visibleTrees := 0
	for _, row := range forest {
		for _, tree := range row {
			if !tree.isVisible {
				continue
			}
			visibleTrees++
		}
	}
	return visibleTrees
}

func getHighestScenicScore(contents string) int {
	forest := convertToMatrix(contents)
	setScenicScores(forest)
	highestScenicScore := 0
	for _, row := range forest {
		for _, tree := range row {
			if tree.scenicScore > highestScenicScore {
				highestScenicScore = tree.scenicScore
			}
		}
	}
	return highestScenicScore
}

func setScenicScores(forest [][]*tree) {
	for i := 1; i < len(forest); i++ {
		for j := 1; j < len(forest[i]); j++ {
			tree := forest[i][j]
			topScore := 0
			for ts := i - 1; ts >= 0; ts-- {
				topTree := forest[ts][j]
				topScore++
				if topTree.height >= tree.height {
					break
				}
			}
			bottomScore := 0
			for bs := i + 1; bs < len(forest); bs++ {
				bottomTree := forest[bs][j]
				bottomScore++
				if bottomTree.height >= tree.height {
					break
				}
			}
			rightScore := 0
			for rs := j + 1; rs < len(forest[i]); rs++ {
				rightTree := forest[i][rs]
				rightScore++
				if rightTree.height >= tree.height {
					break
				}
			}
			leftScore := 0
			for ls := j - 1; ls >= 0; ls-- {
				leftTree := forest[i][ls]
				leftScore++
				if leftTree.height >= tree.height {
					break
				}
			}
			tree.scenicScore = topScore * bottomScore * rightScore * leftScore
		}
	}
}

func convertToMatrix(data string) [][]*tree {
	rows := strings.Split(data, "\n")
	rowsLength := len(rows)
	var forest [][]*tree
	for _, row := range rows {
		var trees []*tree
		if len(row) != rowsLength {
			panic("invalid column length")
		}
		for i := 0; i < len(row); i++ {
			treeHeight, err := strconv.Atoi(row[i : i+1])
			if err != nil {
				panic("failed to convert string to int")
			}
			tree := tree{height: treeHeight, c: len(trees), r: len(forest)}
			trees = append(trees, &tree)
		}
		forest = append(forest, trees)
	}
	return forest
}

func markVisible(forest [][]*tree) {
	columnLength := len(forest[0])
	for r, row := range forest {
		if len(row) != columnLength {
			panic("column length is ")
		}
		for c, column := range row {
			if isVisibleFromTop(r, c, forest) || isVisibleOnLeft(r, c, forest) || isVisibleFromBottom(r, c, forest) || isVisibleFromRight(r, c, forest) {
				column.isVisible = true
			}
		}
	}
}

func isVisibleOnLeft(r int, c int, forest [][]*tree) bool {
	if c == 0 {
		return true
	}
	tree := forest[r][c]
	for i := 0; i < c; i++ {
		currentTree := forest[r][i]
		if currentTree.height >= tree.height {
			return false
		}
	}
	return true
}

func isVisibleFromRight(r int, c int, forest [][]*tree) bool {
	if c == len(forest[r])-1 {
		return true
	}
	tree := forest[r][c]
	for i := len(forest[r]) - 1; i > c; i-- {
		currentTree := forest[r][i]
		if currentTree.height >= tree.height {
			return false
		}
	}
	return true
}

func isVisibleFromTop(r int, c int, forest [][]*tree) bool {
	if r == 0 {
		return true
	}
	tree := forest[r][c]
	for i := 0; i < r; i++ {
		currentTree := forest[i][c]
		if currentTree.height >= tree.height {
			return false
		}
	}
	return true
}

func isVisibleFromBottom(r int, c int, forest [][]*tree) bool {
	if r == len(forest)-1 {
		return true
	}
	tree := forest[r][c]
	for i := len(forest) - 1; i > r; i-- {
		currentTree := forest[i][c]
		if currentTree.height >= tree.height {
			return false
		}
	}
	return true
}

type tree struct {
	isVisible   bool
	height      int
	scenicScore int
	r           int
	c           int
}
