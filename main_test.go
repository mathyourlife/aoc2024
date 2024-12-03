package main

import (
	"bufio"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDays(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		solver Solver
		input  *bufio.Scanner
		want   int
	}{
		{
			name:   "Day 01 Part 1",
			solver: Day01Part1,
			input:  ExampleScanner(1),
			want:   11,
		}, {
			name:   "Day 01 Part 2",
			solver: Day01Part2,
			input:  ExampleScanner(1),
			want:   31,
		}, {
			name:   "Day 02 Part 1",
			solver: Day02Part1,
			input:  ExampleScanner(2),
			want:   2,
		}, {
			name:   "Day 02 Part 2",
			solver: Day02Part2,
			input:  ExampleScanner(2),
			want:   4,
		}, {
			name:   "Day 03 Part 1",
			solver: Day03Part1,
			input:  ExampleScanner(3),
			want:   161,
		}, {
			name:   "Day 03 Part 2",
			solver: Day03Part2,
			input:  ExampleScanner(3),
			want:   48,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.solver(tc.input)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("mismatch (-got, +want):\n%v", diff)
			}
		})
	}
}
