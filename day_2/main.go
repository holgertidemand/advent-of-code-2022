package main

import (
	"bufio"
	"strings"
)

type rule struct {
	player   string
	opponant string
	score    int
}

const Win = "Z"   // Scissors
const Draw = "Y"  // Paper
const Loose = "X" // Rock

// opponant is the key
const Rock = "A"
const Paper = "B"
const Scissors = "C"

var classicRules = [9]rule{
	// Rock
	{"X", Rock, 4},
	{"X", Paper, 1},
	{"X", Scissors, 7},
	// Paper
	{"Y", Rock, 8},
	{"Y", Paper, 5},
	{"Y", Scissors, 2},
	// Scissors
	{"Z", Rock, 3},
	{"Z", Paper, 9},
	{"Z", Scissors, 6},
}

var rules = map[string]map[string]string{
	Loose: {
		Rock:     "Z",
		Paper:    "X",
		Scissors: "Y",
	},
	Draw: {
		Rock:     "X",
		Paper:    "Y",
		Scissors: "Z",
	},
	Win: {
		Rock:     "Y",
		Paper:    "Z",
		Scissors: "X",
	},
}

func Day2(input string) (int, int) {
	var score, classicScore int

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), " ")
		opponant := game[0]
		player := game[1]

		choose := rules[player][opponant]

		for _, r := range classicRules {
			if r.player == choose && r.opponant == opponant {
				score += r.score

			}
			if r.player == player && r.opponant == opponant {
				classicScore += r.score
			}
		}

	}

	return score, classicScore
}
