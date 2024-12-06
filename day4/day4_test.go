package day4

import "testing"

const input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestSearchXMas(t *testing.T) {
	result := SearchXMas(input)
	expected := 18
	if result != expected {
		t.Fatalf("expected %d receved %d", expected, result)
	}
}

func TestSearchMasX(t *testing.T) {
	result := SearchMasX(input)
	expected := 9
	if result != expected {
		t.Fatalf("expected %d receved %d", expected, result)
	}
}
