package main

import "unicode"

type tokenKind int64

const (
	NoneKind tokenKind = iota
	Other
	WhiteSpace
	Punctuation
	Quote
	Operation
)

type token struct {
	str      string
	kind     tokenKind
	function *func(str string) string
	count    int
}

func getTokenKind(r rune) tokenKind {
	if unicode.IsSpace(r) {
		return WhiteSpace
	}

	// TODO: Make Global
	punctuation := map[rune]bool{
		'.': false,
		',': false,
		'!': false,
		'?': false,
		':': false,
		';': false,
	}

	if _, exists := punctuation[r]; exists {
		return Punctuation
	}

	if r == '\'' {
		return Quote
	}

	return Other
}

// TODO: Test
func tokenize(str string) []token {
	var tokens []token

	currentKind := NoneKind
	start := 0
	// end := 0

	// Skips last character bug
	for i, r := range str {
		if i < start {
			continue
		}

		addToken := func(tokenStr string, kind tokenKind, function *func(str string) string, count int) {
			if currentKind != NoneKind {
				tokens = append(tokens, token{tokenStr, kind, function, count})
			}
		}

		addCurrentToken := func() {
			addToken(str[start:i], currentKind, nil, 0)
		}

		// 1. Checks if it is an operation Operation and creates token
		// 2. Checks other token types
		// 3. Get's last token
		if valid, count, function, snippet := getOperation(str[i:]); valid {
			addCurrentToken()
			addToken(snippet, Operation, &function, count)

			start = i + len(snippet)
			currentKind = NoneKind
		} else {
			if runeKind := getTokenKind(r); runeKind != currentKind {
				addCurrentToken()

				start = i
				currentKind = runeKind
			}

			if i == len(str)-1 {
				addToken(str[start:i+1], currentKind, nil, 0)
			}
		}
	}
	return tokens
}
