// Test file for lexer, part of the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 11. May 2018
// This version: 11. May 2018

package lexer

import "testing"

func TestFindBinEOW(t *testing.T) {
	var tests = []struct {
		input []rune
		want  int
	}{
		{[]rune("0"), 1},
		{[]rune("1"), 1},
		{[]rune("00"), 2},
		{[]rune("11"), 2},
		{[]rune("0:"), 2}, // colon counts as part of the length
		{[]rune("0."), 2}, // period counts as part of the length
		{[]rune("0010:"), 5},
		{[]rune("0010:"), 5},
		{[]rune("0010,"), 4},
		{[]rune("0010}"), 4},
		{[]rune("2001"), 0},
	}

	for _, test := range tests {
		got := findBinEOW(test.input)
		if got != test.want {
			t.Errorf("findBinEOW(%q) = %v", test.input, got)
		}
	}
}

func TestFindDecEOW(t *testing.T) {
	var tests = []struct {
		input []rune
		want  int
	}{
		{[]rune("0"), 1},
		{[]rune("1"), 1},
		{[]rune("00"), 2},
		{[]rune("11"), 2},
		{[]rune("0:"), 1}, // colon not allowed
		{[]rune("0."), 1}, // period not allowed
		{[]rune("0010"), 4},
		{[]rune("2001"), 4},
		{[]rune("2009+"), 4},
		{[]rune("2009}"), 4},
		{[]rune("+1"), 0},
	}

	for _, test := range tests {
		got := findDecEOW(test.input)
		if got != test.want {
			t.Errorf("findDecEOW(%q) = %v", test.input, got)
		}
	}
}

func TestFindHexEOW(t *testing.T) {
	var tests = []struct {
		input []rune
		want  int
	}{
		{[]rune("0"), 1},
		{[]rune("1"), 1},
		{[]rune("00"), 2},
		{[]rune("aa"), 2},
		{[]rune("FF"), 2},
		{[]rune("Ff"), 2},
		{[]rune("a:a"), 3}, // colon allowed
		{[]rune("a."), 2},  // period allowed
		{[]rune("beef"), 4},
		{[]rune("2001"), 4},
		{[]rune("2009+"), 4},
		{[]rune("aa:2009}"), 7},
		{[]rune("+1"), 0},
	}

	for _, test := range tests {
		got := findHexEOW(test.input)
		if got != test.want {
			t.Errorf("findHexEOW(%q) = %v", test.input, got)
		}
	}
}
