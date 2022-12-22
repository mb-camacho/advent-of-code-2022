package main

import "testing"

func TestCountPositionsTailVisited(t *testing.T) {
	var actual int
	var expected int
	actual = countPositionsTailVisited(getFileContents("input-test.txt"))
	expected = 13
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = countPositionsTailVisited(getFileContents("input.txt"))
	expected = 6503
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}

func TestCountPositionsTail9Visited(t *testing.T) {
	var actual int
	var expected int
	actual = countPositionsTail9Visited(getFileContents("input-test.txt"))
	expected = 1
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = countPositionsTail9Visited(getFileContents("input-test2.txt"))
	expected = 36
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = countPositionsTail9Visited(getFileContents("input.txt"))
	expected = 2724
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}
