package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := getLines("./input.txt")

	if err != nil {
		fmt.Print(err)
	}

	// sum, err := getSum(data)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	sum, err := getSumWithWords(data)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Sum of numbers is: %d\n", sum)
}

func getLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(readFile), "\n")
	return lines, err
}

func getSum(data []string) (int, error) {
	var sum int
	for _, str := range data {
		var numbers []string
		var sumOfNumbers string

		for _, char := range str {
			if _, err := strconv.Atoi(string(char)); err == nil {
				numbers = append(numbers, string(char))
			}
		}
		sumOfNumbers = numbers[0] + numbers[len(numbers)-1]
		number, _ := strconv.Atoi(sumOfNumbers)
		sum += number
	}

	return sum, nil
}

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getSumWithWords(data []string) (int, error) {
	var sum int
	for _, str := range data {
		var numbers []int
		var sumOfNumbers string

		for i, char := range str {
			if num, err := strconv.Atoi(string(char)); err == nil {
				numbers = append(numbers, num)
			}

			for d, word := range words {
				if strings.HasPrefix(str[i:], word) {
					numbers = append(numbers, d+1)
				}
			}
		}
		sumOfNumbers = fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
		number, _ := strconv.Atoi(sumOfNumbers)
		sum += number
	}

	return sum, nil
}
