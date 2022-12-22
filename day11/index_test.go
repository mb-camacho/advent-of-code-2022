package main

import "testing"

func TestLevelOfMonkeyBusiness(t *testing.T) {
	var actual int
	var expected int
	actual = levelOfMonkeyBusiness(getFileContents("input-test.txt"), 20)
	expected = 10605
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = levelOfMonkeyBusiness(getFileContents("input.txt"), 20)
	expected = 117640
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}

func TestLevelOfMonkeyBusinessNotWorried(t *testing.T) {
	var actual int
	var expected int
	actual = levelOfMonkeyBusiness(getFileContents("input-test.txt"), 10000)
	expected = 2713310158
	if actual != expected {
		t.Fatalf("TEST visible trees :: expected:%v. actual:%v", expected, actual)
	}

	actual = levelOfMonkeyBusiness(getFileContents("input.txt"), 10000)
	expected = 30616425600
	if actual != expected {
		t.Fatalf("LIVE visible trees :: expected:%v. actual:%v", expected, actual)
	}
}
