package day_6

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func TestDay6First(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
		want2 int
	}{
		{
			name:  "test 1",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
			want2: 23,
		},
		{
			name:  "test 2",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
			want2: 23,
		},
		{
			name:  "test 3",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
			want2: 29,
		},
		{
			name:  "test 4",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
			want2: 26,
		},
		{
			name:  "real input",
			input: input,
			want:  1109,
			want2: 3965,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := Day6(tt.input); got1 != tt.want || got2 != tt.want2 {
				t.Errorf("Day6() = %v, %v want = %v, %v", got1, got2, tt.want, tt.want2)
			}
		})
	}
}
