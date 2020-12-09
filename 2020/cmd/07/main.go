package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

type bag struct {
	amount int
	name   string
}

func getSubBagCount(bags map[string][]bag, bagname string) int {
	b, ok := bags[bagname]
	if !ok {
		log.Fatalf("Bag %s not yet processed", bagname)
	}

	count := 1
	for _, subbag := range b {
		count += subbag.amount * getSubBagCount(bags, subbag.name)
	}
	return count
}

func hasGoldBagInBag(bags map[string][]bag, bagname string) bool {
	b, ok := bags[bagname]
	if !ok {
		log.Fatalf("Bag %s not yet processed", bagname)
	}

	for _, subbag := range b {
		if subbag.name == "shiny gold" {
			return true
		}
	}

	for _, subbag := range b {
		if hasGoldBagInBag(bags, subbag.name) {
			return true
		}
	}
	return false
}

func main() {
	input, err := internal.ReadFile("cmd/07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^([\w\s]+) bags contain ([\w\s,]+).$`)
	bagrex := regexp.MustCompile(`^(\d+) ([\w\s]+) bags?$`)

	set := make(map[string][]bag)

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
				set[mainColor] = []bag{}
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
			set[mainColor] = append(set[mainColor], bag{
				amount: amInt,
				name:   subColor,
			})
		}
	}

	count := 0
	for outerBagName := range set {
		if hasGoldBagInBag(set, outerBagName) {
			count++
		}
	}
	log.Infof("Part 1: %d", count)

	// -1: Subtract to top golden bag
	log.Infof("Part 2: %d", getSubBagCount(set, "shiny gold")-1)
}
