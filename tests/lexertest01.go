// Lexer tests for GoAsm65816
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 21. April 2018
// This version: 22. April 2018

// This is a lexer test based on the original version of TAN

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"
)

type token struct {
	Type  int
	Text  string
	Row   int
	Index int
}

const (
	EOF        int = iota
	EOL            // included for output formatting
	AMPERSAND      // '&'
	ANON_LABEL     // '@'
	COMMA          // ','
	COMMENT        // starts with ';'
	DIRECTIVE      // eg '.native'
	DOLLAR         // '$'
	LABEL          // eg 'loop'
	MINUS          // '-'
	OPCODE         // eg 'nop'
	PLUS           // '+'
	QUOTE          // '"'
	SEMICOLON      // ';'
	SLASH          // '/'
	STAR           // '*'
	SYMBOL         // Generic term for multi-character thingies
	LBRACE         // '{'
	RBRACE         // '}'
	TICK           // ' sign (single quote)
)

// This should probably be a map for speed, but we'll look at that later
var (
	input = flag.String("i", "", "Input file (REQUIRED)")

	tokenNames = []string{"EOF", "EOL", "AMPERSAND", "ANON_LABEL", "COMMA",
		"COMMENT", "DIRECTIVE", "DOLLAR", "LABEL", "MINUS", "OPCODE", "PLUS",
		"QUOTE", "SEMICOLON", "SLASH", "STAR", "SYMBOL", "LBRACE", "RBRACE",
		"TICK"}

	opcodes = []string{"dex", "inx", "nop", "pha", "lda.#", "ldx.#", "sta.d",
		"sta.x", "bne", "bra", "mvp"}

	singleChars      = []rune{'&', ',', '$', '+', '-', '/', '*', '{', '}', '\''}
	singleCharTokens = []int{AMPERSAND, COMMA, DOLLAR, PLUS, MINUS, SLASH, STAR, LBRACE,
		RBRACE, TICK}

	raw    []string
	tokens []token
)

// printToken is a temporary routine for testing
func printToken(t token) {
	if t.Type != EOL && t.Type != COMMENT {
		fmt.Print("<", tokenNames[t.Type], "> (", t.Row, ":", t.Index, ") [", t.Text, "]\n")
	}
}

// findAnyEOW takes an array of runes and returns the index of the first whitespace
// character. If there is no whitespace, it returns the length of the string.
// Note that this doesn't care what type of character we find, so it is useful
// for comments
func findAnyEOW(rs []rune) int {
	result := len(rs)

	for i := 0; i < len(rs); i++ {
		if unicode.IsSpace(rs[i]) {
			result = i
			break
		}
	}
	return result
}

// findSymbolEOW takes an array of runes and returns the index of the first
// non-symbol character (not a letter or a number or '.' or ':'). If there is none, it returns
// the length of the string
func findSymbolEOW(rs []rune) int {
	result := 0

	for i := 0; i < len(rs); i++ {

		// Skip ':' which is used in numbers to mark the bank byte
		if rs[i] == ':' {
			continue
		}

		// We need to include '.' and '#' because they are part of the
		// opcodes
		if unicode.IsLetter(rs[i]) || unicode.IsNumber(rs[i]) || rs[i] == '.' || rs[i] == '#' {
			result += 1
		} else {
			break
		}
	}
	return result
}

// isOpcode takes a string and checks to see if it is a legal 65816 opcode,
// returning a bool
func isOpcode(s string) bool {
	r := false

	for _, w := range opcodes {
		if s == w {
			r = true
			break
		}
	}
	return r
}

func main() {
	flag.Parse()

	// *** LOAD SOURCE FILE ***

	if *input == "" {
		log.Fatal("No input file provided. ABORTING.")
	}

	inputFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		raw = append(raw, scanner.Text())
	}

	// *** LEXER ***

	for row, line := range raw {

		// Skip empty lines immediately
		if len(line) == 0 {
			tokens = append(tokens, token{EOL, "", row + 1, 0})
			continue
		}

		// Move stuff to runes because we want to be able
		// to deal with Unicode
		chars := []rune(line)

		for i := 0; i < len(chars); i++ {

			// Stuff to skip in a line
			if unicode.IsSpace(chars[i]) {
				continue
			}

			// Get rid of comments, which we'll have a lot of
			// because of the way Scot writes code. This also gets
			// rid of the full line comments that begin at the start
			// of the line
			if chars[i] == ';' {
				tokens = append(tokens, token{COMMENT, string(chars[i:len(chars)]), row + 1, i})
				i = len(chars)
				continue
			}

			// Stuff that depends on being at start of line
			if i == 0 && !unicode.IsSpace(chars[0]) {
				c0 := chars[0]
				switch {

				case c0 == '@':
					tokens = append(tokens, token{ANON_LABEL, "@", row + 1, 0})
					continue

				case unicode.IsLetter(c0):
					endChar := findAnyEOW(chars[i:len(chars)])
					tokens = append(tokens, token{LABEL, string(chars[i : i+endChar]), row + 1, i})
					i = i + endChar - 1 // continue adds one
					continue
				}
			}

			// Multi-Character stuff that is not a label and not a
			// directive -- so numbers, symbols, and opcodes
			if unicode.IsLetter(chars[i]) || unicode.IsNumber(chars[i]) {
				endChar := findSymbolEOW(chars[i:len(chars)])
				word := chars[i : i+endChar]

				// Make sure we have a legal opcode
				if isOpcode(string(word)) {
					tokens = append(tokens, token{OPCODE, string(word), row + 1, i})
					i = i + endChar - 1 // continue adds one
					continue
				}

				tokens = append(tokens, token{SYMBOL, string(word), row + 1, i})
				i = i + endChar - 1 // continue adds one
				continue
			}

			// Single character tokenization (& and friends). Note @
			// is handled with the stuff at the beginning of the
			// line
			for j, r := range singleChars {
				if chars[i] == r {
					nt := token{singleCharTokens[j], string(r), row + 1, i}
					tokens = append(tokens, nt)
					break
				}
			}

			// This check for directives must come after check for
			// opcodes or else we'll find the "." in "ldx.#" etc
			if chars[i] == '.' {
				endChar := findAnyEOW(chars[i:len(chars)])
				tokens = append(tokens, token{DIRECTIVE, string(chars[i : i+endChar]), row + 1, i})
				i = i + endChar
				continue
			}
		}

		// At the end of the line, add a EOL token
		tokens = append(tokens, token{EOL, "", row + 1, len(line)})
	}

	// At the end of the file, add a EOF token
	tokens = append(tokens, token{EOF, "", len(raw) + 1, 0})

	// *** TEST: PRINT TOKENS ***
	for _, t := range tokens {
		printToken(t)
	}
}
