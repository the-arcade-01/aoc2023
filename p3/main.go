package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var input [][]string
	for scanner.Scan() {
		line := scanner.Text()

		row := []string{}
		for _, val := range line {
			row = append(row, string(val))
		}

		input = append(input, row)
	}

	part1(input)
	part2(input)
}

/* part1 */

func checkDigit(ch string) bool {
	if ch >= "0" && ch <= "9" {
		return true
	}
	return false
}

func checkPart(input [][]string, n int, m int, i int, j int) bool {
	part := false
	if i-1 >= 0 && j-1 >= 0 {
		if !checkDigit(input[i-1][j-1]) && input[i-1][j-1] != "." {
			part = true
		}
	}
	if i-1 >= 0 {
		if !checkDigit(input[i-1][j]) && input[i-1][j] != "." {
			part = true
		}
	}
	if i-1 >= 0 && j+1 < m {
		if !checkDigit(input[i-1][j+1]) && input[i-1][j+1] != "." {
			part = true
		}
	}
	if j-1 >= 0 {
		if !checkDigit(input[i][j-1]) && input[i][j-1] != "." {
			part = true
		}
	}
	if j+1 < m {
		if !checkDigit(input[i][j+1]) && input[i][j+1] != "." {
			part = true
		}
	}
	if i+1 < n && j-1 >= 0 {
		if !checkDigit(input[i+1][j-1]) && input[i+1][j-1] != "." {
			part = true
		}
	}
	if i+1 < n {
		if !checkDigit(input[i+1][j]) && input[i+1][j] != "." {
			part = true
		}
	}
	if i+1 < n && j+1 < m {
		if !checkDigit(input[i+1][j+1]) && input[i+1][j+1] != "." {
			part = true
		}
	}
	return part
}

func part1(input [][]string) {
	ans := 0

	n, m := len(input), len(input[0])
	for i := 0; i < n; i++ {
		num, part := "", false
		for j := 0; j < m; j++ {
			if checkDigit(input[i][j]) {
				num += input[i][j]
				part = part || checkPart(input, n, m, i, j)
			} else {
				if num != "" && part {
					value, err := strconv.Atoi(num)
					if err != nil {
						panic(err)
					}
					ans += value
				}
				num = ""
				part = false
			}
		}
		if num != "" && part {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			ans += value
		}
	}

	fmt.Printf("part1: %v\n", ans)
}

/* part2 */

func checkPart2(input [][]string, n int, m int, i int, j int) (bool, []string) {
	part := false
	var position_star []string
	if i-1 >= 0 && j-1 >= 0 {
		if !checkDigit(input[i-1][j-1]) && input[i-1][j-1] != "." && input[i-1][j-1] == "*" {
			part = true
			pos := string(rune(i-1)) + string(rune(j-1))
			position_star = append(position_star, pos)
		}
	}
	if i-1 >= 0 {
		if !checkDigit(input[i-1][j]) && input[i-1][j] != "." && input[i-1][j] == "*" {
			part = true
			pos := string(rune(i-1)) + string(rune(j))
			position_star = append(position_star, pos)
		}
	}
	if i-1 >= 0 && j+1 < m {
		if !checkDigit(input[i-1][j+1]) && input[i-1][j+1] != "." && input[i-1][j+1] == "*" {
			part = true
			pos := string(rune(i-1)) + string(rune(j+1))
			position_star = append(position_star, pos)
		}
	}
	if j-1 >= 0 {
		if !checkDigit(input[i][j-1]) && input[i][j-1] != "." && input[i][j-1] == "*" {
			part = true
			pos := string(rune(i)) + string(rune(j-1))
			position_star = append(position_star, pos)
		}
	}
	if j+1 < m {
		if !checkDigit(input[i][j+1]) && input[i][j+1] != "." && input[i][j+1] == "*" {
			part = true
			pos := string(rune(i)) + string(rune(j+1))
			position_star = append(position_star, pos)
		}
	}
	if i+1 < n && j-1 >= 0 {
		if !checkDigit(input[i+1][j-1]) && input[i+1][j-1] != "." && input[i+1][j-1] == "*" {
			part = true
			pos := string(rune(i+1)) + string(rune(j-1))
			position_star = append(position_star, pos)
		}
	}
	if i+1 < n {
		if !checkDigit(input[i+1][j]) && input[i+1][j] != "." && input[i+1][j] == "*" {
			part = true
			pos := string(rune(i+1)) + string(rune(j))
			position_star = append(position_star, pos)
		}
	}
	if i+1 < n && j+1 < m {
		if !checkDigit(input[i+1][j+1]) && input[i+1][j+1] != "." && input[i+1][j+1] == "*" {
			part = true
			pos := string(rune(i+1)) + string(rune(j+1))
			position_star = append(position_star, pos)
		}
	}
	return part, position_star
}

func part2(input [][]string) {
	ans := 0

	hash := map[string][]int{}
	n, m := len(input), len(input[0])
	for i := 0; i < n; i++ {
		num, part := "", false
		star_positions := map[string]bool{}
		for j := 0; j < m; j++ {
			if checkDigit(input[i][j]) {
				num += input[i][j]
				p, position := checkPart2(input, n, m, i, j)
				part = part || p
				for _, posi := range position {
					star_positions[posi] = true
				}
			} else {
				if num != "" && part {
					value, err := strconv.Atoi(num)
					if err != nil {
						panic(err)
					}
					for posi := range star_positions {
						hash[posi] = append(hash[posi], value)
					}
				}
				num = ""
				part = false
				star_positions = map[string]bool{}
			}
		}
		if num != "" && part {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			for posi := range star_positions {
				hash[posi] = append(hash[posi], value)
			}
		}
	}

	for _, val := range hash {
		if len(val) == 2 {
			ans += (val[0] * val[1])
		}
	}

	fmt.Printf("part2: %v\n", ans)
}
