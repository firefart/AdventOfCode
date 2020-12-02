package main

import (
	"regexp"
	"strconv"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	input, err := internal.ReadFile("cmd/02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	l := len(input)

	valid1 := 0
	valid2 := 0

	// part 1
	for i := 0; i < l; i++ {
		// parse input: 9-12 q: qqqxhnhdmqqqqjz
		re := regexp.MustCompile(`^(\d+)\-(\d+) (\w): (\w+)`)
		match := re.FindStringSubmatch(input[i])
		min := match[1]
		max := match[2]
		letter := match[3]
		password := match[4]

		minInt, err := strconv.Atoi(min)
		if err != nil {
			log.Fatal(err)
		}
		maxInt, err := strconv.Atoi(max)
		if err != nil {
			log.Fatal(err)
		}

		count := 0
		for _, c := range password {
			if string(c) == letter {
				count += 1
			}
		}

		if count < minInt || count > maxInt {
		} else {
			valid1 += 1
		}
	}
	log.Infof("Valid passwords part 1: %d", valid1)

	// part 2
	for i := 0; i < l; i++ {
		// parse input: 9-12 q: qqqxhnhdmqqqqjz
		re := regexp.MustCompile(`^(\d+)\-(\d+) (\w): (\w+)`)
		match := re.FindStringSubmatch(input[i])
		pos1 := match[1]
		pos2 := match[2]
		letter := match[3]
		password := match[4]

		pos1Int, err := strconv.Atoi(pos1)
		if err != nil {
			log.Fatal(err)
		}
		pos2Int, err := strconv.Atoi(pos2)
		if err != nil {
			log.Fatal(err)
		}

		l1 := string(password[pos1Int-1])
		l2 := string(password[pos2Int-1])

		if (l1 == letter || l2 == letter) && !(l1 == letter && l2 == letter) {
			valid2 += 1
		}
	}
	log.Infof("Valid passwords part 2: %d", valid2)
}
