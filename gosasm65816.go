// The GoAsm65816 Assember for the 65816 MPU
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 02. May 2018

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"goasm65816/lexer"
	"goasm65816/parser"
	"goasm65816/token"
)

var (
	input     = flag.String("i", "", "Input file (REQUIRED)")
	f_verbose = flag.Bool("v", false, "Give verbose messages")
	raw       []string
	tokens    []token.Token
)

// Verbose prints the given string if the verbose flag is set
func verbose(s string) {
	if *f_verbose {
		fmt.Println(s)
	}
}

func main() {

	flag.Parse()

	// *** LOAD SOURCE FILE ***

	if *input == "" {
		log.Fatal("FATAL: No input file provided.")
	}

	inputFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		raw = append(raw, scanner.Text())
	}

	// *** LEXER ***

	tokens, ok := lexer.Lexer(raw)
	if !ok {
		log.Fatal("FATAL: Lexer failed.")
	}

	verbose("Lexer run successful.")

	// *** PARSER ***

	ok := paser.Paser(tokens)
	if !ok {
		log.Fatal("FATAL: Parser failed.")
	}

}
