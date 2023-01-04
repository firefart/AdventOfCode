package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
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
	Rows     int
	Columns  int
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

func printFieldToTerminal(field *Playfield, m *Move) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	if m != nil {
		tm.Println(m)
	}
	if field != nil {
		tm.Println(field)
	}
	tm.Flush()
}

func logic(input []byte) error {
	const sleepTime = 1 * time.Second
	// start with a 3x3 grid as we don't know the final size
	field := newPlayfield(3, 3)
	moves := parseMoves(string(input))

	// print initial field
	printFieldToTerminal(field, nil)
	time.Sleep(sleepTime)

	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			field.moveHead(m.Dir)
			printFieldToTerminal(field, &m)
			time.Sleep(sleepTime)
			field.moveTailToHead()
			printFieldToTerminal(field, &m)
			time.Sleep(sleepTime)
		}
	}
	field.Finished = true
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(field)
	tm.Printf("Part1: Tail visited %d tiles\n", field.getTailVisitNumber())
	tm.Flush()
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
		Rows:     rows,
		Columns:  cols,
		Finished: false,
		Content:  matrix,
	}
}

func (p Playfield) String() string {
	var sb strings.Builder
	for rowIndex, row := range p.Content {
		for colIndex, col := range row {
			content := col.Content
			if p.Finished && col.VisitedByTail {
				content = "#"
			}
			if content == "" {
				content = "."
			}

			// print tail (T), head (H) or overlapping (O) when not finished
			if !p.Finished {
				if p.Tail.Row == p.Head.Row && p.Tail.Col == p.Head.Col && p.Tail.Row == rowIndex && p.Tail.Col == colIndex {
					content = "O"
				} else if p.Tail.Row == rowIndex && p.Tail.Col == colIndex {
					content = "T"
				} else if p.Head.Row == rowIndex && p.Head.Col == colIndex {
					content = "H"
				}
			}

			sb.WriteString(fmt.Sprintf("%s", content))
		}
		sb.WriteString("\n")
	}
	return strings.TrimSpace(sb.String())
}

func (p *Playfield) addRowToBottom() {
	newField := make([][]Field, p.Rows+1)
	newField[p.Rows] = make([]Field, len(p.Content[p.Rows-1]))
	for i, row := range p.Content {
		newField[i] = row
	}
	p.Content = newField
	p.Rows += 1
}

func (p *Playfield) addRowToTop() {
	newField := make([][]Field, p.Rows+1)
	newField[0] = make([]Field, p.Columns)
	for i, row := range p.Content {
		newField[i+1] = row
	}
	p.Content = newField
	p.Rows += 1

	// move head and tails one row down as the coordinates change
	p.Head.Row += 1
	p.Tail.Row += 1
}

func (p *Playfield) addColumnToRight() {
	newField := make([][]Field, p.Rows)
	for i, row := range p.Content {
		newField[i] = make([]Field, len(row)+1)
		for j, col := range row {
			newField[i][j] = col
		}
		newField[i][len(row)] = Field{}
	}
	p.Content = newField
	p.Columns += 1
}

func (p *Playfield) addColumnToLeft() {
	newField := make([][]Field, p.Rows)
	for i, row := range p.Content {
		newField[i] = make([]Field, len(row)+1)
		newField[i][0] = Field{}
		for j, col := range row {
			newField[i][j+1] = col
		}
	}
	p.Content = newField
	p.Columns += 1

	// move head and tails one col right as the coordinates change
	p.Head.Col += 1
	p.Tail.Col += 1
}

func (p *Playfield) moveHead(dir Direction) {
	switch dir {
	case DirectionUp:
		// check if we need to add a row to the top
		if p.Head.Row == 0 {
			p.addRowToTop()
		}
		p.Head.Row -= 1
	case DirectionDown:
		// check if we need to add a row to the bottom
		if p.Head.Row == p.Rows-1 {
			p.addRowToBottom()
		}
		p.Head.Row += 1
	case DirectionLeft:
		// check if we need to add a column to the left
		if p.Head.Col == 0 {
			p.addColumnToLeft()
		}
		p.Head.Col -= 1
	case DirectionRight:
		// check if we need to add a column to the right
		if p.Head.Col == p.Columns-1 {
			p.addColumnToRight()
		}
		p.Head.Col += 1
	default:
		panic(fmt.Sprintf("direction %d not implemented", dir))
	}
}

func (p *Playfield) moveTailToHead() {

}

func (p *Playfield) getTailVisitNumber() int {
	if !p.Finished {
		panic("only call this method after finish!")
	}
	count := 0
	for _, row := range p.Content {
		for _, col := range row {
			if col.VisitedByTail {
				count++
			}
		}
	}
	return count
}
