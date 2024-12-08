package day5

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type OrderData struct {
	after map[string]bool
}

func GetMiddlePageNumberSum(input string, onlyInvalidUpdates bool) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	orderMap := make(map[string]OrderData)
	invalidRows := [][]string{}
	isReadingOrder := true
	sum := 0
outer:
	for scanner.Scan() {
		line := scanner.Text()

		if isReadingOrder {
			split := strings.Split(line, "|")
			if len(line) == 0 || len(split) != 2 {
				isReadingOrder = false
				continue
			}

			left := split[0]
			right := split[1]
			if _, found := orderMap[left]; !found {
				orderMap[left] = OrderData{
					after: make(map[string]bool),
				}
			}
			if _, found := orderMap[right]; !found {
				orderMap[right] = OrderData{
					after: make(map[string]bool),
				}
			}
			orderMap[right].after[left] = true
			continue
		}

		pages := strings.Split(line, ",")
		for i := 0; i < len(pages)-1; i++ {
			orderData, found := orderMap[pages[i+1]]
			if !found {
				log.Fatalf("could not find value '%s' in orderMap", pages[i+1])
			}
			if !orderData.after[pages[i]] {
				invalidRows = append(invalidRows, pages)
				continue outer
			}
		}

		middleValue, err := strconv.Atoi(pages[len(pages)/2])
		if err != nil {
			log.Fatalf("could not convert value '%s' to an integer", pages[len(pages)/2])
		}

		if !onlyInvalidUpdates {
			sum += middleValue
		}
	}

	if onlyInvalidUpdates {
		for _, invalidRow := range invalidRows {
			for i, v := range invalidRow {
				afterCount := 0
				for j, c := range invalidRow {
					if i == j {
						continue
					}
					if orderMap[v].after[c] {
						afterCount++
					}
				}
				if afterCount == len(invalidRow)/2 {
					intValue, err := strconv.Atoi(v)
					if err != nil {
						log.Fatalf("could not convert value '%s' to an integer", v)
					}
					sum += intValue
				}
			}
		}
	}

	return sum
}
