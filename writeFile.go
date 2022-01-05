package main

import "os"

// TODO: write tests
func writeFile(path string, text string) {
	// TODO: handle error if WriteFile fails.
	// TODO: hanle if file already exists?
	bytes := []byte(text)
	os.WriteFile(path, bytes, 0755)
}
