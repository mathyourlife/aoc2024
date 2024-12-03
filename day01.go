package main

import (
	"bufio"
	"sort"
	"strings"
)

func Day01Part2(scanner *bufio.Scanner) int {
	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		var row []int
		for _, part := range strings.Split(line, "   ") {
			row = append(row, mustInt(part))
		}
		return row
	})

	var a []int
	b := map[int]int{}

	for _, line := range lines {
		a = append(a, line[0])
		b[line[1]]++
	}

	var total int
	sort.Ints(a)
	for _, num := range a {
		total += int(num * b[num])
	}
	return total
}

func Day01Part1(scanner *bufio.Scanner) int {
	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		var row []int
		for _, part := range strings.Split(line, "   ") {
			row = append(row, mustInt(part))
		}
		return row
	})

	var a, b []int
	for _, line := range lines {
		a = append(a, line[0])
		b = append(b, line[1])
	}

	sort.Ints(a)
	sort.Ints(b)

	total := 0
	for i := 0; i < len(a); i++ {
		diff := b[i] - a[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	return total
}
