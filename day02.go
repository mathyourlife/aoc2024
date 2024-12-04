package main

import (
	"bufio"
	"slices"
	"sort"
	"strings"
)

func Day02Part2(scanner *bufio.Scanner) int {
	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		parts := strings.Split(line, " ")
		row := make([]int, 0, len(parts))
		for _, part := range parts {
			row = append(row, mustInt(part))
		}
		return row
	})

	isAscending := func(line ...int) bool {
		cmp := slices.Clone(line)
		sort.Ints(cmp)
		if slices.Equal(cmp, line) {
			return true
		}
		return false
	}
	isDescending := func(line ...int) bool {
		cmp := slices.Clone(line)
		sort.Sort(sort.Reverse(sort.IntSlice(cmp)))
		if slices.Equal(cmp, line) {
			return true
		}
		return false
	}

	isSafe := func(line ...int) bool {
		if isDescending(line...) {
			safe := true
			for i := 1; i < len(line); i++ {
				diff := line[i-1] - line[i]
				switch {
				case diff == 0:
					safe = false
				case diff >= 4:
					safe = false
				}
			}
			if safe {
				return true
			}
		}
		if isAscending(line...) {
			safe := true
			for i := 1; i < len(line); i++ {
				diff := line[i] - line[i-1]
				switch {
				case diff == 0:
					safe = false
				case diff >= 4:
					safe = false
				}
			}
			if safe {
				return true
			}
		}
		return false
	}

	var total int
	for _, line := range lines {
		if isSafe(line...) {
			total++
			continue
		}
		for i := 0; i < len(line); i++ {
			copy := slices.Clone(line)
			short := append(copy[:i], copy[i+1:]...)
			if isSafe(short...) {
				total++
				break
			}
		}
	}
	return total
}

func Day02Part1(scanner *bufio.Scanner) int {
	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		parts := strings.Split(line, " ")
		row := make([]int, 0, len(parts))
		for _, part := range parts {
			row = append(row, mustInt(part))
		}
		return row
	})

	isAscending := func(num ...int) bool {
		cmp := slices.Clone(num)
		sort.Ints(cmp)
		if slices.Equal(cmp, num) {
			return true
		}
		return false
	}
	isDescending := func(num ...int) bool {
		cmp := slices.Clone(num)
		sort.Sort(sort.Reverse(sort.IntSlice(cmp)))
		if slices.Equal(cmp, num) {
			return true
		}
		return false
	}

	isSafe := func(line ...int) bool {
		if isDescending(line...) {
			safe := true
			for i := 1; i < len(line); i++ {
				diff := line[i-1] - line[i]
				switch {
				case diff == 0:
					safe = false
				case diff >= 4:
					safe = false
				}
			}
			if safe {
				return true
			}
		}
		if isAscending(line...) {
			safe := true
			for i := 1; i < len(line); i++ {
				diff := line[i] - line[i-1]
				switch {
				case diff == 0:
					safe = false
				case diff >= 4:
					safe = false
				}
			}
			if safe {
				return true
			}
		}
		return false
	}

	var total int
	for _, line := range lines {
		if isSafe(line...) {
			total++
		}
	}
	return total
}
