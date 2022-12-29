package main

import (
	"fmt"
	"testing"
)

var in1 = "" +
	"    [D]    \n" +
	"[N] [C]    \n" +
	"[Z] [M] [P]\n" +
	" 1   2   3 \n"

var in2 = "" +
	"                        [Z] [W] [Z]\n" +
	"        [D] [M]         [L] [P] [G]\n" +
	"    [S] [N] [R]         [S] [F] [N]\n" +
	"    [N] [J] [W]     [J] [F] [D] [F]\n" +
	"[N] [H] [G] [J]     [H] [Q] [H] [P]\n" +
	"[V] [J] [T] [F] [H] [Z] [R] [L] [M]\n" +
	"[C] [M] [C] [D] [F] [T] [P] [S] [S]\n" +
	"[S] [Z] [M] [T] [P] [C] [D] [C] [D]\n" +
	" 1   2   3   4   5   6   7   8   9 \n"

var in3 = "" +
	"    [D]    \n" +
	"[N] [C]    \n" +
	"[Z] [M] [P]\n" +
	" 1   2   3 \n" +
	"\n" +
	"move 1 from 2 to 1\n" +
	"move 3 from 1 to 3\n" +
	"move 2 from 2 to 1\n" +
	"move 1 from 1 to 2\n"

func TestLogic(t *testing.T) {
	if err := logic([]byte(in3)); err != nil {
		t.Fatal(err)
	}
}

func TestParseCrateString(t *testing.T) {
	cr1 := parseCrates(in1)
	if len(cr1) != 3 {
		t.Errorf("cr1 len == %d, want %d", len(cr1), 3)
	}
	if len(cr1[0]) != 3 {
		t.Errorf("cr1 len2 == %d, want %d", len(cr1[0]), 3)
	}
	cr2 := parseCrates(in2)
	if len(cr2) != 8 {
		t.Errorf("cr2 len == %d, want %d", len(cr2), 8)
	}
	if len(cr2[0]) != 9 {
		t.Errorf("cr2 len2 == %d, want %d", len(cr2[0]), 9)
	}
}

func TestEmptyCrate(t *testing.T) {
	cr1 := parseCrates(in1)
	cr2 := parseCrates(in2)

	tests := []struct {
		want   int
		column int
		crates [][]rune
	}{
		{
			want:   0,
			crates: cr1,
			column: 0,
		},
		{
			want:   -1,
			crates: cr1,
			column: 1,
		},
		{
			want:   1,
			crates: cr1,
			column: 2,
		},

		{
			want:   3,
			crates: cr2,
			column: 0,
		},
		{
			want:   1,
			crates: cr2,
			column: 1,
		},
		{
			want:   0,
			crates: cr2,
			column: 2,
		},
		{
			want:   0,
			crates: cr2,
			column: 3,
		},
		{
			want:   4,
			crates: cr2,
			column: 4,
		},
		{
			want:   2,
			crates: cr2,
			column: 5,
		},
		{
			want:   -1,
			crates: cr2,
			column: 6,
		},
		{
			want:   -1,
			crates: cr2,
			column: 7,
		},
		{
			want:   -1,
			crates: cr2,
			column: 8,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.column), func(t *testing.T) {
			got := emptyCrate(tc.column, tc.crates)
			want := tc.want
			if got != want {
				t.Errorf("firstEmptyIndexInColumn() got %d, want %d", got, want)
			}
		})
	}
}

func TestTopCrate(t *testing.T) {
	cr1 := parseCrates(in1)
	cr2 := parseCrates(in2)

	tests := []struct {
		wantIndex int
		wantRune  rune
		column    int
		crates    [][]rune
	}{
		{
			wantRune:  'N',
			wantIndex: 1,
			crates:    cr1,
			column:    0,
		},
		{
			wantRune:  'D',
			wantIndex: 0,
			crates:    cr1,
			column:    1,
		},
		{
			wantRune:  'P',
			wantIndex: 2,
			crates:    cr1,
			column:    2,
		},
		{
			wantRune:  'N',
			wantIndex: 4,
			crates:    cr2,
			column:    0,
		},
		{
			wantRune:  'S',
			wantIndex: 2,
			crates:    cr2,
			column:    1,
		},
		{
			wantRune:  'D',
			wantIndex: 1,
			crates:    cr2,
			column:    2,
		},
		{
			wantRune:  'M',
			wantIndex: 1,
			crates:    cr2,
			column:    3,
		},
		{
			wantRune:  'H',
			wantIndex: 5,
			crates:    cr2,
			column:    4,
		},
		{
			wantRune:  'J',
			wantIndex: 3,
			crates:    cr2,
			column:    5,
		},
		{
			wantRune:  'Z',
			wantIndex: 0,
			crates:    cr2,
			column:    6,
		},
		{
			wantRune:  'W',
			wantIndex: 0,
			crates:    cr2,
			column:    7,
		},
		{
			wantRune:  'Z',
			wantIndex: 0,
			crates:    cr2,
			column:    8,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.column), func(t *testing.T) {
			got := topCrate(tc.column, tc.crates)
			if got.rune != tc.wantRune {
				t.Errorf("topCrate() got rune %c, want %c", got.rune, tc.wantRune)
			}
			if rune(got.index) != rune(tc.wantIndex) {
				t.Errorf("topCrate() got index %d, want %d", got.index, tc.wantIndex)
			}
		})
	}
}
