package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Row int
	Col int
}

type Field struct {
	Content       string
	VisitedByTail bool
}

type Playfield struct {
	Content  [][]Field
	Head     Position
	Tail     Position
	Finished bool
}

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionUp
	DirectionDown
)

func (d Direction) String() string {
	switch d {
	case DirectionDown:
		return "down"
	case DirectionLeft:
		return "left"
	case DirectionRight:
		return "right"
	case DirectionUp:
		return "up"
	default:
		panic(fmt.Sprintf("invalid direction %d", d))
	}
}

type Move struct {
	Dir   Direction
	Count int
}

func (m Move) String() string {
	plural := ""
	if m.Count > 1 {
		plural = "s"
	}
	return fmt.Sprintf("move %d step%s %s", m.Count, plural, m.Dir)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err := logic(content); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func logic(input []byte) error {
	moves := parseMoves(string(input))
	for _, m := range moves {
		fmt.Println(m)
	}
	return nil
}

func parseMoves(in string) []Move {
	tmp := strings.Split(in, "\n")
	moves := make([]Move, len(tmp))
	for i, l := range tmp {
		split := strings.Split(l, " ")
		count, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		moves[i] = Move{
			Dir:   convertDirection(split[0]),
			Count: count,
		}
	}

	return moves
}

func convertDirection(dir string) Direction {
	switch strings.TrimSpace(dir) {
	case "R":
		return DirectionRight
	case "L":
		return DirectionLeft
	case "U":
		return DirectionUp
	case "D":
		return DirectionDown
	default:
		panic(fmt.Sprintf("invalid direction %q", dir))
	}
}

func newPlayfield(rows, cols int) *Playfield {
	matrix := make([][]Field, rows)
	for i := range matrix {
		matrix[i] = make([]Field, cols)
	}
	return &Playfield{
		// starting bottom left
		Head: Position{
			Row: rows - 1,
			Col: 0,
		},
		// starting bottom left
		Tail: Position{
			Row: rows - 1,
			Col: 0,
		},
		Finished: false,
		Content:  matrix,
	}
}

func (m Playfield) String() string {
	var sb strings.Builder
	for _, row := range m.Content {
		for _, col := range row {
			content := col.Content
			if m.Finished && col.VisitedByTail {
				content = "#"
			}
			if content == "" {
				content = "."
			}
			sb.WriteString(fmt.Sprintf("%s", content))
		}
		sb.WriteString("\n")
	}
	return strings.TrimSpace(sb.String())
}

func (m *Playfield) addRowToBottom() {
	rowCount := len(m.Content)
	newField := make([][]Field, rowCount+1)
	newField[rowCount] = make([]Field, len(m.Content[rowCount-1]))
	for i, row := range m.Content {
		newField[i] = row
	}
	m.Content = newField
}

func (m *Playfield) addRowToTop() {

	// move head and tails one row down as the coordinates change
	m.Head.Row += 1
	m.Tail.Row += 1
}

func (m *Playfield) addColumnToRight() {
	rowCount := len(m.Content)
	newField := make([][]Field, rowCount)
	for i, row := range m.Content {
		newField[i] = make([]Field, len(row)+1)
		for j, col := range row {
			newField[i][j] = col
		}
		newField[i][len(row)] = Field{}
	}
	m.Content = newField
}

func (m *Playfield) addColumnToLeft() {

	// move head and tails one col right as the coordinates change
	m.Head.Col += 1
	m.Tail.Col += 1
}
