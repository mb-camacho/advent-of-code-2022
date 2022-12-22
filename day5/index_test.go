package main

import "testing"

func TestGetCratesOnTopOfEachStack(t *testing.T) {
	crates := getCratesOnTopOfEachStack("input-test.txt")
	expected := "CMZ"
	if crates != expected {
		t.Fatalf("expected:%v. actual:%v", expected, crates)
	}
}

func TestGetCratesOnTopOfEachStackV2(t *testing.T) {
	crates := getCratesOnTopOfEachStackV2("input-test.txt")
	expected := "MCD"
	if crates != expected {
		t.Fatalf("expected:%v. actual:%v", expected, crates)
	}
}
