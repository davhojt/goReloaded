package main

import "strconv"

// TODO: Consider handling leading '0x' and 'OX'
func hexToDecimal(str string) string {
	// TODO Handle error.
	value, _ := strconv.ParseUint(str, 16, 64)

	return strconv.FormatUint(value, 10)
}
