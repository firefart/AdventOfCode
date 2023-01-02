package main

import (
	"fmt"
	"strings"
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
			crates: cloneMatrix(cr1),
			column: 0,
		},
		{
			want:   -1,
			crates: cloneMatrix(cr1),
			column: 1,
		},
		{
			want:   1,
			crates: cloneMatrix(cr1),
			column: 2,
		},

		{
			want:   3,
			crates: cloneMatrix(cr2),
			column: 0,
		},
		{
			want:   1,
			crates: cloneMatrix(cr2),
			column: 1,
		},
		{
			want:   0,
			crates: cloneMatrix(cr2),
			column: 2,
		},
		{
			want:   0,
			crates: cloneMatrix(cr2),
			column: 3,
		},
		{
			want:   4,
			crates: cloneMatrix(cr2),
			column: 4,
		},
		{
			want:   2,
			crates: cloneMatrix(cr2),
			column: 5,
		},
		{
			want:   -1,
			crates: cloneMatrix(cr2),
			column: 6,
		},
		{
			want:   -1,
			crates: cloneMatrix(cr2),
			column: 7,
		},
		{
			want:   -1,
			crates: cloneMatrix(cr2),
			column: 8,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.column), func(t *testing.T) {
			crates := matrix{
				content: tc.crates,
			}
			got := crates.emptyCrate(tc.column)
			want := tc.want
			if got != want {
				t.Errorf("firstEmptyIndexInColumn() got %d, want %d", got, want)
			}
		})
	}
}

func TestTopCrateInColumn(t *testing.T) {
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
			crates:    cloneMatrix(cr1),
			column:    0,
		},
		{
			wantRune:  'D',
			wantIndex: 0,
			crates:    cloneMatrix(cr1),
			column:    1,
		},
		{
			wantRune:  'P',
			wantIndex: 2,
			crates:    cloneMatrix(cr1),
			column:    2,
		},
		{
			wantRune:  'N',
			wantIndex: 4,
			crates:    cloneMatrix(cr2),
			column:    0,
		},
		{
			wantRune:  'S',
			wantIndex: 2,
			crates:    cloneMatrix(cr2),
			column:    1,
		},
		{
			wantRune:  'D',
			wantIndex: 1,
			crates:    cloneMatrix(cr2),
			column:    2,
		},
		{
			wantRune:  'M',
			wantIndex: 1,
			crates:    cloneMatrix(cr2),
			column:    3,
		},
		{
			wantRune:  'H',
			wantIndex: 5,
			crates:    cloneMatrix(cr2),
			column:    4,
		},
		{
			wantRune:  'J',
			wantIndex: 3,
			crates:    cloneMatrix(cr2),
			column:    5,
		},
		{
			wantRune:  'Z',
			wantIndex: 0,
			crates:    cloneMatrix(cr2),
			column:    6,
		},
		{
			wantRune:  'W',
			wantIndex: 0,
			crates:    cloneMatrix(cr2),
			column:    7,
		},
		{
			wantRune:  'Z',
			wantIndex: 0,
			crates:    cloneMatrix(cr2),
			column:    8,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.column), func(t *testing.T) {
			crates := matrix{
				content: tc.crates,
			}
			got := crates.topCrateInColumn(tc.column)
			if got.rune != tc.wantRune {
				t.Errorf("topCrateInColumn() got rune %c, want %c", got.rune, tc.wantRune)
			}
			if rune(got.index) != rune(tc.wantIndex) {
				t.Errorf("topCrateInColumn() got index %d, want %d", got.index, tc.wantIndex)
			}
		})
	}
}

func TestTopCratesInColumn(t *testing.T) {
	cr1 := parseCrates(in1)
	// cr2 := parseCrates(in2)

	tests := []struct {
		column int
		amount int
		crates [][]rune
		want   string
	}{
		{
			crates: cloneMatrix(cr1),
			column: 0,
			amount: 2,
			want:   "NZ",
		},
		{
			crates: cloneMatrix(cr1),
			column: 1,
			amount: 3,
			want:   "DCM",
		},
		{
			crates: cloneMatrix(cr1),
			column: 1,
			amount: 2,
			want:   "DC",
		},
		{
			crates: cloneMatrix(cr1),
			column: 1,
			amount: 1,
			want:   "D",
		},
		{
			crates: cloneMatrix(cr1),
			column: 1,
			amount: 0,
			want:   "",
		},
		{
			crates: cloneMatrix(cr1),
			column: 1,
			amount: 1000,
			want:   "DCM",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d from %d - %s", tc.amount, tc.column, tc.want), func(t *testing.T) {
			crates := matrix{
				content: tc.crates,
			}
			got := crates.topCratesInColumn(tc.column, tc.amount)
			sb := strings.Builder{}
			for _, r := range got {
				sb.WriteRune(r.rune)
			}
			gotString := sb.String()
			if gotString != tc.want {
				t.Errorf("topCratesInColumn() got %s, want %s", gotString, tc.want)
			}
		})
	}
}

func TestMoveStackToColumn(t *testing.T) {
	cr1 := parseCrates(in1)
	// cr2 := parseCrates(in2)

	tests := []struct {
		crateStack    string
		toColumn      int
		crates        [][]rune
		want          string
		wantRowsAdded int
	}{
		{
			crates:        cloneMatrix(cr1),
			toColumn:      2,
			crateStack:    "AB",
			want:          "ABP",
			wantRowsAdded: 0,
		},
		{
			crates:        cloneMatrix(cr1),
			toColumn:      2,
			crateStack:    "ABCDEFG",
			want:          "ABCDEFGP",
			wantRowsAdded: 5,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s to %d - %s", tc.crateStack, tc.toColumn, tc.want), func(t *testing.T) {
			crates := matrix{
				content: tc.crates,
			}
			crateStack := make([]crate, len(tc.crateStack))
			for i, s := range tc.crateStack {
				crateStack[i] = crate{
					rune: rune(s),
				}
			}
			rowsAdded := crates.moveStackToColumn(crateStack, tc.toColumn)
			if rowsAdded != tc.wantRowsAdded {
				t.Errorf("rowsAdded: got %d, want %d", rowsAdded, tc.wantRowsAdded)
			}
			sb := strings.Builder{}
			for _, r := range crates.content {
				sb.WriteRune(r[tc.toColumn])
			}
			gotString := sb.String()
			if gotString != tc.want {
				t.Errorf("moveStackToColumn() got %q, want %q", gotString, tc.want)
			}
		})
	}
}
