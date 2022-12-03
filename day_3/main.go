package main

import (
	"bufio"
	"sort"
	"strings"
)

func Day3FirstHalf(input string) int {
	var score int
	alphabet := make([]string, 0, 52)
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")

		length := len(text) / 2
		// split text in the middle
		first := text[:length]
		second := text[length:]

		cache := make(map[string]struct{})
		for i := 0; i < length; i++ {

			for j := 0; j < length; j++ {

				if first[i] == second[j] {
					if _, ok := cache[first[i]]; ok {
						continue
					}
					cache[first[i]] = struct{}{}

					for k := 0; k < len(alphabet); k++ {

						if first[i] == alphabet[k] {
							score += k + 1
							break
						}
					}

					break
				}

			}
		}
	}

	return score
}

func Day3SecondHalf(input string) int {
	var score int
	alphabet := make([]string, 0, 52)
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	buffer := make([]string, 0, 3)

	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())

		if len(buffer) == 3 {

			//sort buffer by length
			sort.Slice(buffer, func(i, j int) bool {
				return len(buffer[i]) < len(buffer[j])
			})

			cache := make(map[string]struct{})
			for i := 0; i < len(buffer[0]); i++ {
				for j := 0; j < len(buffer[1]); j++ {
					for k := 0; k < len(buffer[2]); k++ {

						if buffer[0][i] == buffer[1][j] && buffer[1][j] == buffer[2][k] {
							_, ok := cache[string(buffer[0][i])]
							if ok {
								continue
							}

							cache[string(buffer[0][i])] = struct{}{}

							for l := 0; l < len(alphabet); l++ {
								if string(buffer[0][i]) == alphabet[l] {
									score += l + 1
									break
								}
							}
							break
						}
					}
				}
			}

			buffer = make([]string, 0, 3)
		}

	}

	return score
}
