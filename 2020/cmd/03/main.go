package main

import (
	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

func getTrees(matrix [][]string, columns, rows, right, down int) int {
	trees := 0
	column := 1
	for i := down; i < rows; i = i + down {
		column = column + right
		if column > columns {
			column = column - columns
		}
		item := matrix[i][column-1]
		if item == "#" {
			trees += 1
		}
	}
	return trees
}

func main() {
	input, err := internal.ReadFile("cmd/03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// generate matrix from input
	rows := 0
	columns := 0
	matrix := make([][]string, len(input))
	for i, line := range input {
		columns = 0
		row := i
		rows += 1
		matrix[row] = make([]string, len(line))
		for j, c := range line {
			column := j
			columns += 1
			matrix[row][column] = string(c)
		}
	}

	trees := getTrees(matrix, columns, rows, 3, 1)

	log.Infof("Part 1 Trees: %d", trees)

	trees = getTrees(matrix, columns, rows, 1, 1)
	trees *= getTrees(matrix, columns, rows, 3, 1)
	trees *= getTrees(matrix, columns, rows, 5, 1)
	trees *= getTrees(matrix, columns, rows, 7, 1)
	trees *= getTrees(matrix, columns, rows, 1, 2)

	log.Infof("Part 2 Trees: %d", trees)
}
