package main

import (
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	input, err := internal.ReadFile("cmd/06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	countPart1 := 0
	countPart2 := 0
	peoplePerGroup := 0
	set := make(map[rune]int)
	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			// new group
			// overall count for part 1
			countPart1 += len(set)

			// get count for part 2 where all people had the same answer
			for _, v := range set {
				if v == peoplePerGroup {
					countPart2 += 1
				}
			}

			peoplePerGroup = 0
			set = make(map[rune]int)
			continue
		}

		for _, char := range line {
			set[char] = set[char] + 1
		}
		peoplePerGroup += 1
	}
	// last lines contains no new line so add last counts again
	countPart1 += len(set)
	for _, v := range set {
		if v == peoplePerGroup {
			countPart2 += 1
		}
	}

	log.Infof("Part 1: %d", countPart1)
	log.Infof("Part 2: %d", countPart2)
}
