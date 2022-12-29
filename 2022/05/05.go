package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	content, err := ioutil.ReadAll(f)
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

	cratesPart1 := parseCrates(crateString.String())
	cratesPart2 := cloneMatrix(cratesPart1)

	printMatrix(cratesPart1)

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
		cratesPart1 = moveCratePart1(howMany, from, to, cratesPart1)
		cratesPart2 = moveCratePart2(howMany, from, to, cratesPart2)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("Part1: %s\n", getTopCrateString(cratesPart1))
	fmt.Printf("Part2: %s\n", getTopCrateString(cratesPart2))

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

func addRow(crates [][]rune) [][]rune {
	new := make([][]rune, len(crates)+1)
	new[0] = make([]rune, len(crates[0]))
	for i := range crates[0] {
		new[0][i] = ' '
	}
	for i, row := range crates {
		new[i+1] = row
	}
	return new
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
	for i := 1; i <= cols; i++ {
		sb.WriteString(fmt.Sprintf("   %d   ", i))
	}
	fmt.Printf("Rows: %d\n", rows)
	fmt.Printf("Columns: %d\n", cols)
	fmt.Println(sb.String())
}

func moveCratePart1(howMany, from, to int, crates [][]rune) [][]rune {
	for i := 0; i < howMany; i++ {
		topCrateObj := topCrate(from-1, crates)
		topCrateRune := topCrateObj.rune
		topCrateIndex := topCrateObj.index
		emptyCrateIndex := emptyCrate(to-1, crates)
		if emptyCrateIndex == -1 {
			crates = addRow(crates)
			emptyCrateIndex = 0
			topCrateIndex += 1
		}
		crates[emptyCrateIndex][to-1] = topCrateRune
		crates[topCrateIndex][from-1] = ' '
		// printMatrix(crates)
	}
	return crates
}

func moveCratePart2(howMany, from, to int, crates [][]rune) [][]rune {
	topCrateObjs := topCrates(from-1, howMany, crates)
	topCrateRune := topCrateObj.rune
	topCrateIndex := topCrateObj.index
	emptyCrateIndex := emptyCrate(to-1, crates)
	if emptyCrateIndex == -1 {
		crates = addRow(crates)
		emptyCrateIndex = 0
		topCrateIndex += 1
	}
	crates[emptyCrateIndex][to-1] = topCrateRune
	crates[topCrateIndex][from-1] = ' '
	// printMatrix(crates)
	return crates
}

func emptyCrate(column int, crates [][]rune) int {
	// bottom up
	for i := len(crates) - 1; i >= 0; i-- {
		row := crates[i]
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

func topCrate(column int, crates [][]rune) crate {
	return topCrates(column, 1, crates)[0]
}

func topCrates(column, howMany int, crates [][]rune) []crate {
	// top down
	var ret []crate
	for i, row := range crates {
		if row[column] != ' ' {
			for j := 0; j < howMany; j++ {
				ret = append(ret, crate{
					index: i + j,
					rune:  crates[i+j][column],
				})
			}
		}
	}
	return ret
}

func cloneMatrix(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func getTopCrateString(crates [][]rune) string {
	var ret strings.Builder
	for i := 0; i < len(crates[0]); i++ {
		r, _ := topCrate(i, crates)
		_, err := ret.WriteRune(r)
		if err != nil {
			panic(err)
		}
	}
	return ret.String()
}
