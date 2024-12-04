package main

import (
	"bufio"
	"sort"
	"strings"
)

// This time, you'll need to figure out exactly how often each number from the left list appears
// in the right list. Calculate a total similarity score by adding up each number in the left list
// after multiplying it by the number of times that number appears in the right list.
func Day01Part2(scanner *bufio.Scanner) int {
	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		parts := strings.Split(line, "   ")
		row := make([]int, 0, len(parts))
		for _, part := range parts {
			row = append(row, mustInt(part))
		}
		return row
	})

	a := make([]int, 0, len(lines))
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

// Pair up the smallest number in the left list with the smallest number in the right
// list, then the second-smallest left number with the second-smallest right number, and so on.
//
// Within each pair, figure out how far apart the two numbers are; you'll need to add up all of
// those distances. For example, if you pair up a 3 from the left list with a 7 from the right
// list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.
func Day01Part1(scanner *bufio.Scanner) int {
	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		parts := strings.Split(line, "   ")
		row := make([]int, 0, len(parts))
		for _, part := range parts {
			row = append(row, mustInt(part))
		}
		return row
	})

	a := make([]int, 0, len(lines))
	b := make([]int, 0, len(lines))

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
