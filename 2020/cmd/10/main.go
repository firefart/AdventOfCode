package main

import (
	"sort"
	"strconv"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

type adapter struct {
	joltage int
}

func main() {
	inputString, err := internal.ReadFile("cmd/10/input.txt")
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

	log.Print(input)

	sort.Ints(input)

	log.Print(input)

	builtInAdapter := input[len(input)-1] + 3
	// add our adapter to the list
	input = append(input, builtInAdapter)

	start := 0
	oneJoltDiff := 0
	twoJoltDiff := 0
	threeJoltDiff := 0
	for _, i := range input {
		diff := i - start
		// log.Info(i, " ", start, " ", diff)
		switch diff {
		case 1:
			oneJoltDiff++
		case 2:
			twoJoltDiff++
		case 3:
			threeJoltDiff++
		default:
			log.Fatalf("Missing adapter! %d", diff)
		}
		start = i
	}

	log.Infof("One Jolt Diff: %d", oneJoltDiff)
	log.Infof("Two Jolt Diff: %d", twoJoltDiff)
	log.Infof("Three Jolt Diff: %d", threeJoltDiff)
	log.Infof("Part 1: %d", oneJoltDiff*threeJoltDiff)
}
