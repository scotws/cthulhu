// The goasm65 Assember
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 07. May 2018

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"goasm65816/formatter"
	"goasm65816/lexer"
	"goasm65816/lister"
	"goasm65816/parser"
	"goasm65816/token"
)

var (
	f_debut   = flag.Bool("d", false, "Print lots and lots of debugging information")
	input     = flag.String("i", "", "Input file (REQUIRED)")
	f_format  = flag.Bool("f", false, "Return formatted version of source")
	f_verbose = flag.Bool("v", false, "Give verbose messages")
	f_listing = flag.Bool("l", false, "Generate listing file")
	mpu       = flag.String("m", "65c02", "MPU (default 65c02)")
	notation  = flag.String("n", "wdc", "Assembler notation (wdc or san)")
	f_symbols = flag.Bool("s", false, "Generate symbol table file")

	raw    []string
	tokens []token.Token
)

// Verbose prints the given string if the verbose flag is set
func verbose(s string) {
	if *f_verbose {
		fmt.Println(s)
	}
}

func main() {

	flag.Parse()

	if *notation != "wdc" && *notation != "san" {
		log.Fatalf("FATAL: Notation '%s' not supported", *notation)
	}

	if *mpu != "6502" && *mpu != "65c02" && *mpu != "65816" {
		log.Fatalf("FATAL: MPU '%s' not supported", *mpu)
	}

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

	// The lexer doesn't care about the MPU we are using, just about
	// the notation
	tokens, ok := lexer.Lexer(raw, *notation, *mpu)
	if !ok {
		log.Fatal("FATAL: Lexer failed.")
	}

	verbose("Lexer run successful.")

	// *** FORMATTER ***

	if *f_format {
		formatter.Formatter(tokens)
	}

	verbose("Formatter run successful.")

	// *** PARSER ***

	ok = parser.Parser(tokens)
	if !ok {
		log.Fatal("FATAL: Parser failed.")
	}

	verbose("Parser run successful.")

	// *** LISTER ***

	if *f_listing {
		lister.Lister()
	}

	verbose("Lister run successful.")
}
