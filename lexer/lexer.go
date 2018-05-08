// Lexer package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 08. May 2018

package lexer

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"goasm65816/data"
	"goasm65816/token"
)

var (
	tokens []token.Token

	// We can handle single-character tokens with this table and a loop
	singleChars = []rune{
		',', '-', '+', '@', '/', '*', '[', ']', '(', ')', '>', '<',
		'#', '{', '}',
	}
	singleCharTokens = []int{
		token.T_comma,
		token.T_minus,
		token.T_plus,
		token.T_anonLabel,
		token.T_slash,
		token.T_star,
		token.T_leftSquare,
		token.T_rightSquare,
		token.T_leftParens,
		token.T_rightParens,
		token.T_greater,
		token.T_less,
		token.T_hash,
		token.T_leftCurly,  // UNUSED
		token.T_rightCurly, // UNUSED
	}
)

// addToken takes the token identifier, the actual text of the token from the
// source code, the row and index the token was found in, and adds it to the
// token stream.
func addToken(ti int, s string, r int, i int) {
	s0 := strings.TrimSpace(s)
	r0 := r + 1 // computers count row from 0, we from one
	tokens = append(tokens, token.Token{ti, s0, r0, i})
}

// findBinEOW takes an array of runes and returns the index of the first
// non-binary number character (0 and 1). If there is none, it returns the
// length of the string
func findBinEOW(rs []rune) int {
	e := len(rs)

	for i := 0; i < len(rs); i++ {

		// Allow '.' and ':' as cosmetic separators in binary numbers.
		// They will be actually filtered out by the parser during
		// conversion
		if rs[i] == ':' || rs[i] == '.' {
			continue
		}

		if rs[i] != '0' && rs[i] != '1' {
			e = i
			break
		}
	}
	return e
}

// findDecEOW takes an array of runes and returns the index of the first
// non-decimal number character (0 to 9). If there is none, it returns the
// length of the string. In contrast to binary and hex numbers, we don't allow
// '.' and ':' as cosmetic separators
func findDecEOW(rs []rune) int {
	e := len(rs)

	for i := 0; i < len(rs); i++ {

		if !unicode.IsNumber(rs[i]) {
			e = i
			break
		}
	}
	return e
}

// findDirectiveEOW takes an array of runes and returns the index of the first
// character that is not a legal part of a directive. If there is none, it returns
// length of the string.
func findDirectiveEOW(rs []rune) int {
	e := len(rs)

	// Start one character in to skip '.'
	for i := 1; i < len(rs); i++ {

		if !unicode.IsNumber(rs[i]) &&
			!unicode.IsLetter(rs[i]) &&
			rs[i] != '!' {
			e = i
			break
		}
	}
	return e
}

// findHexEOW takes an array of runes and returns the index of the first
// non-hexadecimal character (0 to 9, a to f or A to F). If there is none, it
// returns the length of the string
func findHexEOW(rs []rune) int {
	var s string
	e := len(rs)

	for i := 0; i < len(rs); i++ {

		// Allow '.' and ':' as cosmetic separators in hex numbers, which
		// are used especially for the bank byte. They will be actually
		// filtered out by the parser during conversion
		if rs[i] == ':' || rs[i] == '.' {
			continue
		}

		s = string(rs[i])

		// We check if this is a legal character by running it through
		// the library routine
		_, err := strconv.ParseUint(s, 16, 64)

		if err != nil {
			e = i
			break
		}
	}
	return e
}

// findSANMneEOW takes an array of runes and returns the index of the
// the first rune that doesn't belong in a SAN mnemonic. If there is none,
// returns the length of the rune array
func findSANMneEOW(rs []rune) int {
	e := len(rs)

	// Start one character in to skip '.'
	for i := 1; i < len(rs); i++ {

		if !unicode.IsLetter(rs[i]) &&
			rs[i] != '.' &&
			rs[i] != '#' {
			e = i
			break
		}
	}
	return e
}

