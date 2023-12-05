package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var specialSymbols = []string{"$", "&", "*", "+", "%", "/", "-", "=", "@", "#"}
var specialSymbolsForMulti = []string{"*"}

func main() {

	lines, err := parseLines("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	grid := createGrid(lines)
	ratiosSum, err := calculateRatiosSum(grid)
	if err != nil {
		fmt.Print(err)
	}
	ratiosMulti, err := calculateRatiosMulti(grid)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Sum of first part: %d\n", ratiosSum)
	fmt.Printf("Sum of second part: %d\n", ratiosMulti)
}

func parseLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(readFile)), "\n")

	return lines, nil
}

func createGrid(lines []string) [][]string {
	var grid [][]string

	for _, row := range lines {

		rowSlice := strings.Split(row, "")

		grid = append(grid, rowSlice)
	}

	return grid
}

func calculateRatiosSum(grid [][]string) (int, error) {

	var result [][]string
	var checkSymbols [][]bool

	for i, rowValues := range grid {
		var symbols []bool
		var num []string
		for j, value := range rowValues {

			if _, err := strconv.Atoi(value); err == nil {

				symbols = append(symbols, checkAdjacentSymbol(grid, specialSymbols, i, j))
				num = append(num, value)
			}

			if value == "." || j == len(grid[i])-1 || isInSpecialSymbols(value) {
				if len(num) > 0 {
					result = append(result, num)
					checkSymbols = append(checkSymbols, symbols)
				}

				num = []string{}
				symbols = []bool{}
			}
		}
	}

	sum, err := filterArrays(result, checkSymbols)
	if err != nil {
		return 0, err
	}

	return sum, nil

}

func calculateRatiosMulti(grid [][]string) (int, error) {

	var result [][]string
	var checkSymbols [][]string

	for i, rowValues := range grid {
		var symbols []string
		var num []string
		for j, value := range rowValues {

			if _, err := strconv.Atoi(value); err == nil {

				if checkAdjacentSymbol(grid, specialSymbolsForMulti, i, j) {
					row, col := checkIndexOfSymbol(grid, specialSymbolsForMulti, i, j)
					position := fmt.Sprintf("%d %d", row, col)
					symbols = append(symbols, position)
				}
				num = append(num, value)

			}

			if value == "." || j == len(grid[i])-1 || isInSpecialSymbols(value) {
				if len(num) > 0 {
					result = append(result, num)
					checkSymbols = append(checkSymbols, symbols)
				}

				num = []string{}
				symbols = []string{}
			}

		}
	}

	sum, err := filterArraysBaseIndex(result, checkSymbols)
	if err != nil {
		return 0, err
	}

	return sum, nil

}

func checkAdjacentSymbol(grid [][]string, symbols []string, row, col int) bool {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}

			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
				for _, symbol := range symbols {
					if grid[i][j] == symbol {
						return true
					}
				}
			}
		}
	}
	return false
}

func checkIndexOfSymbol(grid [][]string, symbols []string, row, col int) (int, int) {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}

			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
				for _, symbol := range symbols {
					if grid[i][j] == symbol {
						return i, j
					}
				}
			}
		}
	}
	return 0, 0
}

func containsTrue(arr []bool) bool {
	for _, v := range arr {
		if v {
			return true
		}
	}
	return false
}

func isInSpecialSymbols(value string) bool {
	for _, symbol := range specialSymbols {
		if symbol == value {
			return true
		}
	}
	return false
}

func filterArrays(arr1 [][]string, boolArr [][]bool) (int, error) {
	var sum int

	for i, row := range arr1 {

		if containsTrue(boolArr[i]) {
			concatenated := strings.Join(row, "")
			num, err := strconv.Atoi(concatenated)
			if err != nil {
				return 0, err
			}
			sum += num
		}
	}
	return sum, nil
}

func filterArraysBaseIndex(arr1 [][]string, arr2 [][]string) (int, error) {
	var currentNumber []int
	var sum int
	for i := range arr1 {
		if i < len(arr2) && arr2[i] != nil {
			for j := range arr2 {
				if j < len(arr2) && arr2[j] != nil {
					if len(arr2[j]) > 0 && len(arr2[i]) > 0 {

						if j != i {
							if arr2[j][0] == arr2[i][0] {
								num1arr1 := strings.Join(arr1[i], "")
								num2arr2 := strings.Join(arr1[j], "")
								num1, _ := strconv.Atoi(num1arr1)
								num2, _ := strconv.Atoi(num2arr2)

								multi := num1 * num2

								sum += multi

								currentNumber = append(currentNumber, multi)
							}
						}

					}
				}

			}
		}
	}

	return sum / 2, nil
}

func uniqueRows(input []int) [][]string {
	seen := make(map[string]bool)
	var result [][]string

	for _, row := range input {
		rowString := fmt.Sprint(row)
		if _, exists := seen[rowString]; !exists {
			seen[rowString] = true
			result = append(result, []string{rowString})
		}
	}

	return result
}
