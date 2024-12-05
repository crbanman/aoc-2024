package day3

import (
	"log"
	"strconv"
)

const (
	MULT = "mul("
	DO   = "do()"
	DONT = "don't()"
)

func ParseAndMult(input string, doDont bool) int {
	sum := 0
	do := true
	parsingLeft := true
	var left string
	var right string

	for i := 0; i < len(input); {
		if i+len(MULT)+4 >= len(input) {
			// no room for mul(x,y), do(), or don't()
			break
		}

		if doDont && input[i:i+len(DO)] == DO {
			do = true
			i += len(DO)
			continue
		} else if doDont && input[i:i+len(DONT)] == DONT {
			do = false
			i += len(DONT)
			continue
		} else if input[i:i+len(MULT)] != MULT {
			i++
			continue
		}

		if !do {
			i += len(MULT)
			continue
		}

		i += len(MULT)

		for j := i; j < len(input); j++ {
			if parsingLeft && input[j] == ',' {
				if len(left) == 0 {
					// mul by zero
					left = ""
					right = ""
					break
				}
				i = j + 1
				parsingLeft = false
				continue
			}

			if !parsingLeft && input[j] == ')' {
				if len(right) == 0 {
					left = ""
					right = ""
					break
				}
				parsingLeft = true
				leftInt, err := strconv.Atoi(left)
				if err != nil {
					log.Fatalf("parseandmult: could not convert %s into an int", left)
				}
				rightInt, err := strconv.Atoi(right)
				if err != nil {
					log.Fatalf("parseandmult: could not convert %s into an int", right)
				}
				// mult values
				sum += leftInt * rightInt
			}

			if input[j] < '0' || input[j] > '9' {
				right = ""
				left = ""
				parsingLeft = true
				break
			}

			if parsingLeft {
				left += string(input[j])
			} else {
				right += string(input[j])
			}
		}

	}

	return sum
}
