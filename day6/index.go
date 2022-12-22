package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents := getFileContents("input.txt")
	firstMarker := getFirstMarker(contents)
	fmt.Printf("First Marker: %v\n", firstMarker)
	firstMarker = getStartOfMessageMarker(contents)
	fmt.Printf("Start of Message Marker: %v\n", firstMarker)
}

func getFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getFirstMarker(signal string) int {
	length := 4
	for i := 0; i <= len(signal)-length; i++ {
		fourChars := signal[i : length+i]
		count := 1
		for j := 0; j < len(fourChars); j++ {
			char := fourChars[j : j+1]
			count = strings.Count(fourChars, char)
			if count > 1 {
				break
			}
		}
		if count == 1 {
			return i + 4
		}
	}
	panic("did not detec four chars. signal:" + signal)
}

func getStartOfMessageMarker(signal string) int {
	length := 14
	for i := 0; i <= len(signal)-length; i++ {
		fourChars := signal[i : length+i]
		count := 1
		for j := 0; j < len(fourChars); j++ {
			char := fourChars[j : j+1]
			count = strings.Count(fourChars, char)
			if count > 1 {
				break
			}
		}
		if count == 1 {
			return i + length
		}
	}
	panic("did not detec four chars. signal:" + signal)
}
