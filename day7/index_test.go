package main

import "testing"

func TestSumOfTotalSizesTest(t *testing.T) {
	var sum int
	var expected int
	sum = sumOfTotalSizes(getFileContents("input-test.txt"), 1)
	expected = 48381165
	if sum != expected {
		t.Fatalf("TEST :: expected:%v. actual:%v", expected, sum)
	}

	sum = sumOfTotalSizes(getFileContents("input-test.txt"), 2)
	expected = 95437
	if sum != expected {
		t.Fatalf("TEST :: expected:%v. actual:%v", expected, sum)
	}

	sum = sumOfTotalSizes(getFileContents("input-test.txt"), 3)
	expected = 24933642
	if sum != expected {
		t.Fatalf("TEST :: expected:%v. actual:%v", expected, sum)
	}
}

func TestSumOfTotalSizesLive(t *testing.T) {
	var sum int
	var expected int
	sum = sumOfTotalSizes(getFileContents("input.txt"), 1)
	expected = 40913445
	if sum != expected {
		t.Fatalf("LIVE :: expected:%v. actual:%v", expected, sum)
	}

	sum = sumOfTotalSizes(getFileContents("input.txt"), 2)
	expected = 1443806
	if sum != expected {
		t.Fatalf("LIVE :: expected:%v. actual:%v", expected, sum)
	}

	sum = sumOfTotalSizes(getFileContents("input.txt"), 3)
	expected = 942298
	if sum != expected {
		t.Fatalf("LIVE :: expected:%v. actual:%v", expected, sum)
	}
}
