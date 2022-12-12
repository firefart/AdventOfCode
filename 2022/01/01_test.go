package main

import (
	"testing"
)

var input = []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

func TestParseElves(t *testing.T) {
	elves, err := parseElves(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(elves) != 5 {
		t.Fatalf("expected 5 elves, got %d", len(elves))
	}
}

func TestBiggest(t *testing.T) {
	elves, err := parseElves(input)
	if err != nil {
		t.Fatal(err)
	}
	want := 24000
	got := biggest(elves)
	if got != want {
		t.Fatalf("biggest() got %d, want %d", got, want)
	}
}

func TestTop3Sum(t *testing.T) {
	elves, err := parseElves(input)
	if err != nil {
		t.Fatal(err)
	}
	want := 45000
	got := top3Sum(elves)
	if got != want {
		t.Fatalf("top3Sum() got %d, want %d", got, want)
	}
}
