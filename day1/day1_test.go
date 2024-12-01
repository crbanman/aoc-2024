package day1

import "testing"

const input = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestGetDistance(t *testing.T) {
	result := GetDistance(input)

	if 11 != result {
		t.Fatalf("got %d, expected %d", result, 11)
	}
}

func TestGetSimilarity(t *testing.T) {
	result := GetSimilarity(input)

	if 31 != result {
		t.Fatalf("got %d, expected %d", result, 31)
	}
}
