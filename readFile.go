package main

import "os"

// TODO: write tests
func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic("Failed while reading file")
	}
	return string(bytes)
}
