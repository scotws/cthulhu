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
	//	"unicode"
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

func printToken(t token) {
	fmt.Print("<", tokenNames[t.Type], ">")
}

func main() {
	flag.Parse()

	tokenNames[EOL] = "EOL"

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

		// At the end of the line, add a EOL token
		fmt.Println(line_number, "-->", line)
		tokens = append(tokens, token{EOL, ""})
	}

	// *** TEST: PRINT TOKENS ***
	for _, t := range tokens {
		printToken(t)
	}
}
