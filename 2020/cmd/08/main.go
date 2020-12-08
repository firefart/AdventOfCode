package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/firefart/adventofcode/internal"
	log "github.com/sirupsen/logrus"
)

type instruction struct {
	ins    string
	amount int
	hit    bool
}

func main() {
	input, err := internal.ReadFile("cmd/08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instructions []instruction

	parseRegex := regexp.MustCompile(`^(\w+) ([-+])(\d+)$`)

	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		match := parseRegex.FindStringSubmatch(line)
		if match == nil {
			log.Fatalf("unmatched input %q", line)
		}
		amount, err := strconv.Atoi(match[3])
		if err != nil {
			log.Fatal(err)
		}
		if match[2] == "-" {
			amount *= -1
		}
		instructions = append(instructions, instruction{
			ins:    match[1],
			amount: amount,
			hit:    false,
		})
	}

	accumulator := 0
	stackpointer := 0
	for {
		ins := instructions[stackpointer]
		if ins.hit {
			break
		}
		instructions[stackpointer].hit = true
		switch ins.ins {
		case "nop":
			stackpointer++
		case "jmp":
			stackpointer += ins.amount
		case "acc":
			accumulator += ins.amount
			stackpointer++
		default:
			log.Fatalf("unknown instruction %s", ins)
		}
	}
	log.Infof("Part 1: %d", accumulator)
}