// findSymbolEOW takes an array of runes and returns the index of the first rune
// that doesn't belong in a label or a symbol. If there is none, it returns the
// length of the rune array
func findSymbolEOW(rs []rune) int {
	e := len(rs)

	// Start one character in to skip '_' and ':'
	for i := 1; i < len(rs); i++ {

		if !unicode.IsLetter(rs[i]) &&
			!unicode.IsNumber(rs[i]) &&
			!isLegalSymbolChar(rs[i]) {
			e = i
			break
		}
	}
	return e
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

// isCommentLine takes a string and checks to see if it is a full-line comment
func isCommentLine(s string) bool {
	f := false
	s0 := strings.TrimSpace(s)
	if s0[0] == ';' {
		f = true
	}

	return f
}

// isDirective takes a string and checks to see if it is a recognized
// directive
func isDirective(s string) bool {
	_, ok := data.Directives[s]
	return ok
}

// isEmptyLine take a complete line and checks to see if it is all whitespace
func isEmpty(s string) bool {
	f := false

	if len(strings.TrimSpace(s)) == 0 {
		f = true
	}

	return f
}

// isLegalSymbolChar takes a rune and returns a bool depending if it is a legal
// character for a symbol or a label. Note that this doesn't mean that a symbol
// or label may begin with these runes, they just are allowed after the
// initial character
func isLegalSymbolChar(r rune) bool {

	lsc := map[rune](bool){
		'?': true, '_': true, '!': true, '&': true, '\'': true,
		'~': true, '#': true, '|': true, '.': true, '=': true, '^': true,
	}
	f := false
	_, ok := lsc[r]

	if unicode.IsLetter(r) ||
		unicode.IsNumber(r) ||
		ok {
		f = true
	}
	return f
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

// proSANMne takes a list of runes and checks to see if it contains a legal
// mnemonic for the Simpler Assembler Notation (SAN). If yes, it returns a
// token, the index of the first rune past the mnemonic and a bool to designate
// success or failure
func procSANMne(rs []rune, mpu string) (int, int, bool) {

	var o int

	f := false

	e := findSANMneEOW(rs)
	r := rs[0:e]

	mt, ok := whichSANMne(r, mpu)

	if ok {
		f = true

		// SAN lets us already know which opcode we have and how many
		// operands it has just from the mnenomic, so we might as well
		// use that information
		switch mt {
		case 0:
			o = token.T_opcode0
		case 1:
			o = token.T_opcode1
		case 2:
			o = token.T_opcode2
		}
	}
	return o, e, f
}

// proWDCMne takes a list of runes and checks to see if it contains a legal
// mnemonic for the traditional WDC notation. If yes, it returns a token, the
// index of the first rune past the mnemonic and a bool to designate success or
// failure
func procWDCMne(rs []rune, mpu string) (int, int, bool) {
	var o int

	f := false

	// We allow uppercase letters for mnenomics because we are nice that
	// way, but internally all is lower case
	s0 := strings.ToLower(string(rs))

	// Any WDC opcode must be three characters long exactly
	if len(s0) >= 3 {
		s1 := s0[0:3]
		_, ok := data.OpcodesWDC[mpu][s1]

		if ok {
			o = token.T_opcodeWDC
			f = true
		}
	}
	return o, 3, f
}

// whichSANMne takes an array of runes and returns a int signaling the number
// of operands the SAN mnemonic takes (0, 1, or 2) and a flag if this is in fact
// a mnemonic. We can't do this yet for the WDC mnemonics because we'd have to
// figure out the operand which can be tricky
func whichSANMne(rs []rune, mpu string) (int, bool) {
	ok := false
	oc, ok := data.OpcodesSAN[mpu][string(rs)]
	return oc.Operands, ok
}

// Lexer takes a list of raw code lines and returns a list of tokens and a flag
// indicating if the conversion was successful or not. Error are handled by the
// main function.
func Lexer(ls []string, notation, mpu string) (*[]token.Token, bool) {

	ok := true

	// OUTER LOOP: Proceed line-by-line
	for ln, l := range ls {

		// Check for empty lines. We add an EOL token to allow
		// formatting
		if isEmpty(l) {
			addToken(token.T_eol, "\n", ln, 0)
			continue
		}

		if isCommentLine(l) {
			addToken(token.T_comment, l, ln, 0)
			addToken(token.T_eol, "\n", ln, 0)
			continue
		}

		// INNER LOOP: Proceed char-by-char
		cs := []rune(l)
		for i := 0; i < len(cs); i++ {

			// Skip over whitespace
			if unicode.IsSpace(cs[i]) {
				continue
			}

			// Single character tokenization (@ and friends).
			t, got := isSingleCharToken(cs[i])
			if got {
				addToken(t, string(cs[i]), ln, i)
				continue
			}

			// Detect numbers
			if unicode.IsNumber(cs[i]) {
				e := findDecEOW(cs[i:len(cs)])
				word := cs[i : i+e]
				addToken(token.T_decimal, string(word), ln, i)
				i = i + e - 1 // continue adds one
				continue
			}

			switch cs[i] {
			case ';':
				word := string(cs[i:len(cs)])
				addToken(token.T_comment, word, ln, i)
				i = len(cs)
				continue
			case '.':
				e := findDirectiveEOW(cs[i:len(cs)])
				word := string(cs[i : i+e])

				if isDirective(word) {
					addToken(token.T_directive, word, ln, i)
					i = i + e - 1 // continue adds one
				}
				continue

			// Global label. The first character after the colon
			// must be an uppercase or lowercase letter
			case ':':
				i += 1 // skip ':' symbol

				// First character after the colon must be an
				// upper- or lowercase letter because a label is
				// basically just a symbol
				if !unicode.IsLetter(cs[i]) {
					log.Fatal("LEXER FATAL (", ln+1, ":", i, "): First char after colon must be a letter")
				}
				e := findSymbolEOW(cs[i:len(cs)])
				word := cs[i-1 : i+e] // Include colon
				addToken(token.T_label, string(word), ln, i)
				i = i + e - 1 // continue adds one
				continue

			case '%':
				i += 1 // skip '%' symbol
				e := findBinEOW(cs[i:len(cs)])
				word := cs[i : i+e]
				addToken(token.T_binary, string(word), ln, i)
				i = i + e - 1 // continue adds one
				continue

			// Hex number. We allow uppercase and lowercase heximal
			// digits
			case '$':
				i += 1 // skip '$' symbol
				e := findHexEOW(cs[i:len(cs)])
				word := cs[i : i+e]
				addToken(token.T_hex, string(word), ln, i)
				i = i + e - 1 // continue adds one
				continue

			// Local label. First character after the underscore
			// must be a letter
			case '_':
				i += 1 // skip '_' symbol

				// First character after the underscore must be an
				// upper- or lowercase letter because a label is
				// basically just a symbol
				if !unicode.IsLetter(cs[i]) {
					log.Fatal("LEXER FATAL (", ln+1, ":", i, "): First char after underscore must be a letter")
				}

				e := findSymbolEOW(cs[i:len(cs)])
				word := cs[i-1 : i+e] // Include underscore
				addToken(token.T_localLabel, string(word), ln, i)
				i = i + e - 1 // continue adds one
				continue

			// String. Use double quote instead of single quote.
			// Note we currently don't allow backslashes to get the
			// quotation mark itself
			case '"':
				i += 1 // skip leading quote
				e, ok := findStringEOW(cs[i:len(cs)])

				if !ok {
					log.Fatal("LEXER FATAL (", ln+1, ":", i, "): No closing quote")
				}

				word := cs[i : i+e]
				addToken(token.T_string, string(word), ln, i)
				i = i + e // skip over final quote
				continue
			}

			// See if we are dealing with a mnemonic. This step is
			// more complicated because we accept more than one
			// notation
			if unicode.IsLetter(cs[i]) {

				var e, tt int

				// WDC and SAN give us totally different types
				// of tokens because SAN allows us to immediatey
				// say how many operands an instruction has.
				// This means we let the specialized routines
				// take care of the token
				switch notation {

				case "wdc":
					tt, e, ok = procWDCMne(cs[i:len(cs)], mpu)
				case "san":
					tt, e, ok = procSANMne(cs[i:len(cs)], mpu)
				default:
					log.Fatalf("LEXER FATAL: Received illegal notation '%s'\n")
				}

				if ok {
					addToken(tt, string(cs[i:i+e]), ln, i)
					i = i + e - 1 // continue adds one
				} else {

					// If this is not some opcode and none
					// of the above, it has to be some sort
					// of a symbol
					e = findSymbolEOW(cs[i:len(cs)])
					word := cs[i : i+e]
					addToken(token.T_symbol, string(word), ln, i)
					i = i + e - 1 // continue adds one
				}
				continue
			}
		}
		addToken(token.T_eol, "\n", ln, len(cs))
	}

	addToken(token.T_eof, "That's all, folks!", len(ls), 0)

	return &tokens, ok
}
