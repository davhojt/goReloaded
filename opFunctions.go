package main

import (
	"strconv"
	"strings"
)

// TODO: Consider handling leading '0x' and 'OX'
func hexToDecimal(str string) string {
	value, _ := strconv.ParseUint(str, 16, 64)

	return strconv.FormatUint(value, 10)
}

func binToDecimal(str string) string {
	value, _ := strconv.ParseUint(str, 2, 64)

	return strconv.FormatUint(value, 10)
}

// TODO: This capitalises after every whitespace or hyphen. Same as toLower Case.
func toUpperCase(str string) string {
	return strings.ToUpper(str)
}

// TODO: This capitalises after every whitespace or hyphen. Same as toUpperCase.
func toLowerCase(str string) string {
	return strings.ToLower(str)
}

func capitalize(str string) string {
	return strings.Title(str)
}
