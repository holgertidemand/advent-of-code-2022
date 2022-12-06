package day_5

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed test_input.txt
var test_input string

// test day 5
func TestDay5(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Test", test_input, "CMZ"},
		{"Real", input, "GFTNRBZPF"},
	}

	// run tests
	for _, test := range tests {
		result := Day5First(test.input)
		if result != test.want {
			t.Errorf("Day5First(%s) = %s; expected %s", test.name, result, test.want)
		}
	}
}

func TestDay5Second(t *testing.T) {
	// test table
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Test", test_input, "MCD"},
		{"Real", input, "VRQWPDSGP"},
	}

	// run tests
	for _, test := range tests {
		result := Day5Second(test.input)
		if result != test.want {
			t.Errorf("Day5Second(%s) = %s; expected %s", test.name, result, test.want)
		}
	}
}
