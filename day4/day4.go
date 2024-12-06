package day4

import (
	"errors"
	"fmt"
	"log"
)

func SearchXMas(input string) int {
	count := 0

	cols, err := getNumCols(input)
	if err != nil {
		log.Fatal("search: could not determine columns from input")
	}

	rows := len(input) / (cols + 1)

	xmasLen := len("XMAS")
	fmt.Printf("rows: %d; columns %d; len %d\n", rows, cols, len(input))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			upRightValues := []byte{}
			rightValues := []byte{}
			downRightValues := []byte{}
			downValues := []byte{}
			for i := 0; i < xmasLen; i++ {
				if cols-x >= xmasLen {
					if y+1 >= xmasLen {
						upRightValues = append(upRightValues, input[coordToIndex(x+i, y-i, cols)])
					}
					rightValues = append(rightValues, input[coordToIndex(x+i, y, cols)])
				}
				if rows-y >= xmasLen {
					if cols-x >= xmasLen {
						downRightValues = append(downRightValues, input[coordToIndex(x+i, y+i, cols)])
					}
					downValues = append(downValues, input[coordToIndex(x, y+i, cols)])
				}
			}

			upRightString := string(upRightValues)
			if upRightString == "XMAS" || upRightString == "SAMX" {
				count++
			}
			rightString := string(rightValues)
			if rightString == "XMAS" || rightString == "SAMX" {
				count++
			}
			downRightString := string(downRightValues)
			if downRightString == "XMAS" || downRightString == "SAMX" {
				count++
			}
			downString := string(downValues)
			if downString == "XMAS" || downString == "SAMX" {
				count++
			}
		}
	}

	return count
}

func SearchMasX(input string) int {
	count := 0

	cols, err := getNumCols(input)
	if err != nil {
		log.Fatal("search: could not determine columns from input")
	}

	rows := len(input) / (cols + 1)

	for y := 1; y < rows-1; y++ {
		for x := 1; x < cols-1; x++ {
			if input[coordToIndex(x, y, cols)] != 'A' {
				continue
			}
			topLeft := input[coordToIndex(x-1, y-1, cols)]
			topRight := input[coordToIndex(x+1, y-1, cols)]
			bottomLeft := input[coordToIndex(x-1, y+1, cols)]
			bottomRight := input[coordToIndex(x+1, y+1, cols)]

			if (topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M') &&
				(bottomLeft == 'M' && topRight == 'S' || bottomLeft == 'S' && topRight == 'M') {
				count++
			}
		}
	}

	return count
}

func coordToIndex(x, y, cols int) int {
	// Add y at the end to account for the \n
	return cols*y + x + y
}

func getNumCols(input string) (int, error) {
	for i, r := range input {
		if r == '\n' {
			return i, nil
		}
	}

	return 0, errors.New("getnumcols: could not find '\\n' in input")
}
