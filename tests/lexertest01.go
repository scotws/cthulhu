// Lexer tests for GoAsm65816
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 21. April 2018
// This version: 21. April 2018

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type token struct {
	Type int
	Text string
}

const (
	EOF          int = iota
	EOL              // included for output formatting
	AMPERSAND        // '&'
	ANON_LABEL       // '@'
	COMMA            // ','
	COMMENT          // starts with ';'
	DIRECT_NAKED     // eg '.native'
	DIRECT_PARA      // eg '.byte' (with parameters)
	DOLLAR           // '$'
	LABEL            // eg 'loop'
	MINUS            // '-'
	OPCODE_0         // eg 'nop' (no operand)
	OPCODE_1         // eg 'lda.# 00' (one operand)
	OPCODE_2         // eg 'mvp 00,01' (two operands, move commands)
	PLUS             // '+'
	QUOTE            // '"'
	SEMICOLON        // ';'
	SLASH            // '/'
	STAR             // '*'
	WS               // all whitespace
)

var (
	opcodes_0     = []string{"nop", "inx", "pha"}
	opcodes_1     = []string{"lda.#", "sta.d", "bra"}
	opcodes_2     = []string{"mvp"}
	directs_naked = []string{".native"}
	directs_para  = []string{".byte", ".origin", ".equ", ".mpu"}

	input = flag.String("i", "", "Input file (REQUIRED)")

	raw    []string
	tokens []token

	tokenNames = make(map[int]string)
)

// printToken is a temporary routine for testing
func printToken(t token) {
	fmt.Print("<", tokenNames[t.Type], "> ", t.Text, "\n")
}

func main() {
	flag.Parse()

	tokenNames[EOL] = "EOL"
	tokenNames[WS] = "WS"
	tokenNames[LABEL] = "LABEL"
	tokenNames[COMMENT] = "COMMENT"
	tokenNames[ANON_LABEL] = "ANON_LABEL"

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

	for line_number, line := range raw {

		// Deal with empty lines
		if len(line) == 0 || strings.TrimSpace(line) == "" {
			tokens = append(tokens, token{EOL, ""})
			continue
		}

		// Move stuff to runes because we want to be able
		// to deal with Unicode
		chars := []rune(line)

		// Do stuff with first character
		c0 := chars[0]
		if !unicode.IsSpace(c0) {
			switch c0 {

			case ';':
				tokens = append(tokens, token{COMMENT, line})
				tokens = append(tokens, token{EOL, ""})
				continue

			case '@':
				tokens = append(tokens, token{ANON_LABEL, "@"})
				line = line[1:len(line)]

			// This must then be a label
			default:
				w0 := strings.SplitN(line, " ", 2)
				tokens = append(tokens, token{LABEL, w0[0]})

				// If there was just the label in the line,
				// we're done
				if len(w0) == 1 {
					tokens = append(tokens, token{EOL, ""})
					continue
				}

				// Otherwise, move past the label and continue
				line = w0[1]
			}
		}

		// HIER HIER HIER TODO

		// At the end of the line, add a EOL token
		fmt.Println(line_number, "-->", line)
		tokens = append(tokens, token{EOL, ""})
	}

	// *** TEST: PRINT TOKENS ***
	for _, t := range tokens {
		printToken(t)
	}
}
