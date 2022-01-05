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

// TODO: TEST
func getPunctuation() map[rune]bool {
	return map[rune]bool{
		'.': false,
		',': false,
		'!': false,
		'?': false,
		':': false,
		';': false,
	}
}

func getTokenKind(r rune) tokenKind {
	if unicode.IsSpace(r) {
		return WhiteSpace
	}

	// TODO: Make Global
	punctuation := getPunctuation()

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

// TODO: TEST
func correctArticle(article, noun string) string {
	type sound int64

	const (
		Consonant sound = iota
		Vowel
		NonLetter
	)

	soundType := func(r rune) sound {
		vowels := map[rune]bool{
			'a': false,
			'e': false,
			'i': false,
			'o': false,
			'u': false,
			'h': false,
		}

		if unicode.IsLetter(r) {
			_, isVowel := vowels[unicode.ToLower(r)]
			if isVowel {
				return Vowel
			}
			return Consonant
		}

		return NonLetter
	}

	wordBeginsWith := func(str string) sound {
		var r rune
		if len(str) > 0 {
			runes := []rune(str)
			r = runes[0]
		}
		return soundType(r)
	}

	if len(noun) > 0 {
		if toLowerCase(article) == "a" && wordBeginsWith(noun) == Vowel {
			return article[0:1] + "n"
		}

		if toLowerCase(article) == "an" && wordBeginsWith(noun) == Consonant {
			return article[0:1]
		}
	}

	return article
}

// TODO: TEST
func mergeTokens(tokens []token) string {
	str := ""

	inQuoute := false

	for i := range tokens {
		buffer := ""
		spaceBefore := ""
		spaceAfter := ""

		writeString := func() {
			str += spaceBefore + buffer + spaceAfter
			buffer, spaceBefore, spaceAfter = "", "", ""
		}

		// Adds punctuation
		if tokens[i].kind == Punctuation {
			buffer = tokens[i].str
			spaceAfter = " "
		}

		if tokens[i].kind == Other {
			// TODO: This runs every time, even if the current word is not an article.
			// Indefinite article check
			for nextToken := i + 1; nextToken < len(tokens); nextToken++ {
				if tokens[nextToken].kind == Punctuation {
					break
				}

				if tokens[nextToken].kind == Other {
					tokens[i].str = correctArticle(tokens[i].str, tokens[nextToken].str)
					break
				}
			}

			buffer = tokens[i].str

			// Add space after word
			for nextToken := i + 1; nextToken < len(tokens); nextToken++ {
				if tokens[nextToken].kind == WhiteSpace || tokens[nextToken].kind == Operation {
					continue
				} else if tokens[nextToken].kind == Other {
					spaceAfter = " "
				} else {
					break
				}
			}
		}

		// Adds quotes
		if tokens[i].kind == Quote {
			buffer = tokens[i].str
			if i == 0 || tokens[i-1].kind == WhiteSpace || i >= len(tokens)-1 || tokens[i+1].kind == WhiteSpace {
				if inQuoute {
					spaceAfter = " "
				} else {
					spaceBefore = " "
				}
				inQuoute = !inQuoute
			}
		}

		// Removes spaces from start and end of string
		if i == 0 {
			spaceBefore = ""
		}
		if i == len(tokens)-1 {
			spaceAfter = ""
		}

		writeString()
	}

	return str
}
