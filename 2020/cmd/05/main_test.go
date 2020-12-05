package main

import "testing"

func TestGetSeatRowAndColumn(t *testing.T) {
	tests := []struct {
		input  string
		row    int
		column int
	}{
		{input: "BFFFBBFRRR", row: 70, column: 7},
		{input: "FFFBBBFRRR", row: 14, column: 7},
		{input: "BBFFBBFRLL", row: 102, column: 4},
	}

	for _, tc := range tests {
		row, column := getSeatRowAndColumn(tc.input)
		if row != tc.row {
			t.Errorf("row wrong. Expected %d got %d", tc.row, row)
		}

		if column != tc.column {
			t.Errorf("column wrong. Expected %d got %d", tc.column, column)
		}
	}
}

func TestGetSeatID(t *testing.T) {
	tests := []struct {
		row    int
		column int
		id     int
	}{
		{row: 70, column: 7, id: 567},
		{row: 14, column: 7, id: 119},
		{row: 102, column: 4, id: 820},
	}

	for _, tc := range tests {
		id := getSeatID(tc.row, tc.column)
		if id != tc.id {
			t.Errorf("column seatid. Expected %d got %d", tc.id, id)
		}
	}
}

func TestGetFirstMissingNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: -1},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, expected: -1},
		{input: []int{2, 3, 4, 5, 6, 7, 8, 9, 0}, expected: 1},
		{input: []int{4, 2, 3, 1, 5, 6, 7, 9, 0}, expected: 8},
	}

	for _, tc := range tests {
		missing := getFirstMissingNumber(tc.input)
		if missing != tc.expected {
			t.Errorf("missing number wrong. Expected %d got %d", tc.expected, missing)
		}
	}
}
