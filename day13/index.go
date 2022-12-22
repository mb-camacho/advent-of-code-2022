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

func sumOfIndices(data string) int {
	pairs := getPairs(data)
	indices := make([]int, 0)
	for _, pair := range pairs {
		if len(pair) != 2 {
			panic("expects len to be 2")
		}
		index := getIndex(pair[0], pair[1])
		if index == 0 {
			continue
		}
		indices = append(indices, index)
	}
	sum := 0
	for _, i := range indices {
		sum += i
	}
	return sum
}

func getIndex(object1 *node, object2 *node) int {
	var list1, list2 []*node
	if object1.value > 0 {
		list1 = append(list1, object1)
	} else {
		list1 = object1.indices
	}
	if object2.value > 0 {
		list2 = append(list2, object2)
	} else {
		list2 = object2.indices
	}
	for i, obj2 := range list2 {
		if i > len(list1)-1 {
			return 1
		}
		obj1 := list1[i]
		if obj1.value < obj2.value {
			return 1
		}
		if obj1.value > obj2.value {
			return 0
		}
	}
	if len(list1) > len(list2) {
		return 0
	}
	return 1
}

func getPairs(data string) [][]*node {
	pairsArray := strings.Split(data, "\n\n")
	var pairObjects [][]*node
	for _, pairString := range pairsArray {
		pairArray := strings.Split(pairString, "\n")
		if len(pairsArray) != 2 {
			panic("expected pair. string:" + pairString)
		}
		var objects []*node
		for _, v := range pairArray {
			object := node{indices: make([]*node, 0)}
			buildNodes(v, 0, &object)
			objects = append(objects, object.indices[0])
		}
		pairObjects = append(pairObjects, objects)
	}
	return pairObjects
}

func buildNodes(data string, index int, nodeObj *node) int {
	value := ""
	for i := index; i < len(data); i++ {
		v := rune(data[i])
		if v == '[' {
			nodeObjChild := node{indices: make([]*node, 0)}
			nodeObj.indices = append(nodeObj.indices, &nodeObjChild)
			i = buildNodes(data, i+1, &nodeObjChild)
			value = ""
			continue
		}
		if v == ',' || v == ']' {
			if value == "" {
				continue
			}
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			nodeObjChild := node{value: valueInt, indices: make([]*node, 0)}
			nodeObj.indices = append(nodeObj.indices, &nodeObjChild)
			value = ""
			if v == ']' {
				return i
			}
			continue
		}
		value += string(v)
	}
	return len(data)
}

type node struct {
	value   int
	indices []*node
}
