package main

import "testing"

func TestSumOfSignalStrengths(t *testing.T) {
	var actual int
	var expected int
	actual = sumOfSignalStrengths(getFileContents("input-test.txt"))
	expected = 13140
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = sumOfSignalStrengths(getFileContents("input.txt"))
	expected = 13760
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}
