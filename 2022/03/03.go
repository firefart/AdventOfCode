package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err := logic(content); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func logic(input []byte) error {
	if err := part1(input); err != nil {
		return err
	}
	if err := part2(input); err != nil {
		return err
	}
	return nil
}

func part1(input []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	overallPrio := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		diff := findDifference(s)
		subPrio := calculatePriority(diff)
		overallPrio += subPrio
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Printf("Part 1 Overall Prio: %d\n", overallPrio)
	return nil
}

func part2(input []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	overallPrio := 0
	var rucksacks []string
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		rucksacks = append(rucksacks, s)

		if len(rucksacks) == 3 {
			badge := findBadge(rucksacks)
			subPrio := calculatePriority(badge)
			overallPrio += subPrio
			rucksacks = []string{}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Printf("Part 2 Overall Prio: %d\n", overallPrio)
	return nil
}

func findDifference(in string) rune {
	inLen := len(in)
	if inLen%2 != 0 {
		panic("invalid input string")
	}
	part1 := in[0 : inLen/2]
	part2 := in[inLen/2 : inLen]

	for _, c := range part1 {
		if strings.ContainsRune(part2, c) {
			return c
		}
	}

	for _, c := range part2 {
		if strings.ContainsRune(part1, c) {
			return c
		}
	}

	panic("are the same")
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func calculatePriority(in rune) int {
	ret := strings.IndexRune(alphabet, in)
	if ret == -1 {
		panic("invalid rune")
	}
	return ret + 1
}

func findBadge(rucksacks []string) rune {
	if len(rucksacks) != 3 {
		panic("invalid rucksack len")
	}
	for _, c := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], c) && strings.ContainsRune(rucksacks[2], c) {
			return c
		}
	}
	panic("no badge found")
}
