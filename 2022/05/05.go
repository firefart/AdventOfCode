package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type matrix struct {
	content [][]rune
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
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var crateString strings.Builder
	var inputString strings.Builder
	crateSection := true
	for scanner.Scan() {
		s := scanner.Text()
		if strings.TrimSpace(s) == "" {
			crateSection = false
			continue
		}

		if crateSection {
			crateString.WriteString(fmt.Sprintf("%s\n", s))
		} else {
			inputString.WriteString(fmt.Sprintf("%s\n", s))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	cratesPart1 := matrix{content: parseCrates(crateString.String())}
	cratesPart2 := matrix{content: cloneMatrix(cratesPart1.content)}

	printMatrix(cratesPart1.content)

	for _, command := range strings.Split(inputString.String(), "\n") {
		if strings.TrimSpace(command) == "" {
			continue
		}
		commandSplit := strings.Split(command, " ")
		if len(commandSplit) != 6 {
			panic(fmt.Sprintf("invalid command %q", command))
		}
		howManyS, fromS, toS := commandSplit[1], commandSplit[3], commandSplit[5]
		howMany, err := strconv.Atoi(howManyS)
		if err != nil {
			return err
		}
		from, err := strconv.Atoi(fromS)
		if err != nil {
			return err
		}
		to, err := strconv.Atoi(toS)
		if err != nil {
			return err
		}
		cratesPart1.moveCratePart1(howMany, from, to)
		cratesPart2.moveCratePart2(howMany, from, to)
	}

	printMatrix(cratesPart1.content)
	fmt.Printf("Part1: %s\n", cratesPart1.getTopCrateString())
	printMatrix(cratesPart2.content)
	fmt.Printf("Part2: %s\n", cratesPart2.getTopCrateString())

	return nil
}

func parseCrates(in string) [][]rune {
	tmp := strings.Split(in, "\n")
	var s []string
	for _, l := range tmp {
		if strings.TrimSpace(l) != "" {
			s = append(s, l)
		}
	}

	numRows := len(s) - 1
	lastRow := s[len(s)-1]
	numCols := (len(lastRow) + 1) / 4

	matrix := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		// create column
		matrix[i] = make([]rune, numCols)

		rowContent := s[i]
		colCount := 0
		for j := 0; j < len(rowContent); j += 4 {
			field := rowContent[j : j+3]
			content := []rune(field)[1]
			matrix[i][colCount] = content
			colCount++
		}
	}
	return matrix
}

func (m *matrix) addRow() {
	new := make([][]rune, len(m.content)+1)
	new[0] = make([]rune, len(m.content[0]))
	for i := range m.content[0] {
		new[0][i] = ' '
	}
	for i, row := range m.content {
		new[i+1] = row
	}
	m.content = new
}

func printMatrix(matrix [][]rune) {
	rows := len(matrix)
	cols := len(matrix[0])
	var sb strings.Builder
	for _, row := range matrix {
		for _, col := range row {
			sb.WriteString(fmt.Sprintf(" [ %c ] ", col))
		}
		sb.WriteString("\n")
	}
	for i := 0; i < cols; i++ {
		sb.WriteString(fmt.Sprintf("   %d   ", i))
	}
	fmt.Printf("Rows: %d\n", rows)
	fmt.Printf("Columns: %d\n", cols)
	fmt.Println(sb.String())
}

func (m *matrix) moveCratePart1(howMany, from, to int) {
	for i := 0; i < howMany; i++ {
		topCrateObj := m.topCrateInColumn(from - 1)
		rowsAdded := m.moveStackToColumn([]crate{topCrateObj}, to-1)
		// remove crate in old location
		m.content[topCrateObj.index+rowsAdded][from-1] = ' '
	}
}

func (m *matrix) moveCratePart2(howMany, from, to int) {
	// printMatrix(m.content)
	// get crate stack
	crateStack := m.topCratesInColumn(from-1, howMany)
	// move crate stack to new location
	rowsAdded := m.moveStackToColumn(crateStack, to-1)
	// remove crate stack in old location
	for _, crate := range crateStack {
		m.content[crate.index+rowsAdded][from-1] = ' '
	}
	// printMatrix(m.content)
}

func (m *matrix) emptyCrate(column int) int {
	// bottom up
	for i := len(m.content) - 1; i >= 0; i-- {
		row := m.content[i]
		if row[column] == ' ' {
			return i
		}
	}
	return -1
}

type crate struct {
	rune  rune
	index int
}

func (m *matrix) topCrateInColumn(column int) crate {
	return m.topCratesInColumn(column, 1)[0]
}

func (m *matrix) topCratesInColumn(column, howMany int) []crate {
	// printMatrix(m.content)
	// fmt.Printf("Col: %d\n", column)
	// fmt.Printf("Amount: %d\n", howMany)
	totalCratesInColumn := m.totalCratesInColumn(column)
	maxIter := howMany
	if maxIter > totalCratesInColumn {
		maxIter = totalCratesInColumn
	}
	var ret []crate
	for i, row := range m.content {
		if row[column] != ' ' {
			for j := 0; j < maxIter; j++ {
				// fmt.Printf("i: %d\n", i)
				// fmt.Printf("j: %d\n", j)
				// fmt.Printf("rune: %c\n", m.content[i+j][column])
				ret = append(ret, crate{
					index: i + j,
					rune:  m.content[i+j][column],
				})
			}
			return ret
		}
	}
	return ret
}

func (m *matrix) totalCratesInColumn(column int) int {
	total := 0
	for _, row := range m.content {
		if row[column] != ' ' {
			total++
		}
	}
	return total
}

func (m *matrix) numberOfRows() int {
	return len(m.content)
}

func (m *matrix) moveStackToColumn(crateStack []crate, to int) int {
	// printMatrix(m.content)

	numberOfElementsInTargetColumn := m.numberOfRows() - m.emptyCrate(to) - 1
	rowDifference := (numberOfElementsInTargetColumn + len(crateStack)) - m.numberOfRows()
	rowsAdded := 0
	// extend matrix if equasion is positive
	for i := rowDifference; i > 0; i-- {
		rowsAdded++
		m.addRow()
	}
	// fmt.Printf("Added %d rows\n", rowsAdded)
	// get index of first empty crate from bottom up
	firstEmptyCrate := m.emptyCrate(to)
	// printMatrix(m.content)
	// fmt.Printf("Stack to move: %v\n", crateStack)
	// fmt.Printf("First Empty Crate: %d\n", firstEmptyCrate)
	// fmt.Printf("To Column: %d\n", to)
	for i := 1; i <= len(crateStack); i++ {
		targetRow := firstEmptyCrate - i + 1

		runeToMove := crateStack[len(crateStack)-i]
		// fmt.Printf("TargetRow: %d\n", targetRow)
		// fmt.Printf("RuneToMove: %c\n", runeToMove.rune)
		m.content[targetRow][to] = runeToMove.rune
	}
	// printMatrix(m.content)
	return rowsAdded
}

func cloneMatrix(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func (m *matrix) getTopCrateString() string {
	var ret strings.Builder
	for i := 0; i < len(m.content[0]); i++ {
		r := m.topCrateInColumn(i)
		_, err := ret.WriteRune(r.rune)
		if err != nil {
			panic(err)
		}
	}
	return ret.String()
}
