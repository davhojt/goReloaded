package main

import (
	"testing"
)

// TODO: Does not test if correct function is returned
func TestGetTokenKind(t *testing.T) {
	var tests = []struct {
		name  string
		input rune
		want  tokenKind
	}{
		{"identifies space", ' ', WhiteSpace},
		{"identifies tab", '\t', WhiteSpace},
		{"identifies new line", '\n', WhiteSpace},
		{"identifies characters", 'a', Other},
		{"identifies number", '4', Other},
		{"identifies zero", '0', Other},
		{"identifies non-punctuation", '-', Other},
		{"identifies punctuation", '.', Punctuation},
		{"identifies single quote", '\'', Quote},
		{"handles null", 0, Other},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTokenKind(tt.input)
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}
