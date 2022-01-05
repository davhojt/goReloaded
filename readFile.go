package main

import "os"

// TODO: write tests
func ReadFile(path string) string {
	// TODO: handle error
	// TODO: hanle if file does not exist
	bytes, _ := os.ReadFile(path)
	return string(bytes)
}
