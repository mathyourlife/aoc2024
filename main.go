package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Day01PartTwo()
}

func Day01PartTwo() {
	f, err := os.OpenFile("day01.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var a []int
	b := map[int]int{}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
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
	fmt.Println(a, b)

	if err := scan.Err(); err != nil {
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
	f, err := os.OpenFile("day01.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var a, b []int

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			log.Fatal("invalid line: ", line)
		}
		fmt.Println(parts[0], parts[1])
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

	if err := scan.Err(); err != nil {
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
