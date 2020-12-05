package main

import (
	"sort"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

func getSeatID(row, column int) int {
	return row*8 + column
}

func getSeatRowAndColumn(input string) (int, int) {
	first := input[0:7] // indicates column
	last := input[7:]   // indicates row

	row := 0
	column := 0

	low := 0
	high := 128 // number of rows in the plane

	for _, c := range first {
		x := string(c)
		median := (low + high) / 2
		switch x {
		case "F":
			high = median
		case "B":
			low = median
		default:
			log.Fatalf("invalid character %s", x)
		}
	}
	row = low

	low = 0
	high = 8 // number of columns in the plane

	for _, c := range last {
		x := string(c)
		median := (low + high) / 2
		switch x {
		case "L":
			high = median
		case "R":
			low = median
		default:
			log.Fatalf("invalid character %s", x)
		}
	}
	column = low

	// log.Infof("%s %s %s %d %d", input, first, last, row, column)
	return row, column
}

func getFirstMissingNumber(input []int) int {
	sort.Ints(input)
	last := len(input) - 1
	for idx, currentItem := range input {
		if idx >= last {
			break
		}
		nextItem := input[idx+1]
		if nextItem-currentItem == 1 {
			continue
		}

		// if we reach here we've found the missing item
		missing := (nextItem + currentItem) / 2
		return missing
	}
	return -1
}

func main() {
	input, err := internal.ReadFile("cmd/05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var seatIDs []int
	maxSeatID := 0
	for _, line := range input {
		row, column := getSeatRowAndColumn(line)
		id := getSeatID(row, column)
		seatIDs = append(seatIDs, id)
		if id > maxSeatID {
			maxSeatID = id
		}
	}

	log.Infof("Part1: Max Seat ID: %d", maxSeatID)

	missing := getFirstMissingNumber(seatIDs)

	log.Infof("Part2: Missing ID: %d", missing)
}
