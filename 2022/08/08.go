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
	trees, err := parseTrees(input)
	if err != nil {
		return err
	}
	count := countVisibleTrees(trees)
	fmt.Printf("Part 1: %d\n", count)
	return nil
}

func parseTrees(input []byte) ([][]int, error) {
	var trees [][]int
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		sLen := len(s)
		row := make([]int, sLen)
		for i, tree := range s {
			height, err := strconv.Atoi(string(tree))
			if err != nil {
				return nil, err
			}
			row[i] = height
		}
		trees = append(trees, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return trees, nil
}

func countVisibleTrees(trees [][]int) int {
	visible := 0
	for rowIndex, row := range trees {
		for colIndex := range row {
			if treeVisible(trees, rowIndex, colIndex) {
				visible += 1
			}
		}
	}
	return visible
}

func treeVisible(trees [][]int, row, col int) bool {
	// edge pieces are always visible
	if row == 0 || row == len(trees)-1 || col == 0 || col == len(trees[0])-1 {
		return true
	}

	// fmt.Printf("%d/%d visibleFromRight: %t\n", row, col, visibleFromRight(trees, row, col))
	// fmt.Printf("%d/%d visibleFromLeft: %t\n", row, col, visibleFromLeft(trees, row, col))
	// fmt.Printf("%d/%d visibleFromTop: %t\n", row, col, visibleFromTop(trees, row, col))
	// fmt.Printf("%d/%d visibleFromBottom: %t\n", row, col, visibleFromBottom(trees, row, col))

	if visibleFromRight(trees, row, col) ||
		visibleFromLeft(trees, row, col) ||
		visibleFromTop(trees, row, col) ||
		visibleFromBottom(trees, row, col) {
		return true
	}

	return false
}

func visibleFromLeft(trees [][]int, row, col int) bool {
	tree := trees[row][col]
	for i := col - 1; i >= 0; i-- {
		if trees[row][i] >= tree {
			return false
		}
	}
	return true
}

func visibleFromRight(trees [][]int, row, col int) bool {
	tree := trees[row][col]
	for i := col + 1; i < len(trees[row]); i++ {
		if trees[row][i] >= tree {
			return false
		}
	}
	return true
}

func visibleFromTop(trees [][]int, row, col int) bool {
	tree := trees[row][col]
	for i := row - 1; i >= 0; i-- {
		if trees[i][col] >= tree {
			return false
		}
	}
	return true
}

func visibleFromBottom(trees [][]int, row, col int) bool {
	tree := trees[row][col]
	for i := row + 1; i < len(trees); i++ {
		if trees[i][col] >= tree {
			return false
		}
	}
	return true
}
