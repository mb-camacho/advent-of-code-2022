package main

import "testing"

func TestGetFirstMarker(t *testing.T) {
	firstMarker := getFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	expected := 7
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
	expected = 5
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getFirstMarker("nppdvjthqldpwncqszvftbrmjlhg")
	expected = 6
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	expected = 10
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getFirstMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	expected = 11
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}
}

func TestGetStartOfMessageMarker(t *testing.T) {
	firstMarker := getStartOfMessageMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	expected := 19
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getStartOfMessageMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
	expected = 23
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getStartOfMessageMarker("nppdvjthqldpwncqszvftbrmjlhg")
	expected = 23
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getStartOfMessageMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	expected = 29
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}

	firstMarker = getStartOfMessageMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	expected = 26
	if firstMarker != expected {
		t.Fatalf("expected:%v. actual:%v", expected, firstMarker)
	}
}
