package main

import (
	"testing"

	// testMain

	_ "embed"
)

//go:embed input.txt
var input string

func TestDay3FirstHalf(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"2", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"3", "PmmdzqPrVvPwwTWBwg", 42},
		{"4", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"4", input, 8401},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day3FirstHalf(tt.input); got != tt.want {
				t.Errorf("Day3() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestDay3SeconedHalf(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg", 18},
		{"2", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 52},
		{"4", input, 2641},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day3SecondHalf(tt.input); got != tt.want {
				t.Errorf("Day3() = %v, want %v", got, tt.want)
			}
		})
	}

}
