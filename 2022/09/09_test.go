package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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

func TestAddRowAndColumn(t *testing.T) {
	m := newPlayfield(3, 3)
	assert := assert.New(t)
	assert.Equal(3, len(m.Content))
	assert.Equal(3, len(m.Content[0]))
	m.addColumnToLeft()
	assert.Equal(3, len(m.Content))
	assert.Equal(4, len(m.Content[0]))
	m.addColumnToLeft()
	assert.Equal(3, len(m.Content))
	assert.Equal(5, len(m.Content[0]))
	m.addRowToBottom()
	assert.Equal(4, len(m.Content))
	assert.Equal(5, len(m.Content[0]))
	m.addColumnToRight()
	assert.Equal(4, len(m.Content))
	assert.Equal(6, len(m.Content[0]))
	m.addRowToBottom()
	assert.Equal(5, len(m.Content))
	assert.Equal(6, len(m.Content[0]))
	m.addRowToTop()
	assert.Equal(6, len(m.Content))
	assert.Equal(6, len(m.Content[0]))
}

func TestGetTailVisitNumber(t *testing.T) {
	field := newPlayfield(3, 3)
	moves := parseMoves(input)
	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			field.moveHead(m.Dir)
			field.moveTailToHead()
		}
	}
	field.Finished = true
	got := field.getTailVisitNumber()
	want := 13
	if got != want {
		t.Fatalf("getTailVisitNumber() got %d, want %d", got, want)
	}
}
