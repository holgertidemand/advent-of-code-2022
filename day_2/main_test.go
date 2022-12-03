package main

// testMain

import (
	_ "embed"
	"testing"
)

// embed input.txt into the binary
//
//go:embed input.txt
var input string

func TestDay2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 int
		want2 int
	}{
		{"1", "A Y", 4, 8},
		{"2", "B X", 1, 1},
		{"3", "C Z", 7, 6},
		{"4", input, 13600, 11386},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got2 := Day2(tt.input); got != tt.want1 || got2 != tt.want2 {
				t.Errorf("Day2() = %v, want %v", got, tt.want1)
				t.Errorf("Day2() = %v, want %v", got2, tt.want2)
			}
		})
	}

}
