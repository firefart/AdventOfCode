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

func execute(origInstructions []instruction) (int, bool) {
	// always work with a copy of the slice here
	// as we will change the underlying data otherwise
	instructions := make([]instruction, len(origInstructions))
	copy(instructions, origInstructions)

	accumulator := 0
	stackpointer := 0
	success := false
	for {
		if stackpointer == len(instructions) {
			success = true
			break
		}

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
			log.Fatalf("unknown instruction %s", ins.ins)
		}
	}
	return accumulator, success
}

func main() {
	input, err := internal.ReadFile("cmd/08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := make([]instruction, len(input))

	parseRegex := regexp.MustCompile(`^(\w+) ([-+])(\d+)$`)

	for i, line := range input {
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
		instructions[i] = instruction{
			ins:    match[1],
			amount: amount,
			hit:    false,
		}
	}

	accumulator, _ := execute(instructions)
	log.Infof("Part 1: %d", accumulator)

	for i := 0; i < len(instructions); i++ {
		// always work with a copy of the slice here
		// as we will change the underlying data otherwise
		instructionsCopy := make([]instruction, len(instructions))
		copy(instructionsCopy, instructions)
		x := instructionsCopy[i]
		newIns := ""
		if x.ins == "jmp" {
			newIns = "nop"
		} else if x.ins == "nop" {
			newIns = "jmp"
		} else {
			continue
		}
		tmp := instructionsCopy
		tmp[i].ins = newIns
		accumulatorPart2, success := execute(tmp)
		if success {
			// we exited normally
			log.Infof("Part 2: %d - Changed %d to %q", accumulatorPart2, i, newIns)
		}
	}
}
