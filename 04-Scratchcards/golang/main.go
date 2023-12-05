package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type IndexCount struct {
	Index    int
	Winnings int
	Count    int
}

func getLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(readFile), "\n")
	return lines, err
}

func incrementNextCards(indexCounts *[]IndexCount) {

	for i, item := range *indexCounts {
		if item.Winnings > 0 {
			for j := i + 1; j <= i+item.Winnings; j++ {
				(*indexCounts)[j].Count += item.Count
			}
		}
	}
}

func splitLines(line string) ([]string, []string) {
	game := strings.Split(line, ":")
	numbers := strings.Split(game[1], "|")

	arr1 := strings.Fields(numbers[0])
	arr2 := strings.Fields(numbers[1])

	return arr1, arr2
}

func main() {
	lines, err := getLines("./input.txt")
	if err != nil {
		fmt.Print(err)
	}

	result := 0
	var totalCount int
	var indexCounts []IndexCount

	for i, line := range lines {

		arr1, arr2 := splitLines(line)
		sum, count := checkWinnings(arr1, arr2)

		indexCount := IndexCount{
			Index:    i + 1,
			Winnings: count,
			Count:    1,
		}

		indexCounts = append(indexCounts, indexCount)
		result += sum
	}

	incrementNextCards(&indexCounts)
	for _, item := range indexCounts {
		totalCount += item.Count
	}

	fmt.Printf("Part one: %v\n", result)
	fmt.Printf("Part two: %v\n", totalCount)
}

func checkWinnings(numbers, reference []string) (int, int) {
	var total int
	var count int

	for len(numbers) > 0 {
		isWinner := isWinning(numbers[0], reference)

		if isWinner {
			count++
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}

		numbers = numbers[1:]
	}

	return total, count
}

func isWinning(s string, reference []string) bool {
	return slices.Index(reference, s) != -1
}
