package day_4

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestDay4(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"1", "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n", 2},
		{"2", "2-8,3-7\n11-95,11-97\n6-6,4-6\n66-66,45-66\n", 4},
		{"3", "6-6,4-6", 1},
		{"4", "66-66,45-66\n", 1},
		{"5", "11-95,11-97\n", 1},
		{"6", "92-92,19-93\n3-94,3-96", 2},
		{"7", "32-43,31-53\n17-80,16-18\n22-65,64-64\n62-63,10-62\n3-99,2-98\n23-92,23-24\n92-92,19-93\n3-96,3-94", 5},
		{"8", input, 498},
	}

	for _, test := range tests {
		if got := Day4FirstHalf(test.input); got != test.want {
			t.Errorf("Day4FirstHalf(%q) = want: %d, gort: %d", test.name, test.want, got)
		}
	}

	tests = []struct {
		name  string
		input string
		want  int
	}{
		{"second 1", "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n", 4},
		{"second 2", input, 859},
	}

	for _, test := range tests {
		if got := Day4SecondHalf(test.input); got != test.want {
			t.Errorf("Day4FirstHalf(%q) = want: %d, gort: %d", test.name, test.want, got)
		}
	}

}
