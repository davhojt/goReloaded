package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		input := os.Args[1]
		output := os.Args[2]
		GoReloaded(input, output)
	} else {
		panic("usage: go run . path_to_input path_to_output")
	}
}
