// Lexer package for the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 07. June 2018

package lexer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"cthulhu/data"
	"cthulhu/token"
)

const (
	errTag = "LEXER"
)

var (
	errCount int

	// We can handle single-character tokens with this table and a loop.
	// DOLLAR ('$') is not included currently because it would screw up the
	// lexer's hex number detection; same for PERIOD ('.') because it
	// doesn't work with directive detection
	singleChars = []rune{
		',', '-', '+', '@', '/', '*', '[', ']', '(', ')', '>', '<',
		'#', '{', '}', '=', '&', '|', '~', '^',
	}
	singleCharTokens = []int{
		token.COMMA,
		token.MINUS,
		token.PLUS,
		token.ANON_LABEL,
		token.SLASH,
		token.STAR,
		token.L_SQUARE,
		token.R_SQUARE,
		token.L_PARENS,
		token.R_PARENS,
		token.GREATER,
		token.LESS,
		token.HASH,
		token.L_CURLY,
		token.R_CURLY,
		token.EQUAL,
		token.AMPERSAND,
		token.PIPE,
		token.PERCENT,
		token.TILDE,
	}
)

// reportErr takes a string, the name of the current file, the current line and
// character index and prints an error report to the standard error output
func reportErr(s string, fn string, l, i int) {
	fmt.Fprintf(os.Stderr, "%s ERROR (%s, %d, %d): %s\n",
		errTag, fn, l+1, i+1, s)
	errCount++
}

// addToken takes the token identifier, the actual text of the token from the
// source code, the row and index the token was found in, and adds it to the
// token stream. We have to pass the pointer to the slice of tokens because Go
// doesn't allow nested functions
func addToken(tokens *[]token.Token, ti int, s string, l int, i int, f string) {
	s0 := strings.TrimSpace(s)
	l0 := l + 1 // computers count line from 0, humans from 1
	i0 := i + 1 // computers count from column 0, humans from 1
	*tokens = append(*tokens, token.Token{Type: ti, Text: s0, Line: l0, Index: i0, File: f})
}

// findBinEOW takes an array of runes and returns the index of the first
// non-binary number character (0 and 1). If there is none, it returns the
// length of the string. Note any ':' and '.' count when calculating the length
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
			rs[i] != '.' &&
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

