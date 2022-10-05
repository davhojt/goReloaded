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

func TestRunOperations(t *testing.T) {
	one := func(str string) string { return str + "---one---" }
	two := func(str string) string { return str + "***two***" }

	space := token{
		"   ",
		WhiteSpace,
		nil,
		0,
	}

	characters := token{
		"abc",
		Other,
		nil,
		0,
	}

	charactersOne := token{
		"abc---one---",
		Other,
		nil,
		0,
	}

	charactersOneOne := token{
		"abc---one------one---",
		Other,
		nil,
		0,
	}

	charactersOneTwo := token{
		"abc---one---***two***",
		Other,
		nil,
		0,
	}

	charactersTwo := token{
		"abc***two***",
		Other,
		nil,
		0,
	}

	punctuation := token{
		".,;",
		Punctuation,
		nil,
		0,
	}

	operationOne := token{
		"(one)",
		Operation,
		&one,
		1,
	}

	operationOneCountThree := token{
		"(one)",
		Operation,
		&one,
		3,
	}

	operationTwo := token{
		"(two)",
		Operation,
		&two,
		1,
	}

	operationTwoCountFour := token{
		"(two)",
		Operation,
		&two,
		4,
	}

	var tests = []struct {
		name  string
		input []token
		want  []token
	}{
		{
			"only spaces",
			[]token{space},
			[]token{space},
		},
		{
			"consecutive spaces",
			[]token{space, space, space},
			[]token{space, space, space},
		},
		{
			"only others",
			[]token{characters},
			[]token{characters},
		},
		{
			"consecutive others",
			[]token{characters, characters, characters},
			[]token{characters, characters, characters},
		},
		{
			"only punctuation",
			[]token{punctuation},
			[]token{punctuation},
		},
		{
			"consecutive punctuation",
			[]token{punctuation, punctuation, punctuation},
			[]token{punctuation, punctuation, punctuation},
		},
		{
			"only operation",
			[]token{operationOne},
			[]token{operationOne},
		},
		{
			"consecutive operation",
			[]token{operationOne, operationOne, operationOne},
			[]token{operationOne, operationOne, operationOne},
		},
		{
			"does not modify punctuation",
			[]token{punctuation, characters},
			[]token{punctuation, characters},
		},
		{
			"does not modify spaces",
			[]token{space, characters},
			[]token{space, characters},
		},
		{
			"applies single operation to no word",
			[]token{operationOne, characters},
			[]token{operationOne, characters},
		},
		{
			"applies single operation to single word",
			[]token{characters, operationOne},
			[]token{charactersOne, operationOne},
		},
		{
			"applies single operation to last word",
			[]token{characters, characters, characters, operationOne},
			[]token{characters, characters, charactersOne, operationOne},
		},
		{
			"applies single operation to middle word",
			[]token{characters, operationOne, characters},
			[]token{charactersOne, operationOne, characters},
		},
		{
			"applies single operation across space",
			[]token{characters, space, operationOne},
			[]token{charactersOne, space, operationOne},
		},
		{
			"applies single operation across punctuation",
			[]token{characters, punctuation, operationOne},
			[]token{charactersOne, punctuation, operationOne},
		},
		{
			"applies operation to correct number of elements",
			[]token{characters, characters, characters, punctuation, characters, operationOneCountThree},
			[]token{characters, charactersOne, charactersOne, punctuation, charactersOne, operationOneCountThree},
		},
		{
			"gracefully overflows when count exceeds words",
			[]token{characters, punctuation, characters, operationOneCountThree},
			[]token{charactersOne, punctuation, charactersOne, operationOneCountThree},
		},
		{
			"applies different operations to the same word",
			[]token{characters, operationOne, operationTwo},
			[]token{charactersOneTwo, operationOne, operationTwo},
		},
		{
			"applies same operations multiple times",
			[]token{characters, operationOne, operationOne},
			[]token{charactersOneOne, operationOne, operationOne},
		},
		{
			"applies different multiple operations",
			[]token{characters, characters, punctuation, characters, characters, characters, operationOneCountThree, operationTwoCountFour, characters},
			[]token{characters, charactersTwo, punctuation, charactersOneTwo, charactersOneTwo, charactersOneTwo, operationOneCountThree, operationTwoCountFour, characters},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runOperations(tt.input)

			for i := range tt.want {
				if i >= len(got) || i >= len(tt.want) ||
					got[i].count != tt.want[i].count ||
					got[i].str != tt.want[i].str ||
					got[i].kind != tt.want[i].kind {
					t.Fatalf("\ngot %v\nwant %v", got, tt.want)
				}
			}
		})
	}
}
