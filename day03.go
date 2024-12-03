package main

import (
	"bufio"
	"regexp"
)

func Day03Part2(scanner *bufio.Scanner) int {
	type operation struct {
		a, b int
	}

	reOp := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	enabled := true
	var lines [][]*operation
	lines = parse(scanner, func(line string) []*operation {
		var row []*operation
		for _, op := range reOp.FindAllStringSubmatch(line, -1) {
			switch op[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					row = append(row, &operation{mustInt(op[2]), mustInt(op[3])})
				}
			}
		}
		return row
	})

	var total int
	for _, line := range lines {
		for _, op := range line {
			total += op.a * op.b
		}
	}

	return total
}

func Day03Part1(scanner *bufio.Scanner) int {
	type operation struct {
		a, b int
	}

	reOp := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var lines [][]*operation
	lines = parse(scanner, func(line string) []*operation {
		var row []*operation
		for _, op := range reOp.FindAllStringSubmatch(line, -1) {
			row = append(row, &operation{mustInt(op[1]), mustInt(op[2])})
		}
		return row
	})

	var total int
	for _, line := range lines {
		for _, op := range line {
			total += op.a * op.b
		}
	}

	return total
}
