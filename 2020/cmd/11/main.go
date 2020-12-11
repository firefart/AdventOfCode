package main

import (
	"bytes"
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

type seatPlan [][]string

func (s seatPlan) occupiedSeats() int {
	seats := 0
	for row := range s {
		for column := 0; column < len(s[row]); column++ {
			if s[row][column] == "#" {
				seats++
			}
		}
	}
	return seats
}

func (s seatPlan) print(iteration int) {
	if len(s) == 0 {
		return
	}
	numberOfCols := len(s[0])
	var b bytes.Buffer
	for row := range s {
		for column := 0; column < numberOfCols; column++ {
			b.WriteString(s[row][column])
		}
		b.WriteString("\n")
	}
	log.Infof("Seatplan iteration %d:\n%s", iteration, b.String())
}

func (s seatPlan) same(new seatPlan) bool {
	if len(s) != len(new) {
		return false
	}
	for row := range s {
		if len(s[row]) != len(new[row]) {
			return false
		}
		for column := 0; column < len(new[row]); column++ {
			if s[row][column] != new[row][column] {
				return false
			}
		}
	}
	return true
}

func (s seatPlan) iterate() seatPlan {
	newPlan := make(seatPlan, len(s))
	numberOfCols := len(s[0])
	for row := range s {
		newPlan[row] = make([]string, len(s[row]))
		for column := 0; column < numberOfCols; column++ {
			newPlan[row][column] = s.checkAndChangeState(row, column)
		}
	}
	return newPlan
}

func (s seatPlan) checkAndChangeState(row, column int) string {
	seat := s[row][column]
	if seat == "." {
		return "."
	}
	// log.Info(row, " ", column)
	adj := s.numberOfAdjacentSeats(row, column)
	// log.Infof("Seat %d,%d: %d", row, column, adj)
	if seat == "L" && adj == 0 {
		return "#"
	} else if seat == "#" && adj >= 4 {
		return "L"
	} else {
		return s[row][column]
	}
}

func (s seatPlan) numberOfAdjacentSeats(row, column int) int {
	seatCount := 0
	leftColumn := column - 1
	rightColumn := column + 1
	maxRowIndex := len(s) - 1
	maxColIndex := len(s[row]) - 1

	// top row
	topRow := row - 1
	if topRow >= 0 {
		if leftColumn >= 0 {
			if s[topRow][leftColumn] == "#" {
				seatCount++
			}
		}
		if s[topRow][column] == "#" {
			seatCount++
		}
		if rightColumn <= maxColIndex {
			if s[topRow][rightColumn] == "#" {
				seatCount++
			}
		}
	}
	// middle row
	middleRow := row
	if leftColumn >= 0 {
		if s[middleRow][leftColumn] == "#" {
			seatCount++
		}
	}
	// do not count the current seat
	if rightColumn <= maxColIndex {
		if s[middleRow][rightColumn] == "#" {
			seatCount++
		}
	}
	// bottom row
	bottomRow := row + 1
	if bottomRow <= maxRowIndex {
		if leftColumn >= 0 {
			if s[bottomRow][leftColumn] == "#" {
				seatCount++
			}
		}
		if s[bottomRow][column] == "#" {
			seatCount++
		}
		if rightColumn <= maxColIndex {
			if s[bottomRow][rightColumn] == "#" {
				seatCount++
			}
		}
	}
	return seatCount
}

func main() {
	input, err := internal.ReadFile("cmd/11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// var seats seatPlan
	seats := make(seatPlan, len(input))
	for row, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		seats[row] = make([]string, len(line))
		for column, char := range line {
			seats[row][column] = string(char)
		}
	}

	differ := true
	oldSeats := seats
	var newSeats seatPlan
	// initial state
	oldSeats.print(0)
	count := 1

	for differ {
		newSeats = oldSeats.iterate()
		if newSeats.same(oldSeats) {
			// last iteration did not change anything
			differ = false
		} else {
			newSeats.print(count)
			oldSeats = newSeats
			count++
		}
	}
	seatCount := newSeats.occupiedSeats()
	log.Infof("Took %d rounds. Seatcount: %d", count, seatCount)
}
