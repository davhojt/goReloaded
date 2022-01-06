package main

import "os"

// TODO: write tests
func writeFile(path string, text string) {
	bytes := []byte(text)
	err := os.WriteFile(path, bytes, 0755)
	if err != nil {
		panic("Failed while writing file")
	}
}
