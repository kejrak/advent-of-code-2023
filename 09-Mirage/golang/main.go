package main

import (
	"fmt"
	"os"
	"slices"
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

func GetSum(lastValues []int) int {
	sum := 0
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

func IterateOverSlice(values []int, lastValue int, lastValues []int) []int {
	var result []int

	if checkZeros(values) {
		return lastValues
	}

	i := 0

	for i < len(values)-1 {

		diff := values[i+1] - values[i]

		result = append(result, diff)
		i += 1
	}

	lastValue = result[len(result)-1]
	lastValues = append(lastValues, lastValue)

	return IterateOverSlice(result, lastValue, lastValues)
}

func PartOne(data []string) int {

	var PartOne int
	for _, line := range data {
		var listOfLastValues []int

		values := strings.Split(line, " ")
		intValues := make([]int, len(values))

		for i, s := range values {
			intValues[i], _ = strconv.Atoi(s)
		}

		listOfLastValues = IterateOverSlice(intValues, intValues[len(intValues)-1], []int{})
		slices.Reverse(listOfLastValues)
		listOfLastValues = append(listOfLastValues, intValues[len(intValues)-1])
		getSum := GetSum(listOfLastValues)

		PartOne += getSum
	}

	return PartOne

}

func PartTwo(data []string) int {

	var PartTwo int
	for _, line := range data {
		var listOfLastValues []int

		values := strings.Split(line, " ")
		intValues := make([]int, len(values))

		for i, s := range values {
			intValues[i], _ = strconv.Atoi(s)
		}

		slices.Reverse(intValues)
		listOfLastValues = IterateOverSlice(intValues, intValues[len(intValues)-1], []int{})
		slices.Reverse(listOfLastValues)

		listOfLastValues = append(listOfLastValues, intValues[len(intValues)-1])

		getSum := GetSum(listOfLastValues)

		PartTwo += getSum
	}

	return PartTwo

}

func main() {
	data, err := getLines("./input.txt")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Part Two: %d\n", PartOne(data))
	fmt.Printf("Part Two: %d\n", PartTwo(data))

}
