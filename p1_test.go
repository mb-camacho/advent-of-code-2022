package main

import "testing"

func TestGetMax(t *testing.T) {
	max := getMax("input-test.txt")
	expected := 24000
	if max != expected {
		t.Fatalf("expected:%d. actual:%d", expected, max)
	}
}

func TestGetSumTop3(t *testing.T) {
	max := getSumTop3("input-test.txt")
	expected := 45000
	if max != expected {
		t.Fatalf("expected:%q. actual:%q", expected, max)
	}
}

func TestGetPlayScore(t *testing.T) {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	score := getPlayScore("X")
	expected := 1
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected, score)
	}

	score2 := getPlayScore("Y")
	expected2 := 2
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected2, score2)
	}

	score3 := getPlayScore("Z")
	expected3 := 3
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected3, score3)
	}
}

func TestGetPlayResultScore(t *testing.T) {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors
	score := getPlayResultScore("X", "C")
	expected := 6
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected, score)
	}

	score2 := getPlayResultScore("Y", "B")
	expected2 := 3
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected2, score2)
	}

	score3 := getPlayResultScore("Z", "A")
	expected3 := 0
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected3, score3)
	}
}

func TestGetScore(t *testing.T) {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors
	score := getScore("input-day2-test.txt")
	expected := 15
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected, score)
	}
}

func TestGetScoreV2(t *testing.T) {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors
	score := getScoreV2("input-day2-test.txt")
	expected := 12
	if score != expected {
		t.Fatalf("expected:%d. actual:%d", expected, score)
	}
}

func TestGetSumOfPriorities(t *testing.T) {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors
	sum := getSumOfPriorities("input-day3-test.txt")
	expected := 157
	if sum != expected {
		t.Fatalf("expected:%d. actual:%d", expected, sum)
	}
}

func TestGetSumOfGroupPriorities(t *testing.T) {
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors
	sum := getSumOfGroupPriorities("input-day3-test.txt")
	expected := 70
	if sum != expected {
		t.Fatalf("expected:%d. actual:%d", expected, sum)
	}
}

func TestGetCountOfAssignmentPairFullyContainingTheOther(t *testing.T) {
	count := getCountOfAssignmentPairFullyContainingTheOther("input-day4-test.txt")
	expected := 2
	if count != expected {
		t.Fatalf("expected:%d. actual:%d", expected, count)
	}
}

func TestGetCountOfAssignmentPairOverlapping(t *testing.T) {
	count := getCountOfAssignmentPairOverlapping("input-day4-test.txt")
	expected := 4
	if count != expected {
		t.Fatalf("expected:%d. actual:%d", expected, count)
	}
}
