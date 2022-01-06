package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getOperation(str string) (valid bool, count int, function func(str string) string, snippet string) {
	type operation struct {
		function      func(string) string
		multipleWords bool
	}

	operations := map[string]operation{
		"hex": {hexToDecimal, false},
		"bin": {binToDecimal, false},
		"low": {toLowerCase, true},
		"up":  {toUpperCase, true},
		"cap": {capitalize, true},
	}

	const snippetIndex = 1
	const opIndex = 2
	const countIndex = 3

	puncMap := getPunctuation()
	punctuation := make([]rune, len(puncMap))

	puncIndex := 0
	for k := range puncMap {
		punctuation[puncIndex] = k
		puncIndex++
	}

	pattern := fmt.Sprintf(`^(\({1}([a-z]{1,3})(?:, ([0-9]*)){0,1}\){1})(?:$|\s|[%s]{1,})`, string(punctuation))
	re := regexp.MustCompile(pattern)

	// Extract data from operation
	if re.MatchString(str) {
		data := re.FindStringSubmatch(str)

		if op, exists := operations[data[opIndex]]; exists {
			// Count does not exist
			if len(data[countIndex]) == 0 {
				return true, 1, op.function, data[snippetIndex]
			} else if op.multipleWords {
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
				if tokens[position].kind == Word {
					count--
					tokens[position].str = function(tokens[position].str)
				}
			}
		}
	}

	return tokens
}
