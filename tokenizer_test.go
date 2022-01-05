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

func TestCorrectArticle(t *testing.T) {
	var tests = []struct {
		name    string
		article string
		noun    string
		want    string
	}{
		{"converts 'a' to 'an' when noun begins with vowel", "a", "apple", "an"},
		{"converts 'an' to 'a' when noun begins with consonant", "an", "guitar", "a"},
		{"preserves capitals when converting 'A' to 'An'", "A", "elephant", "An"},
		{"preserves capitals when converting 'An' to 'A'", "An", "table", "A"},
		{"does not convert 'an' when noun begins with vowel", "an", "arm", "an"},
		{"does not convert 'a' when noun begins with consonant", "a", "spoon", "a"},
		{"does not convert 'a' when noun begins with non letter", "a", "+tv", "a"},
		{"does not convert 'an' when noun begins with non letter", "an", "£money", "an"},
		{"does not convert 'a' when article is empty string", "a", "", "a"},
		{"does not convert 'an' when article is empty string", "an", "", "an"},
		{"does not convert when article is not 'a' or 'an'", "nice", "hat", "nice"},
		{"does not convert when article is empty string", "an", "", "an"},
		{"does not convert 'a' when noun is empty string", "a", "", "a"},
		{"does not convert 'an' when noun is empty string", "an", "", "an"},
		{"does not convert when article is only non letters", "2", "guitar", "2"},
		{"does not convert when article is space", " ", "guitar", " "},
		{"does not convert 'a' when noun is space", "a", " ", "a"},
		{"does not convert 'an' when noun is space", "an", " ", "an"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := correctArticle(tt.article, tt.noun)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
