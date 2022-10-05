package main

import (
	"testing"
)

func TestToUpperCase(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"converts single characters", "a", "A"},
		{"converts multiple characters", "abcxyz", "ABCXYZ"},
		{"converts ONLY lower case characters", "aBcD eF123%@£", "ABCD EF123%@£"},
		{"leaves capital-only strings unchanged", "ABC", "ABC"},
		{"leaves strings without lower-case characters unchanged", "42&()", "42&()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toUpperCase(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
