package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err := logic(content); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func logic(input []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	part1 := 0
	part2 := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		overlap := findOverlapComplete(s)
		if overlap {
			part1++
		}
		overlap2 := findOverlap(s)
		if overlap2 {
			part2++
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Printf("Part 1 Overlap Count: %d\n", part1)
	fmt.Printf("Part 2 Overlap Count: %d\n", part2)
	return nil
}

func findOverlapComplete(in string) bool {
	s := strings.Split(in, ",")
	if len(s) != 2 {
		panic(fmt.Sprintf("invalid input %s", in))
	}
	elv1S, elv2S := s[0], s[1]
	elv1 := parseRange(elv1S)
	elv2 := parseRange(elv2S)

	if strings.Contains(elv1, elv2) || strings.Contains(elv2, elv1) {
		return true
	}

	return false
}

func findOverlap(in string) bool {
	s := strings.Split(in, ",")
	if len(s) != 2 {
		panic(fmt.Sprintf("invalid input %s", in))
	}
	elv1S, elv2S := s[0], s[1]
	elv1 := parseRangeToInt(elv1S)
	elv2 := parseRangeToInt(elv2S)

	for _, i := range elv1 {
		if contains(i, elv2) {
			return true
		}
	}

	for _, i := range elv2 {
		if contains(i, elv1) {
			return true
		}
	}

	return false
}

func parseRange(in string) string {
	s := strings.Split(in, "-")
	if len(s) != 2 {
		panic(fmt.Sprintf("invalid input %s", in))
	}
	startS, endS := s[0], s[1]
	start, err := strconv.Atoi(startS)
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(endS)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	for i := start; i <= end; i++ {
		c := strconv.Itoa(i)
		sb.WriteString(fmt.Sprintf("|%s|", c))
	}
	return strings.ReplaceAll(sb.String(), "||", "|")
}

func parseRangeToInt(in string) []int {
	s := strings.Split(in, "-")
	if len(s) != 2 {
		panic(fmt.Sprintf("invalid input %s", in))
	}
	startS, endS := s[0], s[1]
	start, err := strconv.Atoi(startS)
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(endS)
	if err != nil {
		panic(err)
	}

	var ret []int
	for i := start; i <= end; i++ {
		ret = append(ret, i)
	}
	return ret
}

func contains[T comparable](in T, slice []T) bool {
	for _, x := range slice {
		if in == x {
			return true
		}
	}
	return false
}
