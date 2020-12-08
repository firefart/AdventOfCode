package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

type bag struct {
	count            int
	containedInColor []string
}

func resolveColors(set map[string]bag, topColors []string) map[string]int {
	res := make(map[string]int)
	for _, color := range topColors {
		x := set[color]
		// add the main color
		if v, ok := res[color]; ok {
			res[color] = v + x.count
		} else {
			res[color] = x.count
		}

		// recursive call for the other colors
		y := resolveColors(set, x.containedInColor)
		for k, v := range y {
			res[k] = v
		}
	}
	return res
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
				set[mainColor] = bag{
					count: 0,
				}
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
			if val, ok := set[subColor]; ok {
				val.containedInColor = append(val.containedInColor, mainColor)
				val.count += amInt
				set[subColor] = val
			} else {
				set[subColor] = bag{
					count:            amInt,
					containedInColor: []string{mainColor},
				}
			}
		}
	}
	top := set["shiny gold"]
	res := resolveColors(set, top.containedInColor)
	log.Infof("Part 1: %d", len(res))

	log.Info(res)
}
