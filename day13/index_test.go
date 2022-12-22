package main

import (
	"fmt"
	"testing"
)

func TestBuildNodes(t *testing.T) {
	var nodeObj node
	var list []*node

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,2,3]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 3 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:3. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,[11,22],3]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 3 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:3. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[[11,22],3,4]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 3 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:3. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 0 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:0. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[[[]]]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 1 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:1. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,2,3]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 3 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:3. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,[11,22],3]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 3 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:3. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[[11,22],3,4]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 3 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:3. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 0 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:0. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[[[]]]", 0, &nodeObj)
	list = nodeObj.indices[0].indices
	if len(list) != 1 {
		fmt.Printf("%+v\n", nodeObj.indices[0])
		t.Fatalf("TEST default :: expected:1. actual:%v", len(list))
	}
}

func TestGetIndex(t *testing.T) {
	var nodeObj node
	var list []*node

	var index int
	var expected int
	var list1 *node
	var list2 *node

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,2,3]", 0, &nodeObj)
	list1 = nodeObj.indices[0]
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,2,4]", 0, &nodeObj)
	list2 = nodeObj.indices[0]
	index = getIndex(list1, list2)
	if index != 1 {
		fmt.Printf("%+v\n", list1)
		fmt.Printf("%+v\n", list2)
		t.Fatalf("TEST default :: expected:1. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,2,3]", 0, &nodeObj)
	list1 = nodeObj.indices[0]
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[1,2,4]", 0, &nodeObj)
	list2 = nodeObj.indices[0]
	index = getIndex(list1, list2)
	if index != 0 {
		fmt.Printf("%+v\n", list1)
		fmt.Printf("%+v\n", list2)
		t.Fatalf("TEST default :: expected:0. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,1]", 0, &nodeObj)
	list1 = nodeObj.indices[0]
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,1]", 0, &nodeObj)
	list2 = nodeObj.indices[0]
	index = getIndex(list1, list2)
	if index != 1 {
		fmt.Printf("%+v\n", list1)
		fmt.Printf("%+v\n", list2)
		t.Fatalf("TEST default :: expected:1. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,1]", 0, &nodeObj)
	list1 = nodeObj.indices[0]
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,1,2]", 0, &nodeObj)
	list2 = nodeObj.indices[0]
	index = getIndex(list1, list2)
	if index != 1 {
		fmt.Printf("%+v\n", list1)
		fmt.Printf("%+v\n", list2)
		t.Fatalf("TEST default :: expected:1. actual:%v", len(list))
	}

	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,1,2]", 0, &nodeObj)
	list1 = nodeObj.indices[0]
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[2,1]", 0, &nodeObj)
	list2 = nodeObj.indices[0]
	index = getIndex(list1, list2)
	if index != 0 {
		fmt.Printf("%+v\n", list1)
		fmt.Printf("%+v\n", list2)
		t.Fatalf("TEST default :: expected:0. actual:%v", len(list))
	}

	expected = 2
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[[1],[2,3,4]]", 0, &nodeObj)
	list1 = nodeObj.indices[0]
	nodeObj = node{indices: make([]*node, 0)}
	buildNodes("[[1],[4]]", 0, &nodeObj)
	list2 = nodeObj.indices[0]
	index = getIndex(list1, list2)
	if index != expected {
		fmt.Printf("%+v\n", list1)
		fmt.Printf("%+v\n", list2)
		t.Fatalf("TEST default :: expected:%v. actual:%v", expected, len(list))
	}
}

// func TestSumOfIndices(t *testing.T) {
// 	var actual int
// 	var expected int
// 	actual = sumOfIndices(getFileContents("input-test.txt"))
// 	expected = 13
// 	if actual != expected {
// 		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
// 	}

// actual = sumOfIndices(getFileContents("input.txt"))
// expected = 534
// if actual != expected {
// 	t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
// }
// }
