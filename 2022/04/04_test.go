package main

import "testing"

func TestFindOverlap(t *testing.T) {
	test := map[string]struct {
		want  bool
		input string
	}{
		"TC 1": {
			want:  false,
			input: "2-4,6-8",
		},
		"TC 2": {
			want:  false,
			input: "2-3,4-5",
		},
		"TC 3": {
			want:  true,
			input: "5-7,7-9",
		},
		"TC 4": {
			want:  true,
			input: "2-8,3-7",
		},
		"TC 5": {
			want:  true,
			input: "6-6,4-6",
		},
		"TC 6": {
			want:  true,
			input: "2-6,4-8",
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := findOverlap(tc.input)
			if got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestFindOverlapComplete(t *testing.T) {
	test := map[string]struct {
		want  bool
		input string
	}{
		"TC 1": {
			want:  false,
			input: "2-4,6-8",
		},
		"TC 2": {
			want:  false,
			input: "2-3,4-5",
		},
		"TC 3": {
			want:  false,
			input: "5-7,7-9",
		},
		"TC 4": {
			want:  true,
			input: "2-8,3-7",
		},
		"TC 5": {
			want:  true,
			input: "6-6,4-6",
		},
		"TC 6": {
			want:  false,
			input: "2-6,4-8",
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := findOverlapComplete(tc.input)
			if got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestParseRange(t *testing.T) {
	test := map[string]struct {
		want  string
		input string
	}{
		"TC 1": {
			want:  "|2|3|4|",
			input: "2-4",
		},
		"TC 2": {
			want:  "|6|7|8|",
			input: "6-8",
		},
		"TC 3": {
			want:  "|2|3|",
			input: "2-3",
		},
		"TC 4": {
			want:  "|4|5|",
			input: "4-5",
		},
		"TC 5": {
			want:  "|5|6|7|",
			input: "5-7",
		},
		"TC 6": {
			want:  "|7|8|9|",
			input: "7-9",
		},
		"TC 7": {
			want:  "|2|3|4|5|6|7|8|",
			input: "2-8",
		},
		"TC 8": {
			want:  "|3|4|5|6|7|",
			input: "3-7",
		},
		"TC 9": {
			want:  "|6|",
			input: "6-6",
		},
		"TC 10": {
			want:  "|4|5|6|",
			input: "4-6",
		},
		"TC 11": {
			want:  "|2|3|4|5|6|",
			input: "2-6",
		},
		"TC 12": {
			want:  "|4|5|6|7|8|",
			input: "4-8",
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := parseRange(tc.input)
			if got != tc.want {
				t.Fatalf("got %s, want %s", got, tc.want)
			}
		})
	}
}
