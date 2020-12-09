package main

import (
	"strconv"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

const preambleLength = 25

func minMax(input []int) (int, int) {
	if len(input) == 0 {
		return 0, 0
	}
	max := input[0]
	min := input[0]
	for _, v := range input {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return min, max
}

func checkNumber(preamble []int, number int) bool {
	for _, x := range preamble {
		for _, y := range preamble {
			if x+y == number {
				return true
			}
		}
	}
	return false
}

func getSumSlice(input []int, target int) []int {
	done := false
	start := 0

	for !done {
		sum := 0
		var arr []int
		for i := start; i < len(input); i++ {
			sum += input[i]
			arr = append(arr, input[i])
			if sum == target && len(arr) > 1 {
				return arr
			} else if sum > target {
				break
			}
		}
		start++
		if start > len(input) {
			done = true
		}
	}
	return nil
}

func main() {
	inputString, err := internal.ReadFile("cmd/09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := make([]int, len(inputString))
	for i := range inputString {
		number, err := strconv.Atoi(inputString[i])
		if err != nil {
			log.Fatal(err)
		}
		input[i] = number
	}

	start := preambleLength

	for start < len(input) {
		preamble := input[start-preambleLength : start]
		number := input[start]
		// log.Infof("Number: %d", number)
		// log.Infof("Start: %d", start)
		// log.Infof("Preamble: %v", preamble)
		valid := checkNumber(preamble, number)
		if !valid {
			log.Infof("Number %d is invalid", number)
			part2 := getSumSlice(input, number)
			if part2 != nil {
				min, max := minMax(part2)
				log.Infof("Part 2: %d", min+max)
			}
			break
		}
		start++
	}
}
