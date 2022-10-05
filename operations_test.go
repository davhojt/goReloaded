package main

import "testing"

func TestGetOperation(t *testing.T) {
	var tests = []struct {
		name    string
		input   string
		valid   bool
		count   int
		snippet string
	}{
		{"identifies operation without count", "(cap)", true, 1, "(cap)"},
		{"identifies operation with count", "(up, 45)", true, 45, "(up, 45)"},
		{"identifies operation before punctuation", "(hex)...", true, 1, "(hex)"},
		{"identifies operation before white space", "(bin)\n \t", true, 1, "(bin)"},
		{"ignores invalid operations without count", "(bar)", false, 0, ""},
		{"ignores invalid operations with count", "(foo 42)", false, 0, ""},
		{"ignores brackets within words", "dust(bin)bag", false, 0, ""},
		{"ignores brackets at the start of words", "(cap)ital city", false, 0, ""},
		{"ignores brackets at the end of words", "look(up)", false, 0, ""},
		{"ignores operations with invalid count", "(up, f)", false, 0, ""},
		{"ignores operations with invalid white space delimitor count", "(cap,\n3)", false, 0, ""},
		{"ignores operations with no white space delimitor count", "(low,5)", false, 0, ""},
		{"ignores operations with invalid delimitor character", "(up- 7)", false, 0, ""},
		{"ignores operations with count which do not accept count.", "(bin, 5)", false, 0, ""},
		{"ignores operations which are not the next word", "this (cap)", false, 0, ""},
		{"ignores operations which are after spaces", " (cap)", false, 0, ""},
		{"ignores operations which are after punctuation", ".(cap)", false, 0, ""},
		{"ignores operations in caps", "(Hex)", false, 0, ""},
		{"ignores negative count", "(up, -2)", false, 0, ""},
		{"handles empty string", "", false, 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotCount, _, gotSnippet := getOperation(tt.input)
			if gotValid != tt.valid || gotCount != tt.count || gotSnippet != tt.snippet {
				t.Fatalf("got [%t][%d][%s], want [%t][%d][%s]", gotValid, gotCount, gotSnippet, tt.valid, tt.count, tt.snippet)
			}
		})
	}
}
