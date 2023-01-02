package main

import (
	"fmt"
	"testing"
)

func TestCharsUnique(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  false,
		},
		{
			input: "",
			want:  true,
		},
		{
			input: "aaaa",
			want:  false,
		},
		{
			input: "abcdefg",
			want:  true,
		},
		{
			input: "abcdefgg",
			want:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := charsUnique(tc.input)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestFindStartSequence(t *testing.T) {
	tests := []struct {
		input          string
		want           int
		startSeqLength int
	}{
		{
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			startSeqLength: 4,
			want:           7,
		},
		{
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			startSeqLength: 4,
			want:           5,
		},
		{
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			startSeqLength: 4,
			want:           6,
		},
		{
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			startSeqLength: 4,
			want:           10,
		},
		{
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			startSeqLength: 4,
			want:           11,
		},
		{
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			startSeqLength: 14,
			want:           19,
		},
		{
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			startSeqLength: 14,
			want:           23,
		},
		{
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			startSeqLength: 14,
			want:           23,
		},
		{
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			startSeqLength: 14,
			want:           29,
		},
		{
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			startSeqLength: 14,
			want:           26,
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s -> %d", tc.input, tc.want), func(t *testing.T) {
			got := findStartSequence(tc.input, tc.startSeqLength)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
