package main

import (
	"flag"
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
	sleepDuration := flag.Duration("sleep", 1*time.Second, "how long to sleep")
	printToScreen := flag.Bool("print", true, "print to screen")
	flag.Parse()
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
	if err := logic(content, *sleepDuration, *printToScreen); err != nil {
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

func logic(input []byte, sleepDuration time.Duration, printToScreen bool) error {
	// start with a 3x3 grid as we don't know the final size
	field := newPlayfield(3, 3)
	moves := parseMoves(string(input))

	if printToScreen {
		// print initial field
		printFieldToTerminal(field, nil)
		time.Sleep(sleepDuration)
	}

	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			field.moveHead(m.Dir)
			if printToScreen {
				printFieldToTerminal(field, &m)
				time.Sleep(sleepDuration)
			}
			field.moveTailToHead()
			if printToScreen {
				printFieldToTerminal(field, &m)
				time.Sleep(sleepDuration)
			}
		}
	}
	field.Finished = true
	if printToScreen {
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Println(field)
	}
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
	// as we start bottom left this field is always visited by the tail
	matrix[rows-1][0].VisitedByTail = true
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
	// bail out on overlap
	if p.Tail.Row == p.Head.Row && p.Tail.Col == p.Head.Col {
		p.Content[p.Tail.Row][p.Tail.Col].VisitedByTail = true
		return
	}

	// bail out if directly below or above
	if p.Tail.Col == p.Head.Col && (p.Tail.Row == p.Head.Row-1 || p.Tail.Row == p.Head.Row+1) {
		return
	}

	// bail out if left or right
	if p.Tail.Row == p.Head.Row && (p.Tail.Col == p.Head.Col-1 || p.Tail.Col == p.Head.Col+1) {
		return
	}

	// check diagonals
	if (p.Tail.Row == p.Head.Row-1 || p.Tail.Row == p.Head.Row+1) && (p.Tail.Col == p.Head.Col-1 || p.Tail.Col == p.Head.Col+1) {
		return
	}

	// col pos: left
	// col neg: right
	// row pos: up
	// row neg: down
	rowDirection := p.Tail.Row - p.Head.Row
	colDirection := p.Tail.Col - p.Head.Col

	if rowDirection == 0 {
		// only move left or right in this case as we are already in the correct row
		if colDirection > 0 {
			// move tail to the left
			p.Tail.Col -= 1
		} else {
			// move tail to the right
			p.Tail.Col += 1
		}
	} else if colDirection == 0 {
		// only move up or down in this case as we are already in the correct row
		if rowDirection > 0 {
			// move up one row
			p.Tail.Row -= 1
		} else {
			// move down one row
			p.Tail.Row += 1
		}
	} else if rowDirection > 0 && colDirection > 0 {
		// diagonal up left
		p.Tail.Row -= 1 // one row up
		p.Tail.Col -= 1 // move one col left
	} else if rowDirection < 0 && colDirection < 0 {
		// diagonal down right
		p.Tail.Row += 1 // one row down
		p.Tail.Col += 1 // move one col right
	} else if rowDirection > 0 && colDirection < 0 {
		// diagonal up right
		p.Tail.Row -= 1 // one row up
		p.Tail.Col += 1 // move one col right
	} else if rowDirection < 0 && colDirection > 0 {
		// diagonal down left
		p.Tail.Row += 1 // one row down
		p.Tail.Col -= 1 // move one col left
	} else {
		panic("if is invalid")
	}

	// detect head and move one tile into this direction
	p.Content[p.Tail.Row][p.Tail.Col].VisitedByTail = true
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
