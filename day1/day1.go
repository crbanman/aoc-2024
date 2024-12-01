package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetDistance(input string) int {
	var left []int
	var right []int
	numberStrings := strings.Fields(input)

	for index, numString := range numberStrings {
		numInt, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatalf("getdistance: couldn't parse %s into number", numString)
		}
		if index&1 != 1 {
			left = append(left, numInt)
		} else {
			right = append(right, numInt)
		}
	}

	if len(left) != len(right) {
		log.Fatalf("getdistance: the len of left (%d) didn't match the len of right (%d)", len(left), len(right))
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff > 0 {
			totalDistance += diff
		} else {
			totalDistance += -diff
		}
	}

	return totalDistance
}

func GetSimilarity(input string) int {
	left := map[int]int{}
	right := map[int]int{}

	numberStrings := strings.Fields(input)

	for index, numString := range numberStrings {
		numInt, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatalf("getsimilarity: couldn't parse %s into number", numString)
		}
		if index&1 != 1 {
			left[numInt] += 1
		} else {
			right[numInt] += 1
		}
	}

	similarity := 0
	for k, rv := range left {
		if lv, ok := right[k]; ok {
			similarity += k * rv * lv
		}
	}

	return similarity
}
