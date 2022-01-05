package main

import (
	"testing"
)

func TestBinToDecimal(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"converts zero", "0", "0"},
		{"converts one", "1", "1"},
		{"converts small values", "11", "3"},
		{"converts max int64", "1111111111111111111111111111111111111111111111111111111111111111", "18446744073709551615"},
		{"converts max int64 - 1", "1111111111111111111111111111111111111111111111111111111111111110", "18446744073709551614"},
		{"handles leading zeros", "00101010", "42"},
		{"handles leading zeros", "00101010", "42"},
		{"handles empty string", "", "0"},
		{"handles empty string", "z", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binToDecimal(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
