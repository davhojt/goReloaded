package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getOperation(word string) (valid bool, count int, function func(str string) string, snippet string) {
	type operation struct {
		function      func(string) string
		multipleWords bool
	}

	snippetIndex := 1
	opIndex := 2
	countIndex := 3

	// TODO: Handle punctuation
	keys := []string{"."}
	punctuation := strings.Join(keys, "")

	pattern := fmt.Sprintf(`^(\({1}([a-z]{1,3})(?:, ([0-9]*)){0,1}\){0,1})(?:$|\s|[%s]{1,})`, punctuation)

	re := regexp.MustCompile(pattern)

	operations := map[string]operation{
		"hex": {hexToDecimal, false},
		"bin": {binToDecimal, false},
		"low": {toLowerCase, true},
		"up":  {toUpperCase, true},
		"cap": {capitalize, true},
	}

	// Extract data from operation
	if re.MatchString(word) {
		// TODO: Handle spaces and commas. Previously splits by whole words.
		data := re.FindStringSubmatch(word)

		if op, exists := operations[data[opIndex]]; exists {
			// Count does not exist
			if len(data[countIndex]) == 0 {
				return true, 1, op.function, data[snippetIndex]
			} else if op.multipleWords {
				// TODO: operation must certainly handle positive cases.
				// TODO: Handle negatives, here otr in return.
				// TODO: Handle 12ioaiouo is probably valid for ATOI. SHould only be na int.
				count, err := strconv.Atoi(data[countIndex])

				if err == nil {
					return true, count, op.function, data[snippetIndex]
				}
			}
		}
	}

	// Word is not op code
	return false, 0, nil, ""
}

// TODO: Probably shouldn't modify the input (look at pointers).
func runOperations(tokens []token) []token {

	for i, t := range tokens {
		if t.kind == Operation {
			function := *t.function

			count := t.count

			for position := i - 1; position >= 0 && count > 0; position-- {
				if tokens[position].kind == Other {
					count--
					tokens[position].str = function(tokens[position].str)
				}
			}
		}
	}

	return tokens
}
