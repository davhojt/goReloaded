package main

import (
	"testing"
)

func TestHexToDecimal(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"converts zero", "0", "0"},
		{"converts max int64", "FFFFFFFFFFFFFFFF", "18446744073709551615"},
		{"converts max int64 - 1", "FFFFFFFFFFFFFFFE", "18446744073709551614"},
		{"converts values which are equal to deciaml", "5", "5"},
		{"converts small values", "A", "10"},
		{"converts lower case", "1e", "30"},
		{"converts upper case", "1E", "30"},
		{"converts hex without letters", "123", "291"},
		{"handles leading zeros", "00000ABC", "2748"},
		{"handles empty string", "", "0"},
		{"handles empty string", "z", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hexToDecimal(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
