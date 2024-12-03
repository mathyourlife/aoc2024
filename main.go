package main

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//go:embed data/*
var data embed.FS

func DayInput(day int) fs.File {
	filename := fmt.Sprintf("data/day%02d.txt", day)
	file, err := data.Open(filename)
	if err != nil {
		log.Fatalf("error opening file '%s': %s", filename, err)
	}
	return file
}

func DayScanner(day int) *bufio.Scanner {
	return bufio.NewScanner(DayInput(day))
}

func main() {
	Day01PartOne()
	Day01PartTwo()
	Day02PartOne()
	Day02PartTwo()
	Day03PartOne()
	Day03PartTwo()
}

func Day03PartTwo() {
	scanner := DayScanner(3)
	// f := strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

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

	fmt.Println(total)
	if total != 63866497 {
		log.Fatal("total != 63866497")
	}
}

func Day03PartOne() {
	scanner := DayScanner(3)
	// f := strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

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

	fmt.Println(total)
	if total != 171183089 {
		log.Fatal("total != 171183089")
	}
}

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Day02PartTwo() {
	scanner := DayScanner(2)
	// f := strings.NewReader(`7 6 4 2 1
	// 1 2 7 8 9
	// 9 7 6 2 1
	// 1 3 2 4 5
	// 8 6 4 4 1
	// 1 3 6 7 9`)

	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		var row []int
		for _, part := range strings.Split(line, " ") {
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
	fmt.Println(total)
	if total != 626 {
		log.Fatal("total != 626")
	}
}

func Day02PartOne() {
	scanner := DayScanner(2)
	// f := strings.NewReader(`7 6 4 2 1
	// 1 2 7 8 9
	// 9 7 6 2 1
	// 1 3 2 4 5
	// 8 6 4 4 1
	// 1 3 6 7 9`)

	var lines [][]int
	lines = parse(scanner, func(line string) []int {
		var row []int
		for _, part := range strings.Split(line, " ") {
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
	fmt.Println(total)
	if total != 585 {
		log.Fatal("total != 585")
	}
}

func Day01PartTwo() {
	scanner := DayScanner(1)

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
	fmt.Println(total)
	if total != 26674158 {
		log.Fatal("total != 26674158")
	}
}

func Day01PartOne() {
	scanner := DayScanner(1)

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
	fmt.Println(total)
	if total != 1830467 {
		log.Fatal("total != 1830467")
	}
}

func parse[T any](scanner *bufio.Scanner, f func(line string) T) []T {
	var parsedLines []T
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var parsedLine T
		parsedLine = f(line)
		parsedLines = append(parsedLines, parsedLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return parsedLines
}
