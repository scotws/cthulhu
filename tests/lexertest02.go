// Lexer test for GoAsm65816
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 21. April 2018
// This version: 24. April 2018

// This is a lexer test based on an experimental version of the Simpler
// Assembler Notation (SAN). Note that the output is not only used for straight
// assembly but also to produce a listing and a nicely formatted source code

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	EOF           int = iota
	EOL               // included for output formatting
	AMPERSAND         // '&'
	ANON_LABEL        // '@'
	BIN_NUMBER        // starts with '%'
	BRANCH_OPCODE     // "bra" and friends
	COMMA             // ','
	COMMENT           // starts with ';'
	DEC_NUMBER        // starts with 0-9
	DIRECTIVE         // eg '.native'
	HEX_NUMBER        // starts with '$' or '0x'
	JUMP_OPCODE       // "jmp", "rts", "jsr" etc
	LABEL             // eg 'loop'
	LOCAL_LABEL       // start with '_' limited to scopes
	MINUS             // '-'
	NOP_OPCODE        // "nop", "wdm"
	OPCODE            // generic opcode if no special one found
	PERCENT           // '%'
	PLUS              // '+'
	SEMICOLON         // ';'
	SLASH             // '/'
	STAR              // '*'
	STRING            // Does not include quotation marks
	SYMBOL            // Generic term for multi-character thingies
)

type token struct {
	Type int    // see list of constants above
	Text string // original source code representation
	// Value int // decimal value for bin and hex numbers
	// Bytes []byte // for opcodes and parameters, final value as bytes
	// File string
	Row   int
	Index int
}

// Most of these should be stored in maps for speed, but will look into that
// later
var (
	input = flag.String("i", "", "Input file (REQUIRED)")

	tokenNames = []string{"EOF", "EOL", "AMPERSAND", "ANON_LABEL", "BIN_NUMBER",
		"BRANCH_OPCODE",
		"COMMA", "COMMENT", "DEC_NUMBER", "DIRECTIVE", "HEX_NUMBER",
		"JUMP_OPCODE", "LABEL", "LOCAL_LABEL",
		"MINUS", "NOP_OPCODE", "OPCODE", "PERCENT", "PLUS", "SEMICOLON", "SLASH",
		"STAR", "STRING", "SYMBOL"}

	opcodes = []string{"cmp.#", "dex", "dey", "inx", "nop", "pha", "lda.#", "ldx.#", "lda.x",
		"ldy.#", "sta.d", "sta.x", "beq", "bne", "bra", "mvp", "jsr", "jmp", "rts"}

	// Separate detection of branches and jumps allows syntax
	// coloring for output
	branches = []string{"beq", "bne", "bra"}
	jumps    = []string{"jmp", "jsr", "rts"}
	noops    = []string{"nop"}

	directives = []string{".equ", ".mpu", ".origin", ".native", ".byte", ".end",
		".axy16", ".axy8", ".scope", ".scend", ".word"}

	// Minus and plus must be part of this list because they are used for
	// anonymous label branches, not just math
	singleChars      = []rune{',', '-', '+', '@', '+', '-', '/', '*'}
	singleCharTokens = []int{COMMA, MINUS, PLUS, AMPERSAND, PLUS, MINUS,
		SLASH, STAR}

	raw    []string
	tokens []token
)

// contains takes a string and a list of strings (words in our case) and returns
// a bool if the string is found in the list. This is what other languages do
// out of the box with "in"

func contains(s string, ws []string) bool {
	r := false
	for _, w := range ws {
		if s == w {
			r = true
			break
		}
	}
	return r
}

// isSingleCharToken takes a rune and checks if it is one of the characters we
// consider a single-rune token. Returns a bool as a sign of success, and if
// true, the corresponding int as the token type.
func isSingleCharToken(r rune) (int, bool) {
	var t int
	f := false
	for i, c := range singleChars {
		if c == r {
			t = singleCharTokens[i]
			f = true
			break
		}
	}
	return t, f
}

// isLegalSymbolStart takes a rune and checks if it is allowed as the start of
// a symbol.
func isLegalSymbolStart(r rune) bool {
	f := false
	if unicode.IsLetter(r) ||
		r == '.' ||
		r == '_' ||
		r == ':' {
		f = true
	}
	return f
}

// isLegalSymbolChar takes a rune and checks if it is allowed as part of the
// symbol body. There is a separate check for the beginning of the symbol. Note
// the math functions (-,+,/,*) are not legal symbol chars. Note this should
// support unicode symbol names
func isLegalSymbolChar(r rune) bool {
	f := false
	if unicode.IsLetter(r) ||
		unicode.IsNumber(r) ||
		r == '_' ||
		r == '&' ||
		r == '!' ||
		r == '?' ||
		r == '#' ||
		r == '.' || // used in opcode mnemonics
		r == '\'' || // single quote ("tick")
		r == ':' {
		f = true
	}
	return f
}

