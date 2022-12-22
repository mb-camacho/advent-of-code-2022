package main

import "testing"

func TestCountVisibleTrees(t *testing.T) {
	var actual int
	var expected int
	actual = countVisibleTrees(getFileContents("input-test.txt"))
	expected = 21
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = countVisibleTrees(getFileContents("input.txt"))
	expected = 1647
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}

func TestGetHighestScenicScore(t *testing.T) {
	var actual int
	var expected int
	actual = getHighestScenicScore(getFileContents("input-test.txt"))
	expected = 8
	if actual != expected {
		t.Fatalf("TEST highest scenic score :: expected:%v. actual:%v", expected, actual)
	}

	actual = getHighestScenicScore(getFileContents("input.txt"))
	expected = 392080
	if actual != expected {
		t.Fatalf("LIVE highest scenic score :: expected:%v. actual:%v", expected, actual)
	}
}
