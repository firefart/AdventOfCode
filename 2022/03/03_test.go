package main

import "testing"

func TestFindDifference(t *testing.T) {
	test := map[string]struct {
		want     rune
		rucksack string
	}{
		"Rucksack 1": {
			want:     'p',
			rucksack: "vJrwpWtwJgWrhcsFMMfFFhFp",
		},
		"Rucksack 2": {
			want:     'L',
			rucksack: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		},
		"Rucksack 3": {
			want:     'P',
			rucksack: "PmmdzqPrVvPwwTWBwg",
		},
		"Rucksack 4": {
			want:     'v',
			rucksack: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		},
		"Rucksack 5": {
			want:     't',
			rucksack: "ttgJtRGJQctTZtZT",
		},
		"Rucksack 6": {
			want:     's',
			rucksack: "CrZsJsPPZsGzwwsLwLmpwMDw",
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := findDifference(tc.rucksack)
			if got != tc.want {
				t.Fatalf("got %c, want %c", got, tc.want)
			}
		})
	}
}

func TestCalculatePriority(t *testing.T) {
	test := map[string]struct {
		in   rune
		want int
	}{
		"Rucksack 1": {
			in:   'p',
			want: 16,
		},
		"Rucksack 2": {
			in:   'L',
			want: 38,
		},
		"Rucksack 3": {
			in:   'P',
			want: 42,
		},
		"Rucksack 4": {
			in:   'v',
			want: 22,
		},
		"Rucksack 5": {
			in:   't',
			want: 20,
		},
		"Rucksack 6": {
			in:   's',
			want: 19,
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := calculatePriority(tc.in)
			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestFindBadge(t *testing.T) {
	test := map[string]struct {
		want      rune
		rucksacks []string
	}{
		"Group 1": {
			want:      'r',
			rucksacks: []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
		},
		"Group 2": {
			want:      'Z',
			rucksacks: []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
		},
	}
	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := findBadge(tc.rucksacks)
			if got != tc.want {
				t.Fatalf("got %c, want %c", got, tc.want)
			}
		})
	}
}
