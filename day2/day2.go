package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func GetSafeCount(input string, withDampener bool) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		values := strings.Fields(line)
		if validateReport(values, withDampener) {
			safeCount++
			continue
		}

	}

	return safeCount
}

func validateReport(report []string, withDampener bool) bool {
	if len(report) == 0 {
		return false
	}

	prev := 0
	for i, v := range report {
		// Don't evaluate last number as it was evaluated on previous iteration
		if i == len(report)-1 {
			break
		}

		cur, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("getsafecount: couldn't parse cur %s into number\n", v)
			return false
		}

		next, err := strconv.Atoi(report[i+1])
		if err != nil {
			fmt.Printf("getsafecount: couldn't parse next %s into number", v)
			return false
		}

		if cur == next || cur-next > 3 || next-cur > 3 || prev != 0 && prev > cur && cur < next || prev != 0 && prev < cur && cur > next {
			if withDampener {
				if i > 0 {
					withoutPrev, err := copySliceRemoveIndex(report, i-1)
					if err != nil {
						fmt.Printf("could not remove index %d from slice '%v'", i-1, report)
						return false
					}
					if validateReport(withoutPrev, false) {
						return true
					}
				}
				withoutCur, err := copySliceRemoveIndex(report, i)
				if err != nil {
					fmt.Printf("could not remove index %d from slice '%v'", i, report)
					return false
				}
				if validateReport(withoutCur, false) {
					return true
				}
				withoutNext, err := copySliceRemoveIndex(report, i+1)
				if err != nil {
					fmt.Printf("could not remove index %d from slice '%v'", i+1, report)
					return false
				}
				return validateReport(withoutNext, false)
			} else {
				return false
			}
		}
		prev = cur
	}

	return true
}

func copySliceRemoveIndex(slice []string, index int) ([]string, error) {
	if index < 0 || index > len(slice) {
		return slice, fmt.Errorf("copysliceremoveindex: invalid index %d for slice with length %d", index, len(slice))
	}

	newSlice := make([]string, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice, nil
}
