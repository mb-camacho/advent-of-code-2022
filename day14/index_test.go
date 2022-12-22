package main

import (
	"testing"
)

func TestCountUnitsOfSand(t *testing.T) {
	var actual int
	var expected int

	expected = 24
	actual = countUnitsOfSand(getFileContents("input-test.txt"))
	if actual != expected {
		t.Fatalf("TEST :: expected:%v. actual:%v", expected, actual)
	}

	expected = 832
	actual = countUnitsOfSand(getFileContents("input.txt"))
	if actual != expected {
		t.Fatalf("LIVE :: expected:%v. actual:%v", expected, actual)
	}
}

func TestCountUnitsOfSandWithFloor(t *testing.T) {
	var actual int
	var expected int

	expected = 93
	actual = countUnitsOfSandWithFloor(getFileContents("input-test.txt"))
	if actual != expected {
		t.Fatalf("TEST :: expected:%v. actual:%v", expected, actual)
	}

	expected = 27601
	actual = countUnitsOfSandWithFloor(getFileContents("input.txt"))
	if actual != expected {
		t.Fatalf("LIVE :: expected:%v. actual:%v", expected, actual)
	}
}
