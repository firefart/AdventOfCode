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
	trees, err := parseTrees(input)
	if err != nil {
		return err
	}
	count := countVisibleTrees(trees)
	fmt.Printf("Part 1: %d\n", count)
	scenicScore := bestScenicScore(trees)
	fmt.Printf("Part 2: %d\n", scenicScore)
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

	visRight, _ := visibleFromRight(trees, row, col)
	visLeft, _ := visibleFromLeft(trees, row, col)
	visTop, _ := visibleFromTop(trees, row, col)
	visBottom, _ := visibleFromBottom(trees, row, col)

	// fmt.Printf("%d/%d visibleFromRight: %t\n", row, col, visRight)
	// fmt.Printf("%d/%d visibleFromLeft: %t\n", row, col, visLeft)
	// fmt.Printf("%d/%d visibleFromTop: %t\n", row, col, visTop)
	// fmt.Printf("%d/%d visibleFromBottom: %t\n", row, col, visBottom)

	if visRight ||
		visLeft ||
		visTop ||
		visBottom {
		return true
	}

	return false
}

// returns true if visible from the direction and also the scenicscore
func visibleFromLeft(trees [][]int, row, col int) (bool, int) {
	tree := trees[row][col]
	count := 0
	for i := col - 1; i >= 0; i-- {
		count++
		if trees[row][i] >= tree {
			return false, count
		}
	}
	return true, count
}

// returns true if visible from the direction and also the scenicscore
func visibleFromRight(trees [][]int, row, col int) (bool, int) {
	tree := trees[row][col]
	count := 0
	for i := col + 1; i < len(trees[row]); i++ {
		count++
		if trees[row][i] >= tree {
			return false, count
		}
	}
	return true, count
}

// returns true if visible from the direction and also the scenicscore
func visibleFromTop(trees [][]int, row, col int) (bool, int) {
	tree := trees[row][col]
	count := 0
	for i := row - 1; i >= 0; i-- {
		count++
		if trees[i][col] >= tree {
			return false, count
		}
	}
	return true, count
}

// returns true if visible from the direction and also the scenicscore
func visibleFromBottom(trees [][]int, row, col int) (bool, int) {
	tree := trees[row][col]
	count := 0
	for i := row + 1; i < len(trees); i++ {
		count++
		if trees[i][col] >= tree {
			return false, count
		}
	}
	return true, count
}

func bestScenicScore(trees [][]int) int {
	var score []int
	for rowIndex, row := range trees {
		for colIndex := range row {
			score = append(score, scenicScore(trees, rowIndex, colIndex))
		}
	}

	highest := 0
	for _, s := range score {
		if s > highest {
			highest = s
		}
	}
	return highest
}

func scenicScore(trees [][]int, row, col int) int {
	_, right := visibleFromRight(trees, row, col)
	_, left := visibleFromLeft(trees, row, col)
	_, top := visibleFromTop(trees, row, col)
	_, bottom := visibleFromBottom(trees, row, col)

	return top * left * right * bottom
}
