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

func TestToUpperCase(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"converts single characters", "a", "A"},
		{"converts multiple characters", "abcxyz", "ABCXYZ"},
		{"converts ONLY lower case characters", "aBcD eF123%@£", "ABCD EF123%@£"},
		{"leaves capital-only strings unchanged", "ABC", "ABC"},
		{"leaves strings without lower-case characters unchanged", "42&()", "42&()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toUpperCase(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestToLowerCase(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"converts single characters", "A", "a"},
		{"converts multiple characters", "ABCXYZ", "abcxyz"},
		{"converts ONLY upper case characters", "aBcD eF123%@£", "abcd ef123%@£"},
		{"leaves lower-only strings unchanged", "abc", "abc"},
		{"leaves strings without capitals characters unchanged", "42&()", "42&()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toLowerCase(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		want        string
	}{
		{"capitalizes single characters", "a", "A"},
		{"preserves single capitals", "A", "A"},
		{"capitalizes only the first character", "xyz", "Xyz"},
		{"leaves capitalized leading characters unchanged", "Liverpool", "Liverpool"},
		{"preserves numerical leading characters", "1st", "1st"},
		{"preserves mid-string capitalised characters", "mcDonald", "McDonald"},
		{"leaves strings without alpha-numeric characters unchanged", "42&()", "42&()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := capitalize(tt.inputString)
			if got != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}
