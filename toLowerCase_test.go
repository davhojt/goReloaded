package main

import (
	"testing"
)

func TestToLowerCase(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"converts single characters", "A", "a"},
		{"converts multiple characters", "ABCXYZ", "abcxyz"},
		{"converts ONLY upper case characters", "aBcD eF123%@£", "abcd ef123%@£"},
		{"leaves lower-only strings unchanged", "abc", "abc"},
		{"leaves strings without capitals characters unchanged", "42&()", "42&()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toLowerCase(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
