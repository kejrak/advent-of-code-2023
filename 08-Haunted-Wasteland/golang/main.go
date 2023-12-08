package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	LeftMove  string
	RightMove string
}

type Mapping map[string]Coordinates

var moveIndex = 0
var counter = 0

func GetSteps(data Mapping, position string, moves []string) int {

	if position == "ZZZ" {
		return counter
	}

	if moveIndex == len(moves) {
		moveIndex = 0
	}

	if moves[moveIndex] == "R" {
		moveIndex += 1
		counter++
		nextMove := data[position].RightMove
		return GetSteps(data, nextMove, moves)

	} else {
		moveIndex += 1
		counter++
		nextMove := data[position].LeftMove
		return GetSteps(data, nextMove, moves)
	}

}

func GetStartNodes(data Mapping) []string {

	var listNodes []string

	for key := range data {
		if key[2] == 'A' {
			listNodes = append(listNodes, key)
		}
	}

	return listNodes
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GetEndingIndexes(data Mapping, position string, moves []string) int {

	i := 0

	moveIndex := 0

	currentPosition := position

	for i < 100000 {
		if moveIndex == len(moves) {
			moveIndex = 0
		}

		if currentPosition[len(currentPosition)-1] == 'Z' {
			return i
		}

		if moves[moveIndex] == "L" {
			currentPosition = data[currentPosition].LeftMove
		} else {
			currentPosition = data[currentPosition].RightMove
		}

		i += 1
		moveIndex += 1
	}

	return 0
}

func getLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(readFile), "\n")
	return lines, err
}

func main() {
	data, err := getLines("./input.txt")

	moves := strings.Split(data[0], "")

	if err != nil {
		fmt.Print(err)
	}

	coordinatesMap := make(map[string]Coordinates)

	paths := data[2:]

	for _, line := range paths {

		value := strings.Split(line, " = ")[0]
		moves := strings.Split(line, "=")[1]
		options := strings.Split(moves, ",")

		leftValue := options[0][2:]
		rightValue := options[1][1:4]

		coordinate := Coordinates{
			LeftMove:  leftValue,
			RightMove: rightValue,
		}

		coordinatesMap[value] = coordinate
	}

	listOfNodes := GetStartNodes(coordinatesMap)

	var result []int

	for _, node := range listOfNodes {
		endIndex := GetEndingIndexes(coordinatesMap, node, moves)
		result = append(result, endIndex)
	}

	fmt.Println(GetSteps(coordinatesMap, "AAA", moves))
	fmt.Println(LCM(result[0], result[1], result[1:]...))

}
