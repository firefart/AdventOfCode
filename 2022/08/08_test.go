package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var input = `30373
25512
65332
33549
35390`

var parsed = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestParseTrees(t *testing.T) {
	trees, err := parseTrees([]byte(input))
	if err != nil {
		t.Fatalf("parseTrees() err=%v, want nil", err)
	}

	if diff := cmp.Diff(parsed, trees); diff != "" {
		t.Errorf("parseTrees() mismatch (-want +got):\n%s", diff)
	}
}

func TestTreeVisible(t *testing.T) {
	tests := []struct {
		row  int
		col  int
		want bool
	}{
		{
			// edge
			row:  0,
			col:  0,
			want: true,
		},
		{
			// edge
			row:  0,
			col:  1,
			want: true,
		},
		{
			// edge
			row:  0,
			col:  2,
			want: true,
		},
		{
			// edge
			row:  0,
			col:  3,
			want: true,
		},
		{
			// edge
			row:  0,
			col:  4,
			want: true,
		},
		{
			// edge
			row:  1,
			col:  0,
			want: true,
		},
		{
			// The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
			row:  1,
			col:  1,
			want: true,
		},
		{
			// The top-middle 5 is visible from the top and right.
			row:  1,
			col:  2,
			want: true,
		},
		{
			// The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
			row:  1,
			col:  3,
			want: false,
		},
		{
			// edge
			row:  1,
			col:  4,
			want: true,
		},
		{
			// edge
			row:  2,
			col:  0,
			want: true,
		},
		{
			// The left-middle 5 is visible, but only from the right.
			row:  2,
			col:  1,
			want: true,
		},
		{
			// The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
			row:  2,
			col:  2,
			want: false,
		},
		{
			// The right-middle 3 is visible from the right.
			row:  2,
			col:  3,
			want: true,
		},
		{
			// edge
			row:  2,
			col:  4,
			want: true,
		},
		{
			// edge
			row:  3,
			col:  0,
			want: true,
		},
		{
			// In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
			row:  3,
			col:  1,
			want: false,
		},
		{
			// In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
			row:  3,
			col:  2,
			want: true,
		},
		{
			// In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
			row:  3,
			col:  3,
			want: false,
		},
		{
			// edge
			row:  3,
			col:  4,
			want: true,
		},
		{
			// edge
			row:  4,
			col:  0,
			want: true,
		},
		{
			// edge
			row:  4,
			col:  1,
			want: true,
		},
		{
			// edge
			row:  4,
			col:  2,
			want: true,
		},
		{
			// edge
			row:  4,
			col:  3,
			want: true,
		},
		{
			// edge
			row:  4,
			col:  4,
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d/%d -> %t", tc.row, tc.col, tc.want), func(t *testing.T) {
			got := treeVisible(parsed, tc.row, tc.col)
			if got != tc.want {
				t.Errorf("treeVisible() got %t want %t", got, tc.want)
			}
		})
	}
}

func TestCountVisibleTrees(t *testing.T) {
	got := countVisibleTrees(parsed)
	want := 21
	if got != want {
		t.Fatalf("countVisibleTrees() got %d want %d", got, want)
	}
}

func TestScenicScore(t *testing.T) {
	tests := []struct {
		row  int
		col  int
		want int
	}{
		{
			row:  1,
			col:  2,
			want: 4,
		},
		{
			row:  3,
			col:  2,
			want: 8,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d/%d -> %d", tc.row, tc.col, tc.want), func(t *testing.T) {
			got := scenicScore(parsed, tc.row, tc.col)
			if got != tc.want {
				t.Errorf("scenicScore() got %d want %d", got, tc.want)
			}
		})
	}
}

func TestBestScenicScore(t *testing.T) {
	got := bestScenicScore(parsed)
	want := 8
	if got != want {
		t.Fatalf("bestScenicScore() got %d want %d", got, want)
	}
}
