package day3

import (
	"testing"
)

func TestParseAndMultDoDontFalse(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@
^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	result := ParseAndMult(input, false)
	expected := 161
	if result != expected {
		t.Fatalf("expected %d receved %d", expected, result)
	}
}

func TestParseAndMultDoDontTrue(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)
+mul(32,64](mul(11,8)undo()?mul(8,5))`
	result := ParseAndMult(input, true)
	expected := 48
	if result != expected {
		t.Fatalf("expected %d receved %d", expected, result)
	}
}
