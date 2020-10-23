package main

import (
	"bufio"
	"log"
	"os"
)

func check(input string) (bool, bool) {
	counts := make(map[rune]int)
	for _, char := range input {
		if _, ok := counts[char]; ok {
			counts[char]++
		} else {
			counts[char] = 1
		}
	}

	var two, three bool
	for _, v := range counts {
		switch v {
		case 2:
			two = true
		case 3:
			three = true
		}
	}

	return two, three
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	twos := 0
	threes := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		two, three := check(value)
		//fmt.Printf("%s: %v %v\n", value, two, three)
		if two {
			twos++
		}
		if three {
			threes++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%d", twos*threes)
}