// findMneEOW takes an array of runes and returns the index of the
// the first rune that doesn't belong in a SAN mnemonic. If there is none,
// returns the length of the rune array
func findMneEOW(rs []rune) int {
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

	for i := 0; i < len(rs); i++ {

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
// quote mark was found, return a zero as int and a false bool. This can also be
// used to find the beginning of a string
func findStringEOW(rs []rune) (int, bool) {
	f := false
	t := 0
	for i, c := range rs {
		if c == '"' {
			t = i // don't include the quote itself
			f = true
			break
		}
	}
	return t, f
}

// getIncludeFile takes an array of runes and returns a string, the number of
// characters we've skipped, and a flag to indicated if looking for the string
// was a success
func getIncludeFile(rs []rune) (string, int, bool) {

	start, ok := findStringEOW(rs) // get starting quote
	if !ok {
		return ".include needs filename string", 0, false
	}

	// We are now pointing to the character before the quotation mark, so
	// we move up one and then try to find the end
	start++

	end, ok := findStringEOW(rs[start:len(rs)])
	if !ok {
		return ".include needs filename string", 0, false
	}

	fn := rs[start : start+end]

	return string(fn), start + end, true
}

// isCommentLine takes a string and checks to see if it is a comment by getting
// rid of the whitespace around it
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
// character for a symbol or a label. Note that ':' is not a legal symbol
// character, but a flag marking the definition of the label
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

// procMne takes a list of runes and checks to see if it contains a legal
// mnemonic for the Simpler Assembler Notation (SAN). If yes, it returns a
// token, the index of the first rune past the mnemonic and a bool to designate
// success or failure
func procMne(rs []rune, mpu string) (int, int, bool) {

	var o int

	f := false

	e := findMneEOW(rs)
	r := rs[0:e]

	mt, ok := whichMne(r, mpu)

	if ok {
		f = true

		// SAN lets us already know which opcode we have and how many
		// operands it has just from the mnenomic, so we might as well
		// use that information
		switch mt {

		case 0:
			o = token.OPC_0
		case 1:
			o = token.OPC_1
		case 2:
			o = token.OPC_2
		}
	}
	return o, e, f
}

// whichMne takes an array of runes and returns a int signaling the number
// of operands the SAN mnemonic takes (0, 1, or 2) and a flag if this is in fact
// a mnemonic.
func whichMne(rs []rune, mpu string) (int, bool) {
	ok := false
	oc, ok := data.OpcodesSAN[mpu][string(rs)]
	return oc.Operands, ok
}

// Lexer takes the mpu type and the name of a file to scan and returns a list of
// tokens. If there are .include files in the mix, it will call itself
func Lexer(mpu string, filename string) *[]token.Token {

	var tokens []token.Token
	var ls []string

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ls = append(ls, scanner.Text())
	}

	// OUTER LOOP: Proceed line-by-line

	for ln, l := range ls {

		// Check for empty lines. We add a token to allow
		// formatting
		if isEmpty(l) {
			addToken(&tokens, token.EMPTY, "", ln, 1, filename)
			addToken(&tokens, token.EOL, "\n", ln, 1, filename)
			continue
		}

		// See if this is a whole-line comment, which gets a different
		// token than the comments in-line
		if isCommentLine(l) {
			addToken(&tokens, token.COMMENT_LINE, l, ln, 1, filename)
			addToken(&tokens, token.EOL, "\n", ln, 1, filename)
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
				addToken(&tokens, t, string(cs[i]), ln, i, filename)
				continue
			}

			// Detect numbers
			if unicode.IsNumber(cs[i]) {
				e := findDecEOW(cs[i:len(cs)])
				word := cs[i : i+e]
				addToken(&tokens, token.DEC_NUM, string(word), ln, i, filename)
				i = i + e - 1 // continue adds one
				continue
			}

			// Now things get a little tricker, because we have to
			// isolate words

			switch cs[i] {

			// In-line comments. Always run till the end of the line
			case ';':
				word := string(cs[i:len(cs)])
				addToken(&tokens, token.COMMENT, word, ln, i, filename)
				i = len(cs)
				continue

			// Directives
			case '.':
				e := findDirectiveEOW(cs[i:len(cs)])
				word := string(cs[i : i+e])

				if isDirective(word) {

					// ellipsis is a special case because it
					// can be used for various things such
					// as listings of bytes
					if word == "..." {
						addToken(&tokens, token.ELLIPSIS, word, ln, i, filename)
						i = i + e - 1 // continue adds one
						continue
					}

					// .include is a special case because it
					// causes us to call ourselves and parse
					// the file we were given
					if word == ".include" {

						fn, ef, ok := getIncludeFile(cs[i:len(cs)])
						if !ok {
							es := "Error getting .include file"
							reportErr(es, fn, ln, i)
							continue
						}

						inclTokens := Lexer(mpu, fn)

						// Remove the last token, which
						// is an EOF
						*inclTokens = (*inclTokens)[0 : len(*inclTokens)-1]

						tokens = append(tokens, *inclTokens...)

						i = i + ef
						continue
					}

					// We make life easier for the parser
					// by distinguishing between simple
					// directives (default) and those with
					// parameters
					_, ok := data.DirectivesPara[word]

					if ok {
						addToken(&tokens, token.DIREC_PARA, word, ln, i, filename)
					} else {

						addToken(&tokens, token.DIREC, word, ln, i, filename)
					}

					i = i + e - 1 // continue adds one
					continue
				}

				es := fmt.Sprintf("Unknown directive '%s'", word)
				reportErr(es, filename, ln, i)
				continue

			// Binary number
			case '%':
				i++ // skip '%' symbol
				e := findBinEOW(cs[i:len(cs)])
				word := cs[i : i+e]
				addToken(&tokens, token.BIN_NUM, string(word), ln, i, filename)
				i = i + e - 1 // continue adds one
				continue

			// Hex number. We allow uppercase and lowercase heximal
			// digits
			case '$':
				i++ // skip '$' symbol
				e := findHexEOW(cs[i:len(cs)])
				word := cs[i : i+e]
				addToken(&tokens, token.HEX_NUM, string(word), ln, i, filename)
				i = i + e - 1 // continue adds one
				continue

			// Local label or symbol. First character after the underscore
			// must be a letter. There is currently code duplication
			// here with the detection of global labels
			case '_':
				i++ // skip '_' symbol

				// First character after the underscore must be an
				// upper- or lowercase letter because a label is
				// basically just a symbol
				if !unicode.IsLetter(cs[i]) {
					reportErr("Letter required after initial local label underscore",
						filename, ln, i)
				}

				e := findSymbolEOW(cs[i:len(cs)])
				word := cs[i-1 : i+e] // Include underscore

				// if the next character is a colon (':'), we've
				// definied a local label here, otherwise it's a
				// scoped symbol

				if i+e < len(cs) {
					nc := cs[i+e]

					if nc == ':' {
						addToken(&tokens, token.LOCAL_LABEL, string(word), ln, i, filename)
						i = i + e // continue adds one, but skip colon
						continue
					}
				}

				// Not a local label, just some sort of symbol
				addToken(&tokens, token.SYMBOL, string(word), ln, i, filename)
				i = i + e - 1 // continue adds one
				continue

			// String. Use double quote instead of single quote.
			// Note we currently don't allow backslashes to get the
			// quote character itself
			case '"':
				i++ // skip leading quote
				e, ok := findStringEOW(cs[i:len(cs)])
				if !ok {
					reportErr("Can't find closing quotation mark", filename, ln, i)
					continue
				}

				word := cs[i : i+e]
				addToken(&tokens, token.STRING, string(word), ln, i, filename)
				i = i + e // skip over final quote
				continue
			}

			// Our word starts with a letter, which means it can be
			// an instruction, a label definition, or just a normal
			// symbol
			if unicode.IsLetter(cs[i]) {

				var e, tt int
				var ok bool

				// We take some work off the parser by getting
				// the information about what kind of opcode we
				// have -- one with zero, one, or two (65816
				// only) operands
				tt, e, ok = procMne(cs[i:len(cs)], mpu)

				if ok {
					addToken(&tokens, tt, string(cs[i:i+e]), ln, i, filename)
					i = i + e - 1 // continue adds one
					continue
				}

				// We're dealing with some sort of symbol.
				e = findSymbolEOW(cs[i:len(cs)])
				word := cs[i : i+e]

				// if the next character is a colon (':'), we've
				// definied a global label here
				if i+e < len(cs) {
					nc := cs[i+e]

					if nc == ':' {
						addToken(&tokens, token.LABEL, string(word), ln, i, filename)
						i = i + e // continue adds one, but skip colon
						continue
					}
				}

				// This is just a symbol then
				addToken(&tokens, token.SYMBOL, string(word), ln, i, filename)
				i = i + e - 1 // continue adds one
				continue
			}

			// We don't believe in illegal tokens. If we got here,
			// we report it immediately and attempt to continue
			es := fmt.Sprintf("Can't process character '%s'", string(cs[i]))
			reportErr(es, filename, ln, i)
			continue

		}
		addToken(&tokens, token.EOL, "\n", ln, len(cs), filename)
	}

	// Add an end of file token (EOF). If this is an include file, this will
	// be deleted by the call and only the main one will remain
	addToken(&tokens, token.EOF, "EOF", len(ls), 0, filename)

	if errCount != 0 {
		log.Fatalf("LEXER FATAL: Found %d error(s).", errCount)
	}

	return &tokens
}
