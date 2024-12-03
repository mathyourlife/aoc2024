package main

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed data/*
var data embed.FS

func DayInputFile(day int) fs.File {
	filename := fmt.Sprintf("data/day%02d.txt", day)
	file, err := data.Open(filename)
	if err != nil {
		log.Fatalf("error opening file %q: %s", filename, err)
	}
	return file
}

func ExampleInputFile(day int) fs.File {
	filename := fmt.Sprintf("data/example%02d.txt", day)
	file, err := data.Open(filename)
	if err != nil {
		log.Fatalf("error opening file %q: %s", filename, err)
	}
	return file
}

func DayScanner(day int) *bufio.Scanner {
	return bufio.NewScanner(DayInputFile(day))
}

func ExampleScanner(day int) *bufio.Scanner {
	return bufio.NewScanner(ExampleInputFile(day))
}
