package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func mustInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
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
