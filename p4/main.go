package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	part1(input)
	part2(input)
}

/* part1 */

func part1(input []string) {
	ans := 0
	for _, line := range input {
		section := strings.Split(line, ":")
		cards := strings.Split(section[1], "|")
		wins := strings.Split(cards[0], " ")
		ours := strings.Split(cards[1], " ")

		hashWins := map[string]bool{}
		for _, val := range wins {
			if val != "" {
				hashWins[val] = true
			}
		}

		points := 0
		for _, val := range ours {
			if val != "" {
				if _, ok := hashWins[val]; ok {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
				}
			}
		}

		ans += points
	}

	fmt.Printf("part1: %v\n", ans)
}

/* part2 */

func part2(input []string) {
	ans := 0
	hash := map[int]int{}

	for _, line := range input {
		section := strings.Split(line, ":")
		card := strings.Split(section[0], ":")[0]
		cardArray := strings.Split(card, " ")

		numId, err := strconv.Atoi(cardArray[len(cardArray)-1])
		if err != nil {
			panic(err)
		}

		hash[numId]++

		cards := strings.Split(section[1], "|")
		wins := strings.Split(cards[0], " ")
		ours := strings.Split(cards[1], " ")

		hashWins := map[string]bool{}
		for _, val := range wins {
			if val != "" {
				hashWins[val] = true
			}
		}

		matchCards := 0
		for _, val := range ours {
			if val != "" {
				if _, ok := hashWins[val]; ok {
					matchCards++
				}
			}
		}

		for i := 1; i <= matchCards; i++ {
			hash[numId+i] += hash[numId]
		}
	}

	for _, val := range hash {
		ans += val
	}

	fmt.Printf("part2: %v\n", ans)
}
