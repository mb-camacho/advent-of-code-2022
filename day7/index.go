package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	// contents := getFileContents("input.txt")
	// firstMarker := getFirstMarker(contents)
	// fmt.Printf("First Marker: %v\n", firstMarker)
	// firstMarker = getStartOfMessageMarker(contents)
	// fmt.Printf("Start of Message Marker: %v\n", firstMarker)
}

func getFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func sumOfTotalSizes(contents string, mode int) int {
	commands := strings.Split(contents, "\n")
	var currentDir *dir
	// var parentDir *dir
	dirs := make(map[string]*dir)
	main := dir{name: "/"}
	dirs[main.name] = &main
	// currentDir = &main
	for _, command := range commands {
		if strings.HasPrefix(command, "$ cd ") {
			rawDir := strings.ReplaceAll(command, "$ cd ", "")
			if rawDir == ".." {
				currentDir = currentDir.parent
				continue
			}
			// parentDir = currentDir
			var parentName string
			if currentDir == nil {
				parentName = ""
			} else {
				parentName = currentDir.name
			}
			dirName := formatDirName(parentName + strings.ReplaceAll(command, "$ cd ", ""))
			currentDir = dirs[dirName]
			if currentDir == nil {
				panic("dir not found. name:" + dirName)
			}
			continue
		}
		if strings.HasPrefix(command, "$ ls") {
			continue
		}
		if strings.HasPrefix(command, "$") {
			panic("unknown command:" + command)
		}
		array := strings.Split(command, " ")
		arrayLen := len(array)
		if arrayLen != 2 {
			panic("invalid file/dir. expected:2 actual:" + strconv.Itoa(arrayLen) + " command:" + command)
		}
		if array[0] == "dir" {
			child := dir{name: formatDirName(currentDir.name + array[1])}
			if _, ok := dirs[child.name]; ok {
				panic("dir already exists... name:" + child.name)
			}
			dirs[child.name] = &child
			addDir(currentDir, &child)
			continue
		}
		size, err := strconv.Atoi(array[0])
		if err != nil {
			panic("failed to convert to int. " + array[0])
		}
		file := file{name: array[1], size: size}
		addFile(currentDir, &file)
		continue
	}
	updateSize(&main)
	if mode == 1 {
		return main.size
	}

	totalSize := 0
	maxSize := 70000000
	requiredFreeSpace := 30000000
	freeSpace := maxSize - main.size
	spaceToFee := requiredFreeSpace - freeSpace
	dirSizeToDelete := 0
	for _, v := range dirs {
		if v.size >= spaceToFee && (dirSizeToDelete == 0 || v.size < dirSizeToDelete) {
			dirSizeToDelete = v.size
		}
		if v.size > 100000 {
			continue
		}
		totalSize += v.size
	}
	if mode == 2 {
		return totalSize
	}
	if mode == 3 {
		return dirSizeToDelete
	}
	return 0
}

func addDir(parent *dir, child *dir) {
	child.parent = parent
	parent.dirs = append(parent.dirs, child)
	parent.size += child.size
}

func addFile(dir *dir, file *file) {
	dir.files = append(dir.files, file)
	dir.size += file.size
}

func updateSize(dir *dir) {
	dir.size = 0
	for _, file := range dir.files {
		dir.size += file.size
	}
	for _, child := range dir.dirs {
		updateSize(child)
		dir.size += child.size
	}
}

func formatDirName(dir string) string {
	if strings.HasSuffix(dir, "/") {
		return dir
	}
	return dir + "/"
}

func getCombinations(combinations []combination, data map[int]int, length int) []combination {
	if len(combinations) == 0 {
		for i, value := range data {
			combination := combination{}
			combination.indexes = make(map[int]int)
			combination.indexes[i] = value
			combinations = append(combinations, combination)
		}
	}
	if length == 1 {
		return combinations
	}
	var newCombinations []combination
	for _, combinationO := range combinations {
		for i, v := range data {
			if _, ok := combinationO.indexes[i]; ok {
				continue
			}
			newCombination := combination{}
			newCombination.indexes = make(map[int]int)
			total := 0
			for j, cv := range combinationO.indexes {
				newCombination.indexes[j] = cv
				total += cv
			}
			if total+v > 100000 {
				continue
			}
			newCombination.indexes[i] = v
			newCombinations = append(newCombinations, newCombination)
		}
	}
	if len(newCombinations) == 0 {
		return newCombinations
	}
	return getCombinations(newCombinations, data, length-1)
}

type combination struct {
	indexes map[int]int
}

type dir struct {
	name   string
	size   int
	parent *dir
	dirs   []*dir
	files  []*file
}

type file struct {
	name string
	size int
}
