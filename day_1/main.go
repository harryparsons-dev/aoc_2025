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
	path := filepath.Join("day_1.txt")
	allInput, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error opening file: %v", err)
	}

	input := strings.Split(string(allInput), "\n")

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 1: %v\n", part2(input))

}

func part1(input []string) int {
	zeroCount := 0
	currentPtr := 50
	for _, command := range input {
		left := strings.HasPrefix(command, "L")
		distStr := command[1:]
		dist, _ := strconv.Atoi(distStr)

		if left == true { // going left
			currentPtr = (currentPtr - dist + 100) % 100

		} else { // going right
			currentPtr = (currentPtr + dist) % 100
		}
		if currentPtr == 0 {
			zeroCount += 1
		}

	}

	return zeroCount
}

func part2(input []string) int {
	zeroCount := 0
	currentPtr := 50
	for _, command := range input {
		left := strings.HasPrefix(command, "L")
		distStr := command[1:]
		dist, _ := strconv.Atoi(distStr)

		for i := 0; i < dist; i++ {
			if left == true { // going left
				currentPtr = (currentPtr - 1 + 100) % 100

			} else { // going right
				currentPtr = (currentPtr + 1) % 100
			}
			if currentPtr == 0 {
				zeroCount += 1
			}
		}

	}
	return zeroCount
}

// NOT 3156 too low
// 8733
// not 7941
// not 7101
