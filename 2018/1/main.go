package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	run := 0
	i := 0
	alreadyFound := make(map[int]bool)
	for {
		f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			i += value
			if alreadyFound[i] {
				log.Printf("Part 2: %d\n", i)
				if run > 0 {
					os.Exit(0)
				}
			}
			alreadyFound[i] = true
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if run == 0 {
			log.Println(i)
		}
		run++
	}
}
