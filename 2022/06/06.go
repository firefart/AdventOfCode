package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Printf("Part1 index %d\n", findStartSequence(string(input), 4))
	fmt.Printf("Part2 index %d\n", findStartSequence(string(input), 14))
	return nil
}

func findStartSequence(input string, seqLen int) int {
	for i := 0; i < len(input); i++ {
		endIndex := i + seqLen
		if endIndex > len(input) {
			endIndex = len(input)
		}
		if charsUnique(input[i:endIndex]) {
			return endIndex
		}
	}
	return -1
}

func charsUnique(input string) bool {
	chars := make(map[rune]int)
	for _, j := range input {
		if _, ok := chars[j]; ok {
			return false
		}
		chars[j] = 1
	}
	return true
}
