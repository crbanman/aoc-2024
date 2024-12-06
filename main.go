package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/crbanman/aoc-2024/day1"
	"github.com/crbanman/aoc-2024/day2"
	"github.com/crbanman/aoc-2024/day3"
	"github.com/crbanman/aoc-2024/day4"
)

func main() {
	inputDay := flag.String("day", "day1", "The day to run: -day day1")
	flag.Parse()
	data, err := os.ReadFile("./" + *inputDay + "/input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	input := string(data)

	switch *inputDay {
	case "day1":
		runDay1(input)
	case "day2":
		runDay2(input)
	case "day3":
		runDay3(input)
	case "day4":
		runDay4(input)
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

func runDay2(input string) {
	p1 := day2.GetSafeCount(input, false)
	fmt.Printf("Part 1: %d\n", p1)

	p2 := day2.GetSafeCount(input, true)
	fmt.Printf("Part 2: %d\n", p2)
}

func runDay3(input string) {
	p1 := day3.ParseAndMult(input, false)
	fmt.Printf("Part 1: %d\n", p1)

	p2 := day3.ParseAndMult(input, true)
	fmt.Printf("Part 2: %d\n", p2)
}

func runDay4(input string) {
	p1 := day4.SearchXMas(input)
	fmt.Printf("Part 1: %d\n", p1)

	p2 := day4.SearchMasX(input)
	fmt.Printf("Part 2: %d\n", p2)
}
