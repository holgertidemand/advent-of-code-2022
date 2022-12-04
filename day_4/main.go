package day_4

import (
	"bufio"
	"strconv"
	"strings"
)

func Day4SecondHalf(input string) int {
	var score int
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var numbers []int
		for _, v := range strings.Split(strings.Join(strings.Split(scanner.Text(), ","), "-"), "-") {
			n, _ := strconv.Atoi(v)
			numbers = append(numbers, n)
		}

		inRange := func(value, low, high int) bool {
			return low <= value && value <= high
		}

		for i, v := range numbers {
			if i < 2 {
				if inRange(v, numbers[2], numbers[3]) {
					score++
					break
				}
			} else {
				if inRange(v, numbers[0], numbers[1]) {
					score++
					break
				}
			}

		}

	}

	return score
}

func Day4FirstHalf(input string) int {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var score int
	for scanner.Scan() {
		var numbers []int
		for _, v := range strings.Split(strings.Join(strings.Split(scanner.Text(), ","), "-"), "-") {
			n, _ := strconv.Atoi(v)
			numbers = append(numbers, n)
		}

		low1 := numbers[0]
		high1 := numbers[1]
		low2 := numbers[2]
		high2 := numbers[3]

		if low2 <= low1 && high1 <= high2 || low1 <= low2 && high2 <= high1 {
			score++
		}
	}
	return score
}
