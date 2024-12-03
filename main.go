package main

import (
	"bufio"
	"fmt"
	"log"
)

type Solver func(*bufio.Scanner) int

type Day struct {
	Part1 Solver
	Part2 Solver
}

func main() {
	days := []Day{
		{Day01Part1, Day01Part2},
		{Day02Part1, Day02Part2},
		{Day03Part1, Day03Part2},
	}

	inputs := []struct {
		Day     int
		Part    int
		Scanner *bufio.Scanner
		Want    int
	}{
		// Examples
		{1, 1, ExampleScanner(1), 11},
		{1, 2, ExampleScanner(1), 31},
		{2, 1, ExampleScanner(2), 2},
		{2, 2, ExampleScanner(2), 4},
		{3, 1, ExampleScanner(3), 161},
		{3, 2, ExampleScanner(3), 48},

		// Real
		{1, 1, DayScanner(1), 1830467},
		{1, 2, DayScanner(1), 26674158},
		{2, 1, DayScanner(2), 585},
		{2, 2, DayScanner(2), 626},
		{3, 1, DayScanner(3), 171183089},
		{3, 2, DayScanner(3), 63866497},
	}

	for _, input := range inputs {
		day := days[input.Day-1]
		var got int
		switch input.Part {
		case 1:
			got = day.Part1(input.Scanner)
		case 2:
			got = day.Part2(input.Scanner)
		}
		if got != input.Want {
			log.Fatalf("Day%02dPart1 got %d, want %d", input.Day, got, input.Want)
		}
		fmt.Printf("Day%02dPart1 ok\n", input.Day)
	}
}
