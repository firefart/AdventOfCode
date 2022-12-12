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
	score, err := playAllRounds(input)
	if err != nil {
		return err
	}
	fmt.Printf("Part 1 Score: %d\n", score)
	score, err = playAllRoundsPart2(input)
	if err != nil {
		return err
	}
	fmt.Printf("Part 2 Score: %d\n", score)
	return nil
}

func playAllRounds(input []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	overallScore := 0
	for scanner.Scan() {
		s := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		opponentMove, myMove := translateMove(s[0]), translateMove(s[1])
		res := playRound(opponentMove, myMove)
		subScore := calculateRoundScore(res, myMove)
		overallScore += subScore
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return overallScore, nil
}

func playAllRoundsPart2(input []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	overallScore := 0
	for scanner.Scan() {
		s := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		opponentMove, expectedresult := translateMove(s[0]), translateResult(s[1])
		myMove := fakeResult(expectedresult, opponentMove)
		res := playRound(opponentMove, myMove)
		subScore := calculateRoundScore(res, myMove)
		overallScore += subScore
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return overallScore, nil
}

type move int
type result int

const (
	moveRock          move   = iota
	movePaper         move   = iota
	moveScissors      move   = iota
	resultWinOpponent result = iota
	resultWinMe       result = iota
	resultDraw        result = iota
)

func playRound(opponentMove, myMove move) result {
	if opponentMove == myMove {
		return resultDraw
	}

	switch opponentMove {
	case moveRock:
		switch myMove {
		case moveRock:
			return resultDraw
		case movePaper:
			return resultWinMe
		case moveScissors:
			return resultWinOpponent
		default:
			panic("invalid move")
		}
	case movePaper:
		switch myMove {
		case moveRock:
			return resultWinOpponent
		case movePaper:
			return resultDraw
		case moveScissors:
			return resultWinMe
		default:
			panic("invalid move")
		}
	case moveScissors:
		switch myMove {
		case moveRock:
			return resultWinMe
		case movePaper:
			return resultWinOpponent
		case moveScissors:
			return resultDraw
		default:
			panic("invalid move")
		}
	default:
		panic("invalid move")
	}
}

func calculateRoundScore(res result, myChoice move) int {
	var score int
	switch myChoice {
	case moveRock:
		score = 1
	case movePaper:
		score = 2
	case moveScissors:
		score = 3
	default:
		panic("invalid move")
	}

	switch res {
	case resultDraw:
		score += 3
	case resultWinMe:
		score += 6
	case resultWinOpponent:
		score += 0
	default:
		panic("invalid result")
	}

	return score
}

func translateMove(in string) move {
	switch in {
	case "A":
		return moveRock
	case "B":
		return movePaper
	case "C":
		return moveScissors
	case "X":
		return moveRock
	case "Y":
		return movePaper
	case "Z":
		return moveScissors
	default:
		panic("invalid move")
	}
}

func translateResult(in string) result {
	switch in {
	case "X":
		return resultWinOpponent
	case "Y":
		return resultDraw
	case "Z":
		return resultWinMe
	default:
		panic("invalid result")
	}
}

func fakeResult(expectedResult result, opponentMove move) move {
	switch expectedResult {
	case resultDraw:
		switch opponentMove {
		case movePaper:
			return movePaper
		case moveScissors:
			return moveScissors
		case moveRock:
			return moveRock
		default:
			panic("invalid move")
		}
	case resultWinMe:
		switch opponentMove {
		case movePaper:
			return moveScissors
		case moveScissors:
			return moveRock
		case moveRock:
			return movePaper
		default:
			panic("invalid move")
		}
	case resultWinOpponent:
		switch opponentMove {
		case movePaper:
			return moveRock
		case moveScissors:
			return movePaper
		case moveRock:
			return moveScissors
		default:
			panic("invalid move")
		}
	default:
		panic("invalid expectedResult")
	}
}
