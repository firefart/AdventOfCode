package main

import "testing"

var input = []byte(`A Y
B X
C Z`)

func TestPlayAllRounds(t *testing.T) {
	score, err := playAllRounds(input)
	if err != nil {
		t.Fatal(err)
	}
	if score != 15 {
		t.Errorf("got score of %d, want %d", score, 15)
	}
}

func TestCalculateRoundScore(t *testing.T) {
	test := map[string]struct {
		res    result
		myMove move
		want   int
	}{
		"Round 1": {
			res:    resultWinMe,
			myMove: movePaper,
			want:   8,
		},
		"Round 2": {
			res:    resultWinOpponent,
			myMove: moveRock,
			want:   1,
		},
		"Round 3": {
			res:    resultDraw,
			myMove: moveScissors,
			want:   6,
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := calculateRoundScore(tc.res, tc.myMove)
			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPlayRound(t *testing.T) {
	test := map[string]struct {
		want         result
		myMove       move
		opponentMove move
	}{
		"Round 1": {
			want:         resultWinMe,
			myMove:       moveScissors,
			opponentMove: movePaper,
		},
		"Round 2": {
			want:         resultWinOpponent,
			myMove:       moveScissors,
			opponentMove: moveRock,
		},
		"Round 3": {
			want:         resultDraw,
			myMove:       moveScissors,
			opponentMove: moveScissors,
		},
		"Round 4": {
			want:         resultDraw,
			myMove:       movePaper,
			opponentMove: movePaper,
		},
		"Round 5": {
			want:         resultWinMe,
			myMove:       movePaper,
			opponentMove: moveRock,
		},
		"Round 6": {
			want:         resultWinOpponent,
			myMove:       movePaper,
			opponentMove: moveScissors,
		},
		"Round 7": {
			want:         resultWinOpponent,
			myMove:       moveRock,
			opponentMove: movePaper,
		},
		"Round 8": {
			want:         resultDraw,
			myMove:       moveRock,
			opponentMove: moveRock,
		},
		"Round 9": {
			want:         resultWinMe,
			myMove:       moveRock,
			opponentMove: moveScissors,
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := playRound(tc.opponentMove, tc.myMove)
			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestFakeResult(t *testing.T) {
	test := map[string]struct {
		expectedResult result
		want           move
		opponentMove   move
	}{
		"Round 1": {
			expectedResult: resultWinMe,
			want:           moveScissors,
			opponentMove:   movePaper,
		},
		"Round 2": {
			expectedResult: resultWinOpponent,
			want:           moveScissors,
			opponentMove:   moveRock,
		},
		"Round 3": {
			expectedResult: resultDraw,
			want:           moveScissors,
			opponentMove:   moveScissors,
		},
		"Round 4": {
			expectedResult: resultDraw,
			want:           movePaper,
			opponentMove:   movePaper,
		},
		"Round 5": {
			expectedResult: resultWinMe,
			want:           movePaper,
			opponentMove:   moveRock,
		},
		"Round 6": {
			expectedResult: resultWinOpponent,
			want:           movePaper,
			opponentMove:   moveScissors,
		},
		"Round 7": {
			expectedResult: resultWinOpponent,
			want:           moveRock,
			opponentMove:   movePaper,
		},
		"Round 8": {
			expectedResult: resultDraw,
			want:           moveRock,
			opponentMove:   moveRock,
		},
		"Round 9": {
			expectedResult: resultWinMe,
			want:           moveRock,
			opponentMove:   moveScissors,
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := fakeResult(tc.expectedResult, tc.opponentMove)
			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}
