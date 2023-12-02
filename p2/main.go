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

func parser(line string) int {
	colon_break := strings.Split(line, ":")

	gameId, err := strconv.Atoi(strings.Split(colon_break[0], " ")[1])
	if err != nil {
		panic(err)
	}

	sub_games := strings.Split(colon_break[1], ";")

	for _, sub := range sub_games {
		hash := map[string]int{}
		s := strings.Split(sub, ",")
		for _, c := range s {
			record := strings.Split(c, " ")
			val, err := strconv.Atoi(record[1])
			if err != nil {
				panic(err)
			}
			hash[record[2]] = val
		}
		if hash["red"] > 12 || hash["green"] > 13 || hash["blue"] > 14 {
			return 0
		}
	}

	return gameId
}

func part1(input []string) {
	ans := 0

	for _, line := range input {
		id := parser(line)
		ans += id
	}

	fmt.Printf("part1: %v\n", ans)
}

/* part2 */

func parser2(line string) int {
	colon_break := strings.Split(line, ":")

	sub_games := strings.Split(colon_break[1], ";")

	hash := map[string]int{}
	for _, sub := range sub_games {
		s := strings.Split(sub, ",")
		for _, c := range s {
			record := strings.Split(c, " ")
			val, err := strconv.Atoi(record[1])
			if err != nil {
				panic(err)
			}
			if hash[record[2]] < val {
				hash[record[2]] = val
			}
		}
	}
	power := hash["red"] * hash["blue"] * hash["green"]

	return power
}

func part2(input []string) {
	ans := 0

	for _, line := range input {
		ans += parser2(line)
	}

	fmt.Printf("part2: %v\n", ans)
}