// printToken is a temporary routine for testing
func printToken(t token) {
	fmt.Print("<", tokenNames[t.Type], "> (", t.Row, ":", t.Index, ") [", t.Text, "]\n")
}

// findBinNumberEOW takes an array of runes and returns the index of the first
// non-binary number character (0 and 1). If there is none, it returns the
// length of the string
func findBinNumberEOW(rs []rune) int {
	result := len(rs)

	for i := 0; i < len(rs); i++ {

		// Allow '.' and ':' as cosmetic separators in binary numbers. They will
		// be actually filtered out by the parser during conversion
		if rs[i] == ':' || rs[i] == '.' {
			continue
		}

		if rs[i] != '0' && rs[i] != '1' {
			result = i
			break
		}
	}
	return result
}

// findDecNumberEOW takes an array of runes and returns the index of the first
// non-decimal number character (0 to 9). If there is none, it returns the
// length of the string. In contrast to binary and hex numbers, we don't allow
// '.' and ':' as cosmetic separators
func findDecNumberEOW(rs []rune) int {
	result := len(rs)

	for i := 0; i < len(rs); i++ {

		if !unicode.IsNumber(rs[i]) {
			result = i
			break
		}
	}
	return result
}

// findHexNumberEOW takes an array of runes and returns the index of the first
// non-hexadecimal character (0 to 9, a to f or A to F). If there is none, it
// returns the length of the string
func findHexNumberEOW(rs []rune) int {
	var s string
	result := len(rs)

	for i := 0; i < len(rs); i++ {

		// Allow '.' and ':' as cosmetic separators in hex numbers, which
		// are used especially for the bank byte. They will be actually
		// filtered out by the parser during conversion
		if rs[i] == ':' || rs[i] == '.' {
			continue
		}

		s = string(rs[i])
		_, err := strconv.ParseUint(s, 16, 64)

		if err != nil {
			result = i
			break
		}
	}
	return result
}

// findStringEOW starts at a quote mark and searches for the next quote mark,
// returning its last character's index. It also returns a success bool. If no
// quote mark was, found return a zero as an it and a false bool
func findStringEOW(rs []rune) (int, bool) {
	f := false
	t := 0

	for i, c := range rs {
		if c == '"' {
			t = i // don't include the closing quote itself
			f = true
			break
		}
	}
	return t, f
}

// findSymbolEOW takes an array of runes and returns the index of the first
// character that is not legal for a symbol. If there is not such character in
// the array, it returns the length of the string.
func findSymbolEOW(rs []rune) int {
	result := len(rs)

	// We start one character in because if we landed here, we know that the
	// first letter of the symbol is legal
	for i := 1; i < len(rs); i++ {
		if !isLegalSymbolChar(rs[i]) || unicode.IsSpace(rs[i]) {
			result = i
			break
		}
	}
	return result
}

