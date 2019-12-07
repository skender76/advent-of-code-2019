package main

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file(fileName string) []string {
	file, err := os.Open("input")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}


