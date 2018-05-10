// The Cthulhu Assember for the 6502/65c02/65816
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 10. May 2018

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"cthulhu/formatter"
	"cthulhu/lexer"
	"cthulhu/lister"
	"cthulhu/parser"
	"cthulhu/token"
)

var (
	f_debug   = flag.Bool("d", false, "Print lots and lots of debugging information")
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
		log.Fatalf("FATAL Notation '%s' not supported", *notation)
	}

	if *mpu != "6502" && *mpu != "65c02" && *mpu != "65816" {
		log.Fatalf("FATAL MPU '%s' not supported", *mpu)
	}

	// ***** LOAD SOURCE FILE *****

	if *input == "" {
		log.Fatal("FATAL No input file provided")
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

	verbose("Source file loaded.")

	// ***** LEXER *****

	tokens := lexer.Lexer(raw, *notation, *mpu)

	// TODO include .INCLUDE files and lex them
	// Remember to add information on the file in the tokens for error
	// purposes
	verbose("(Includer not written yet.)")

	// Part of the debugging information is a list of tokens
	if *f_debug {
		fmt.Println("=== List of tokens after initial lexing: ===\n")
		lexer.Tokenlister(tokens)
	}

	verbose("Lexer run.")

	// ***** FORMATTER *****

	// The formatter produces a cleanly indented version of the source code,
	// much like the gofmt program included with Go. See the file itself for
	// more detail

	if *f_format {
		formatter.Formatter(tokens)
		verbose("Formatter run.")
	}

	// ***** PARSER *****

	// The parser takes a slice of tokens and returns an Abstract Syntax
	// Tree (AST) built of node.Node elements. This AST is used as the basis
	// for all other work
	AST := parser.Parser(tokens)

	// Part of the debugging information is a Lisp-like list of elements of
	// the AST
	if *f_debug {
		fmt.Println("=== AST after initial parsing: ===\n")
		parser.Lisplister(&AST)
	}

	verbose("Parser run.")

	// *** ANALYZER ***

	// The analyzer examens the AST provided by the parser and runs various
	// processes on it to convert numbers,
	// TODO

	verbose("(Analyzer not written yet.)")

	// TODO print Lisp tree if debugging enabled

	// *** GENERATOR ***

	// The generator takes the assembler instructions and other information
	// and produces the actual bytes that will be saved in the final file.
	// TODO

	verbose("(Generator not written yet.)")

	// TODO print Lisp tree if debugging enabled

	// *** LISTER ***

	// The lister produces a detailed listing of the code with useful
	// information such as the actual byte stored for each instruction and
	// the modes the 65816 was in during each instruction

	if *f_listing {
		lister.Lister(AST)
		verbose("Lister run.")
	}

}
