package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path := filepath.Join("day_2", "input.txt")
	allInput, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}

	input := strings.Split(string(allInput), ",")

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))

}

func part1(input []string) int {
	invalidIdCount := 0
	for i := 0; i < len(input); i++ {
		idRange := strings.Split(input[i], "-")
		if len(idRange) < 2 {
			log.Printf("No range found")
			continue
		}
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		countFromRange := checkValidRange(start, end)
		invalidIdCount = invalidIdCount + countFromRange
	}
	return invalidIdCount
}

func part2(input []string) int {
	invalidIdCount := 0
	for i := 0; i < len(input); i++ {
		idRange := strings.Split(input[i], "-")
		if len(idRange) < 2 {
			log.Printf("No range found")
			continue
		}
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		countFromRange := checkValidRangePart2(start, end)
		invalidIdCount = invalidIdCount + countFromRange
	}
	return invalidIdCount
}

func checkValidRangePart2(start int, end int) int {
	invalidCount := 0

	for i := start; i <= end; i++ {
		str := strconv.Itoa(i)
		splitStr := strings.Split(str, "")
		for j := 0; j < len(splitStr); j++ {
			if checkForRepeats(strings.Join(splitStr[:j+1], ""), str) {
				invalidCount += i
				break
			}
		}

	}

	return invalidCount
}

func checkForRepeats(sub string, input string) bool {
	for i := 1; i < len(input)+1; i++ {
		if strings.Repeat(sub, i) == input && sub != input {
			return true
		}
	}
	return false
}

func checkValidRange(start int, end int) int {
	invalidCount := 0

	for i := start; i <= end; i++ {
		str := strconv.Itoa(i)
		splitStr := strings.Split(str, "")
		for j := 0; j < len(splitStr); j++ {
			if strings.Join(splitStr[j:], "") == strings.Join(splitStr[:j], "") {
				invalidCount += i
			}
		}
	}

	return invalidCount
}
