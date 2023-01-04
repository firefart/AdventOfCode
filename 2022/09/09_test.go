package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var input = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestParseMoves(t *testing.T) {
	want := []Move{
		{
			Dir:   DirectionRight,
			Count: 4,
		},
		{
			Dir:   DirectionUp,
			Count: 4,
		},
		{
			Dir:   DirectionLeft,
			Count: 3,
		},
		{
			Dir:   DirectionDown,
			Count: 1,
		},
		{
			Dir:   DirectionRight,
			Count: 4,
		},
		{
			Dir:   DirectionDown,
			Count: 1,
		},
		{
			Dir:   DirectionLeft,
			Count: 5,
		},
		{
			Dir:   DirectionRight,
			Count: 2,
		},
	}
	got := parseMoves(input)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseMoves() mismatch (-want +got):\n%s", diff)
	}
}

func Test(t *testing.T) {
	p := newPlayfield(3, 3)
	fmt.Println(p)
	p.addColumnToLeft()
	fmt.Println("#####")
	fmt.Println(p)
}

// func TestTopCratesInColumn(t *testing.T) {
// 	tests := []struct {
// 		column int
// 		amount int
// 		crates [][]rune
// 		want   string
// 	}{}

// 	for _, tc := range tests {
// 		t.Run(fmt.Sprintf("%d from %d - %s", tc.amount, tc.column, tc.want), func(t *testing.T) {

// 		})
// 	}
// }
