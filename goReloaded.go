package main

// TODO: TEST
func processString(str string) string {
	tokens := tokenize(str)
	tokens = runOperations(tokens)

	return mergeTokens(tokens)
}

// TODO: TEST
func GoReloaded(input, output string) {
	text := ReadFile(input)
	writeFile(output, processString(text))
}
