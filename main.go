package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/crbanman/aoc-2024/day1"
)

func main() {
	inputDay := flag.String("day", "day1", "The day to run: -day day1")
	filePath := flag.String("input", "./"+*inputDay+"/input.txt", "Path to the input file containing number pairs")
	flag.Parse()

	data, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	input := string(data)

	switch *inputDay {
	case "day1":
		runDay1(input)
	default:
		log.Fatalf("invalid day %s", *inputDay)
	}

}

func runDay1(input string) {
	p1 := day1.GetDistance(input)
	fmt.Printf("Part 1: %d\n", p1)

	p2 := day1.GetSimilarity(input)
	fmt.Printf("Part 2: %d\n", p2)
}
