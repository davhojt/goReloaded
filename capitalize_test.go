package main

import (
	"testing"
)

func TestCapitalize(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"capitalizes single characters", "a", "A"},
		{"preserves single capitals", "A", "A"},
		{"capitalizes only the first character", "xyz", "Xyz"},
		{"leaves capitalized leading characters unchanged", "Liverpool", "Liverpool"},
		{"preserves numerical leading characters", "1st", "1st"},
		{"preserves mid-string capitalised characters", "mcDonald", "McDonald"},
		{"leaves strings without alpha-numeric characters unchanged", "42&()", "42&()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := capitalize(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
