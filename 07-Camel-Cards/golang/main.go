package main

import (
	"fmt"
	"os"
	"slices"
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

func mapValues(a map[string]int) []int {
	x := []int{}
	for _, val := range a {
		x = append(x, val)
	}

	return x
}

func countCards(hand string) []int {
	c := make(map[string]int)

	for _, card := range hand {
		char := string(card)
		if _, ok := c[char]; ok {
			c[char] += 1
		} else {
			c[char] = 1
		}
	}

	values := mapValues(c)
	slices.Sort(values)
	slices.Reverse(values)
	return values
}

func (h *HandCards) GetType() {
	counts := countCards(h.Cards)

	if counts[0] == 5 {
		h.Strength = 7
	} else if counts[0] == 4 {
		h.Strength = 6
	} else if counts[0] == 3 && counts[1] == 2 {
		h.Strength = 5
	} else if counts[0] == 3 {
		h.Strength = 4
	} else if counts[0] == 2 && counts[1] == 2 {
		h.Strength = 3
	} else if counts[0] == 2 {
		h.Strength = 2
	} else {
		h.Strength = 1
	}

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
