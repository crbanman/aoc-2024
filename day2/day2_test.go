package day2

import (
	"log"
	"testing"
)

const input = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
10 1 2 3 4
13 16 13 11 8 5
`

func TestGetSafeCountNoDampener(t *testing.T) {
	result := GetSafeCount(input, false)
	if result != 2 {
		log.Fatalf("invalid safe count without dampener. expected %d, received %d", 2, result)
	}
}

func TestGetSafeCountWithDampener(t *testing.T) {
	result := GetSafeCount(input, true)
	if result != 6 {
		log.Fatalf("invalid safe count with dampener. expected %d, received %d", 6, result)
	}
}
