package day_5

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Day5First(input string) string {
	sections := strings.Split(input, "\n\n")
	stacks := createStacks(sections[0])

	lines := getLines(sections[1])

	for _, l := range lines {
		items, from, to := parseLine(l)

		tempFrom := make([]string, len(stacks[from]))
		tempTo := make([]string, len(stacks[to]))
		copy(tempFrom, stacks[from])
		copy(tempTo, stacks[to])

		for i := 0; i < items; i++ {
			tempFrom := make([]string, len(stacks[from]))
			tempTo := make([]string, len(stacks[to]))
			copy(tempFrom, stacks[from])
			copy(tempTo, stacks[to])

			var move, newFrom, newTo []string
			move = append(move, tempFrom[:1]...)
			stacks[from] = append(newFrom, tempFrom[1:]...)
			stacks[to] = append(move, append(newTo, tempTo...)...)
		}
	}

	var result string
	for i := range stacks {
		if len(stacks[i]) > 0 {
			result += stacks[i][0]
		}
	}

	return result
}

func Day5Second(input string) string {
	sections := strings.Split(input, "\n\n")
	stacks := createStacks(sections[0])
	lines := getLines(sections[1])

	for _, l := range lines {
		items, from, to := parseLine(l)

		tempFrom := make([]string, len(stacks[from]))
		tempTo := make([]string, len(stacks[to]))
		copy(tempFrom, stacks[from])
		copy(tempTo, stacks[to])

		var move, newFrom, newTo []string
		move = append(move, tempFrom[:items]...)
		stacks[from] = append(newFrom, tempFrom[items:]...)
		stacks[to] = append(move, append(newTo, tempTo...)...)
	}

	var result string
	for i := range stacks {
		if len(stacks[i]) > 0 {
			result += stacks[i][0]
		}
	}

	return result
}

func parseLine(line string) (int, int, int) {
	var items, from, to int
	for i, v := range regexp.MustCompile("[0-9]+").FindAllString(line, -1) {
		n, _ := strconv.Atoi(v)
		switch i {
		case 0:
			items = n
		case 1:
			from = n - 1
		case 2:
			to = n - 1
		}

	}

	return items, from, to
}

func getLines(input string) []string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}

func createStacks(input string) [][]string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var stacks [][]string
	for scanner.Scan() {
		columns := strings.Split((strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) || unicode.IsNumber(r) {
				return -1
			}

			if r == '[' || r == ']' {
				return -1
			}
			return r
		}, scanner.Text())), "")

		for i := 0; i < len(columns); i++ {
			if len(stacks) <= i {
				stacks = append(stacks, []string{})
			}

			if columns[i] == "*" {
				continue
			}

			stacks[i] = append(stacks[i], columns[i])
		}

	}

	return stacks
}
