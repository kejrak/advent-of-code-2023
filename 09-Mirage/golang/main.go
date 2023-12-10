package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(readFile), "\n")
	return lines, err
}

func calculateSum(lastValues []int) int {
	sum := lastValues[0]
	add := 0
	i := 0

	for i < len(lastValues)-1 {
		add = sum + lastValues[i+1]
		sum = add

		i += 1
	}

	return sum
}

func checkZeros(values []int) bool {

	for _, value := range values {
		if value != 0 {
			return false
		}
	}
	return true
}

func iterateOverSlice(values []int, lastValue int, lastValues []int) []int {
	var result []int

	if checkZeros(values) {
		return lastValues
	}

	lastValues = append(lastValues, lastValue)

	i := 0

	for i < len(values)-1 {

		diff := values[i+1] - values[i]

		result = append(result, diff)
		i += 1
	}

	lastValue = result[len(result)-1]

	return iterateOverSlice(result, lastValue, lastValues)
}

func reverse(input []int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func main() {
	data, err := getLines("./input.txt")

	if err != nil {
		fmt.Print(err)
	}

	var PartOne int
	var PartTwo int
	for _, line := range data {
		var result1 []int
		var result2 []int

		values := strings.Split(line, " ")
		intValues := make([]int, len(values))

		for i, s := range values {
			intValues[i], _ = strconv.Atoi(s)
		}

		result1 = iterateOverSlice(intValues, intValues[len(intValues)-1], []int{})
		result2 = iterateOverSlice(reverse(intValues), reverse(intValues)[len(intValues)-1], []int{})
		result1Sum := calculateSum(result1)
		result2Sum := calculateSum(result2)

		PartOne += result1Sum
		PartTwo += result2Sum
	}

	fmt.Printf("Part One: %d\n", PartOne)
	fmt.Printf("Part Two: %d\n", PartTwo)

}
