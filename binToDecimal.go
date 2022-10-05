package main

import "strconv"

func binToDecimal(str string) string {
	// TODO Handle error.
	value, _ := strconv.ParseUint(str, 2, 64)

	return strconv.FormatUint(value, 10)
}
