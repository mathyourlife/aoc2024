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
		log.Fatalf("error opening file '%s': %w", filename, err)
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

	reOp := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	var total int
	enabled := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		for _, op := range reOp.FindAllStringSubmatch(line, -1) {
			switch op[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					total += mustInt(op[2]) * mustInt(op[3])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}

func Day03PartOne() {
	scanner := DayScanner(3)
	// f := strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

	reOp := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var total int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		for _, op := range reOp.FindAllStringSubmatch(line, -1) {
			total += mustInt(op[1]) * mustInt(op[2])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
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

	var rows [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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

	isSafe := func(row ...int) bool {
		if isDescending(row...) {
			safe := true
			for i := 1; i < len(row); i++ {
				diff := row[i-1] - row[i]
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
		if isAscending(row...) {
			safe := true
			for i := 1; i < len(row); i++ {
				diff := row[i] - row[i-1]
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
	for _, row := range rows {
		if isSafe(row...) {
			total++
			continue
		}
		for i := 0; i < len(row); i++ {
			copy := slices.Clone(row)
			short := append(copy[:i], copy[i+1:]...)
			if isSafe(short...) {
				total++
				break
			}
		}
	}
	fmt.Println(total)
}

func Day02PartOne() {
	scanner := DayScanner(2)
	// f := strings.NewReader(`7 6 4 2 1
	// 1 2 7 8 9
	// 9 7 6 2 1
	// 1 3 2 4 5
	// 8 6 4 4 1
	// 1 3 6 7 9`)

	var rows [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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

	isSafe := func(row ...int) bool {
		if isDescending(row...) {
			safe := true
			for i := 1; i < len(row); i++ {
				diff := row[i-1] - row[i]
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
		if isAscending(row...) {
			safe := true
			for i := 1; i < len(row); i++ {
				diff := row[i] - row[i-1]
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
	for _, row := range rows {
		if isSafe(row...) {
			total++
		}
	}
	fmt.Println(total)
}

func Day01PartTwo() {
	scanner := DayScanner(1)

	var a []int
	b := map[int]int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			log.Fatal("invalid line: ", line)
		}
		aNum, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		bNum, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, aNum)
		b[bNum]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total int
	sort.Ints(a)
	for _, num := range a {
		total += int(num * b[num])
	}
	fmt.Println(total)
}

func Day01PartOne() {
	scanner := DayScanner(1)

	var a, b []int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			log.Fatal("invalid line: ", line)
		}

		aNum, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		bNum, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, aNum)
		b = append(b, bNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
}
