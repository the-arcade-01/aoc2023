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

/* part 1 */
func checkDigit(ch string) bool {
	if ch >= "0" && ch <= "9" {
		return true
	}
	return false
}

func part1(input []string) {
	ans := 0

	for _, line := range input {
		first, last := "", ""
		for i := 0; i < len(line); i++ {
			if checkDigit(string(line[i])) {
				if first == "" {
					first = string(line[i])
					last = string(line[i])
					continue
				}
				last = string(line[i])
			}
		}

		if first != "" && last != "" {
			str := first + last
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			ans += num
		}
	}
	fmt.Printf("part1: %v\n", ans)
}

/* part 2 */

func digit(ch string) string {
	switch ch {
	case "one", "1":
		return "1"
	case "two", "2":
		return "2"
	case "three", "3":
		return "3"
	case "four", "4":
		return "4"
	case "five", "5":
		return "5"
	case "six", "6":
		return "6"
	case "seven", "7":
		return "7"
	case "eight", "8":
		return "8"
	case "nine", "9":
		return "9"
	case "zero", "0":
		return "0"
	}
	return ch
}

func part2(input []string) {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	ans := 0

	for _, line := range input {
		firstMap := map[string]int{}
		lastMap := map[string]int{}

		for _, val := range nums {
			firstMap[val] = strings.Index(line, val)
			lastMap[val] = strings.LastIndex(line, val)
		}

		firstNum, lastNum, firstIndex, lastIndex := "", "", len(line)-1, 0

		for key, val := range firstMap {
			if val != -1 {
				if firstIndex >= val {
					firstIndex = val
					firstNum = digit(key)
				}
			}
		}
		for key, val := range lastMap {
			if val != -1 {
				if lastIndex <= val {
					lastIndex = val
					lastNum = digit(key)
				}
			}
		}
		if firstNum != "" && lastNum != "" {
			str := firstNum + lastNum
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}

			ans += num
		}
	}
	fmt.Printf("part2: %v\n", ans)
}
