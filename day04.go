package main

import (
	"bufio"
	"strings"
)

func Day04Part2(scanner *bufio.Scanner) int {

	var lines []string
	lines = parse(scanner, func(line string) string {
		return line
	})
	grid := &Grid{
		data:     strings.Join(lines, ""),
		rowLen:   len(lines[0]),
		rowCount: len(lines),
	}

	var answer int
	// don't check the edges
	for row := 1; row < grid.rowCount-1; row++ {
		for col := 1; col < grid.rowLen-1; col++ {
			if grid.CharAt(row, col) == "A" {
				// check here
				cross1 := grid.CharAt(row-1, col-1) + "A" + grid.CharAt(row+1, col+1)
				cross2 := grid.CharAt(row+1, col-1) + "A" + grid.CharAt(row-1, col+1)
				if (cross1 == "MAS" || cross1 == "SAM") &&
					(cross2 == "MAS" || cross2 == "SAM") {
					answer++
				}
			}
		}
	}

	return answer
}

type Grid struct {
	data     string
	rowLen   int
	rowCount int
}

func (g *Grid) String() string {
	return g.data
}

func (g *Grid) CharAt(row, col int) string {
	return string(g.data[row*g.rowLen+col])
}

func Day04Part1(scanner *bufio.Scanner) int {
	var lines []string
	lines = parse(scanner, func(line string) string {
		return line
	})
	grid := &Grid{
		data:     strings.Join(lines, ""),
		rowLen:   len(lines[0]),
		rowCount: len(lines),
	}

	var answer int

	tmpGrid := strings.Join(lines, "\n")
	answer += strings.Count(tmpGrid, "XMAS")
	answer += strings.Count(tmpGrid, "SAMX")

	tmpGrid = ""
	for startRow := grid.rowCount - 1; startRow >= 0; startRow-- {
		for i := 0; i < grid.rowCount; i++ {
			row := (startRow - i + grid.rowCount) % grid.rowCount
			col := i
			tmpGrid += grid.CharAt(row, col)
			// fmt.Print(grid.CharAt(row, col))
			if row == 0 && i != grid.rowCount-1 {
				tmpGrid += "\n"
				// fmt.Println()
			}
		}
		tmpGrid += "\n"
		// fmt.Println()
	}
	answer += strings.Count(tmpGrid, "XMAS")
	answer += strings.Count(tmpGrid, "SAMX")

	tmpGrid = ""
	for startRow := 0; startRow < grid.rowCount; startRow++ {
		for i := 0; i < grid.rowCount; i++ {
			row := (startRow + i) % grid.rowCount
			col := i
			if row == 0 && startRow > 0 {
				tmpGrid += "\n"
				// fmt.Println()
			}
			tmpGrid += grid.CharAt(row, col)
			// fmt.Print(grid.CharAt(row, col))
		}
		tmpGrid += "\n"
		// fmt.Println()
	}
	answer += strings.Count(tmpGrid, "XMAS")
	answer += strings.Count(tmpGrid, "SAMX")

	// Vertical
	tmpGrid = ""
	for col := 0; col < grid.rowLen; col++ {
		for row := 0; row < grid.rowCount; row++ {
			tmpGrid += grid.CharAt(row, col)
		}
		tmpGrid += "\n"
	}
	answer += strings.Count(tmpGrid, "XMAS")
	answer += strings.Count(tmpGrid, "SAMX")

	return answer
}
