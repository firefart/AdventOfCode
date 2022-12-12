package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
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
	elves, err := parseElves(input)
	if err != nil {
		return err
	}
	fmt.Printf("Biggest: %d\n", biggest(elves))
	fmt.Printf("Top 3: %d\n", top3Sum(elves))
	return nil
}

func parseElves(input []byte) (map[int]int, error) {
	elves := make(map[int]int)
	scanner := bufio.NewScanner(bytes.NewReader(input))
	currentElf := 1
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			currentElf++
			continue
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		elves[currentElf] += i
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return elves, nil
}

func biggest(in map[int]int) int {
	biggest := 0
	for _, v := range in {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

func top3Sum(in map[int]int) int {
	var arr []int
	for _, v := range in {
		arr = append(arr, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	top := arr[0:3]
	top3 := 0
	for _, i := range top {
		top3 += i
	}
	return top3
}
