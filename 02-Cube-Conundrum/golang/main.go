package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	regx = `(\d+) (\w+)`
)

func main() {
	lines, err := getLines("./input.txt")
	if err != nil {
		fmt.Print(err)
	}

	sumPart1, sumPart2 := getGames(lines)

	fmt.Printf("Part 1 sum: %d\n", sumPart1)
	fmt.Printf("Part 2 sum: %d\n", sumPart2)

}

func getLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(readFile)), "\n")
	return lines, err
}

func getGames(lines []string) (int, int) {
	var sumPart1 int
	var sumPart2 int

	pattern := regexp.MustCompile(regx)
	for num, line := range lines {

		var maxQuantities = make(map[string]int)

		for _, match := range pattern.FindAllStringSubmatch(line, -1) {
			quantity, _ := strconv.Atoi(match[1])
			maxQuantities[match[2]] = slices.Max([]int{maxQuantities[match[2]], quantity})
		}

		if maxQuantities["red"] <= 12 && maxQuantities["green"] <= 13 && maxQuantities["blue"] <= 14 {
			sumPart1 += num + 1
		}
		sumPart2 += maxQuantities["red"] * maxQuantities["green"] * maxQuantities["blue"]
	}
	return sumPart1, sumPart2
}
