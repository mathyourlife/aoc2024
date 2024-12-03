package main

import (
	"bufio"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay03(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		solver Solver
		input  *bufio.Scanner
		want   int
	}{
		{
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
