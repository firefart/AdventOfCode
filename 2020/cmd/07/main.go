package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

type bag struct {
	bags map[string]int
}

func execute(set map[string]bag, topElement string) map[string]bag {
	ret := make(map[string]bag)
	workingElement := set[topElement]
	ret[topElement] = workingElement
	for k, v := range workingElement.bags {
		log.Info(k, " ", v)
		x := execute(set, k)
		for k2, v2 := range x {
			ret[k2] = v2
		}
	}
	return ret
}

func countBags(set map[string]bag, name string) int {
	workingElement := set[name]
	count := 0
	for k, v := range workingElement.bags {
		count += v
		count += countBags(set, k)
	}
	count *= len(workingElement.bags)
	return count
}

func main() {
	input, err := internal.ReadFile("cmd/07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^([\w\s]+) bags contain ([\w\s,]+).$`)
	bagrex := regexp.MustCompile(`^(\d+) ([\w\s]+) bags?$`)

	set := make(map[string]bag)

	for _, line := range input {
		match := re.FindStringSubmatch(line)
		if match == nil {
			log.Infof("unmatched input: %s", line)
			continue
		}
		mainColor := match[1]
		contains := match[2]
		parts := strings.Split(contains, ", ")
		for _, part := range parts {
			if part == "no other bags" {
				set[mainColor] = bag{}
				continue
			}
			match := bagrex.FindStringSubmatch(part)
			if match == nil {
				log.Infof("unmatched input: %s", part)
				continue
			}
			amount := match[1]
			subColor := match[2]
			amInt, err := strconv.Atoi(amount)
			if err != nil {
				log.Fatal(err)
			}
			if val, ok := set[mainColor]; ok {
				if len(val.bags) == 0 {
					val.bags = make(map[string]int)
				}
				val.bags[subColor] = amInt
				set[mainColor] = val
			} else {
				bags := make(map[string]int)
				bags[subColor] = amInt
				set[mainColor] = bag{
					bags: bags,
				}
			}
		}
	}
	x := execute(set, "shiny gold")
	log.Info(x)
	log.Infof("Part 1: %d", len(x))
	// log.Info(set)
	y := countBags(set, "shiny gold")
	log.Info(y)
}
