package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	A int = 14
	K int = 13
	Q int = 12
	J int = 11
	T int = 10
)

type HandCards struct {
	Cards    string
	Bid      int
	Strength int
}

func (h *HandCards) GetType() {

	occurrences := countCardOccurrences(h.Cards)

	isFiveStrength, isFive := h.isFive(occurrences)
	isFourStrength, isFour := h.isFour(occurrences)
	isFullHouseStrength, isFullHouse := h.isFullHouse(occurrences)
	isThreeStrength, isThree := h.isThree(occurrences)
	isTwoStrength, isTwo := h.isTwo(occurrences)
	isOneStrength, isOne := h.isOne(occurrences)
	isHighStrength, isHigh := h.isHigh(occurrences)

	if isFive {
		h.Strength = isFiveStrength
	}

	if isFour {
		h.Strength = isFourStrength
	}

	if isFullHouse {
		h.Strength = isFullHouseStrength
	}

	if isThree {
		h.Strength = isThreeStrength
	}

	if isTwo {
		h.Strength = isTwoStrength
	}

	if isOne {
		h.Strength = isOneStrength
	}

	if isHigh {
		h.Strength = isHighStrength
	}

}

func (h *HandCards) isFive(occurrences map[rune]int) (int, bool) {

	for _, count := range occurrences {

		if count == 5 {

			return 7, true

		}
	}
	return 0, false

}

func (h *HandCards) isFour(occurrences map[rune]int) (int, bool) {

	for _, count := range occurrences {

		if count == 4 {

			return 6, true

		}
	}
	return 0, false
}

func (h *HandCards) isFullHouse(occurrences map[rune]int) (int, bool) {

	threeCount := 0
	twoCount := 0

	for _, count := range occurrences {

		switch count {
		case 3:
			threeCount++
		case 2:
			twoCount++
		}
	}
	return 5, threeCount == 1 && twoCount == 1
}

func (h *HandCards) isThree(occurrences map[rune]int) (int, bool) {

	threeCount := 0
	oneCount := 0

	for _, count := range occurrences {

		switch count {
		case 3:
			threeCount++
		case 1:
			oneCount++
		}
	}

	return 4, threeCount == 1 && oneCount == 2

}

func (h *HandCards) isTwo(occurrences map[rune]int) (int, bool) {

	twoCount := 0

	for _, count := range occurrences {

		if count == 2 {
			twoCount++
		}
	}

	return 3, twoCount == 2

}

func (h *HandCards) isOne(occurrences map[rune]int) (int, bool) {

	twoCount := 0
	oneCount := 0

	for _, count := range occurrences {

		switch count {
		case 1:
			oneCount++
		case 2:
			twoCount++
		}
	}

	return 2, twoCount == 1 && oneCount == 3

}

func (h *HandCards) isHigh(occurrences map[rune]int) (int, bool) {

	oneCount := 0

	for _, count := range occurrences {

		if count == 1 {
			oneCount++
		}

	}
	return 1, oneCount == 5
}

func countCardOccurrences(s string) map[rune]int {

	cardCount := make(map[rune]int)

	for _, card := range s {

		cardCount[card]++
	}

	return cardCount
}

func getLines(filePath string) ([]string, error) {
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(readFile), "\n")
	return lines, err
}

func getValue(char byte) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}

	switch char {
	case 'A':
		return A
	case 'K':
		return K
	case 'Q':
		return Q
	case 'J':
		return J
	case 'T':
		return T
	default:
		return 0
	}

}

func compareStrings(s1, s2 string) bool {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		val1 := getValue(s1[i])
		val2 := getValue(s2[i])

		if val1 > val2 {
			return true
		} else if val1 < val2 {
			return false
		}
	}

	return len(s1) < len(s2)
}

func main() {

	data, err := getLines("./input.txt")

	if err != nil {
		fmt.Print(err)
	}

	var hands []HandCards

	for _, hand := range data {
		hand := strings.Split(hand, " ")

		cards := hand[0]
		bid, err := strconv.Atoi(hand[1])

		if err != nil {
			fmt.Print(err)
		}

		handCards := HandCards{
			Cards:    cards,
			Bid:      bid,
			Strength: 0,
		}

		hands = append(hands, handCards)
	}

	for i := range hands {
		hands[i].GetType()
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Strength > hands[j].Strength
	})

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Strength != hands[j].Strength {
			return hands[i].Strength > hands[j].Strength
		}
		return compareStrings(hands[i].Cards, hands[j].Cards)
	})

	for i, j := 0, len(hands)-1; i < j; i, j = i+1, j-1 {
		hands[i], hands[j] = hands[j], hands[i]
	}

	totalBid := 0

	for i, hand := range hands {

		count := hand.Bid * (i + 1)

		totalBid += count

	}

	fmt.Printf("First part %d\n", totalBid)

}