// addToken takes the token identifier, the actual text of the token from the
// source code, the row and index the token was found in and adds it to the
// token stream
func addToken(ti int, s string, r int, i int) {
	s0 := strings.TrimSpace(s)
	r0 := r + 1 // computers count row from 0, we from one
	tokens = append(tokens, token{ti, s0, r0, i})
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

	// ***** LEXER *****

	var tt int // temporary token

	for row, line := range raw {

		// Skip empty lines immediately. We put in a EOL to let the user
		// at least decide where they want to be
		if len(line) == 0 {
			addToken(EOL, "", row, 0)
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
			// of the line. Note we don't distinguish between
			// full-line comments (start with ;;) and inline
			// comments (start with ;), that is left up to the
			// formatter at a later stage
			if chars[i] == ';' {
				addToken(COMMENT, string(chars[i:len(chars)]), row, i)
				i = len(chars)
				continue
			}

			// Single character tokenization (@ and friends).
			st, found := isSingleCharToken(chars[i])
			if found {
				addToken(st, string(chars[i]), row, i)
				continue
			}

			// Handle hexadecimal numbers. This must come before the
			// test of for decimal numbers
			if chars[i] == '$' {
				i += 1 // skip '$' symbol
				endChar := findHexNumberEOW(chars[i:len(chars)])
				word := chars[i : i+endChar]
				addToken(HEX_NUMBER, string(word), row, i)
				i = i + endChar - 1 // Continue adds one
				continue
			}

			// Handle binary numbers. This must come before the test
			// for decimal numbers as well
			if chars[i] == '%' {
				i += 1 // skip '%' prefix
				endChar := findBinNumberEOW(chars[i:len(chars)])
				word := chars[i : i+endChar]
				addToken(BIN_NUMBER, string(word), row, i)
				i = i + endChar - 1 // Continue adds one
				continue
			}

			// Handle decimal numbers
			if unicode.IsNumber(chars[i]) {
				endChar := findDecNumberEOW(chars[i:len(chars)])
				word := chars[i : i+endChar]
				addToken(DEC_NUMBER, string(word), row, i)
				i = i + endChar - 1 // continue adds one
				continue
			}

			// Handle strings
			if chars[i] == '"' {
				i += 1 // skip leading quote
				endChar, ok := findStringEOW(chars[i:len(chars)])
				if !ok {
					log.Fatal("LEXER FATAL (", row+1, ":", i, "): No closing quote for string")
				}

				word := chars[i : i+endChar]
				addToken(STRING, string(word), row, i)
				i = i + endChar // skip over final quote
				continue
			}

			// Multi-Character stuff that is not a string -- so
			// labels, opcodes and symbols. We do this by first
			// getting the complete string till the next whitespace
			// and then figuring out what to do with it.
			if isLegalSymbolStart(chars[i]) {

				endChar := findSymbolEOW(chars[i:len(chars)])
				word := chars[i : i+endChar]

				// Check if we have a legal opcode. This must
				// come before the check for directives or the
				// period in opcodes such as "lda.#" will
				// trigger the directive finder
				if contains(string(word), opcodes) {

					// We can futher sort out the branches
					// and the jumps to make it easier for
					// the parser and also to allow
					// formatted output
					switch {

					case contains(string(word), branches):
						tt = BRANCH_OPCODE
					case contains(string(word), jumps):
						tt = JUMP_OPCODE
					case contains(string(word), noops):
						tt = NOP_OPCODE
					default:
						tt = OPCODE
					}

					addToken(tt, string(word), row, i)
					i = i + endChar - 1 // continue adds one
					continue
				}

				// Make sure we have a legal directive
				if contains(string(word), directives) {
					addToken(DIRECTIVE, string(word), row, i)
					i = i + endChar - 1 // continue adds one
					continue
				}

				// If we landed here, we've got a symbol of some
				// sort. First check if it is a label
				if strings.HasSuffix(string(word), ":") {
					// We remove the ':' at the end of the
					// label name to make symbol stuff
					// easier

					if word[0] == '_' {
						tt = LOCAL_LABEL
					} else {
						tt = LABEL
					}

					addToken(tt, string(chars[i:i+endChar-1]), row, i)
					i = i + endChar - 1 // continue adds one
					continue
				}

				// If it is not an opcode, a directive, or a
				// label, it's pretty much got to be a normal
				// symbol of some sort, and we'll let the parser
				// figure it out later. This means we catch all
				// errors as "symbols" at this point
				addToken(SYMBOL, string(chars[i:i+endChar]), row, i)
				i = i + endChar - 1
			}

		}
		// At the end of the line, add an EOL token to make life easier for the
		// parser
		addToken(EOL, "", row, len(line)-1)
	}

	// At the end of the file, add a EOF token
	addToken(EOF, "", len(raw), 0)

	// *** TEST: PRINT TOKENS ***
	for _, t := range tokens {
		printToken(t)
	}

	// *** TEST: PRINT RAW TOKENS ***
	for _, t := range tokens {
		if t.Type == EOL ||
			t.Type == COMMENT ||
			t.Type == NOP_OPCODE {
			continue
		}
		fmt.Print("<", tokenNames[t.Type], ">")
	}
	fmt.Print("\n")

	// *** TEST: PRETTY PRINT CONTENT ***
	indent_1 := "        "
	indent_2 := indent_1 + indent_1
	for _, t := range tokens {

		switch t.Type {
		case COMMA:
			fmt.Print(", ")
		case COMMENT:
			if strings.HasPrefix(t.Text, ";;") {
				fmt.Print(t.Text)
			} else {
				fmt.Print(" ", t.Text)
			}
		case LABEL:
			fmt.Print(t.Text, ":\n")
		case LOCAL_LABEL:
			fmt.Print(t.Text, ":\n")
		case EOL:
			fmt.Print("\n")
		case DIRECTIVE:
			fmt.Print(indent_1, t.Text, " ")
		case OPCODE:
			fmt.Print(indent_2, t.Text, " ")
		case BRANCH_OPCODE:
			fmt.Print(indent_2, t.Text, " ")
		case JUMP_OPCODE:
			fmt.Print(indent_2, t.Text, " ")
		default:
			fmt.Print(t.Text)
		}

	}
	fmt.Print("\n")

}
