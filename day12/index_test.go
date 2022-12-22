package main

import (
	"fmt"
	"testing"
)

func TestDefault(t *testing.T) {
	// var actual int
	// var expected int
	testNode := node{}
	actual := testNode.bNode
	// expected := nil
	if actual != nil {
		t.Fatalf("TEST default :: expected:nil. actual:%v", actual)
	}

	str := "hello"
	fmt.Print("str[0]:" + str[4:] + "\n")
}

func TestGetFewestStep(t *testing.T) {
	var actual int
	var expected int
	actual = getFewestStep(getFileContents("input-test.txt"))
	expected = 31
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = getFewestStep(getFileContents("input.txt"))
	expected = 534
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}

func TestGetFewestStepMultipleStart(t *testing.T) {
	var actual int
	var expected int
	actual = getFewestStepMultipleStart(getFileContents("input-test.txt"))
	expected = 29
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = getFewestStepMultipleStart(getFileContents("input.txt"))
	expected = 525
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}
