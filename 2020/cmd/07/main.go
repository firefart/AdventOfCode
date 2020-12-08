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

func resolveColors(set map[string]bag, topColors []string) map[string]struct{} {
	res := make(map[string]struct{})
	for _, color := range topColors {
		x := set[color]
		// add the main color
		res[color] = struct{}{}
		// recursive call for the other colors
		y := resolveColors(set, x.containedInColor)
		for k := range y {
			res[k] = struct{}{}
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
			if val, ok := set[subColor]; ok {
				val.containedInColor = append(val.containedInColor, mainColor)
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
