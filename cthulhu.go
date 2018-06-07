// The Cthulhu Assember for the 6502/65c02/65816
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 07. June 2018

package main

import (
	"flag"
	"fmt"
	"log"

	"cthulhu/analyzer"
	"cthulhu/data"
	//	"cthulhu/formatter"
	"cthulhu/generator"
	"cthulhu/lexer"
	"cthulhu/lister"
	"cthulhu/parser"
	"cthulhu/token"
)

var (
	fDebug      = flag.Bool("d", false, "Print lots and lots of debugging information")
	fInput      = flag.String("i", "", "Input file (REQUIRED)")
	fFormat     = flag.Bool("f", false, "Return formatted version of source")
	fFormatFile = flag.String("ff", "", "File name to save formatted source")
	fFormatOnly = flag.Bool("fo", false, "Only produce formatted source code")
	fHexdump    = flag.Bool("h", false, "Add hexdump of binary in text file \"cthulhu.hex\"")
	fVerbose    = flag.Bool("v", false, "Give verbose messages")
	fListing    = flag.Bool("l", false, "Generate listing file")
	mpu         = flag.String("m", "65c02", "MPU type")
	fSymbols    = flag.Bool("s", false, "Generate symbol table file")

	tokens []token.Token
)

// Verbose prints the given string if the verbose flag is set
func verbose(s string) {
	if *fVerbose {
		fmt.Println(s)
	}
}

func main() {

	// ***** CONFIRM REQUIRED FLAGS *****

	flag.Parse()

	if *mpu != "6502" && *mpu != "65c02" && *mpu != "65816" {
		log.Fatalf("FATAL MPU '%s' not supported", *mpu)
	}
	if *fInput == "" {
		log.Fatal("FATAL No input file provided")
	}

	// ***** LEXER *****

	// The lexer takes the name of the master source file and the mpu and
	// splits up the source into tokens.

	v := fmt.Sprintf("LEXER: Scanning %s as main source file", *fInput)
	verbose(v)

	tokens := lexer.Lexer(*mpu, *fInput)

	// Add a single end-of-file (EOF) token. Lexer can't do this itself
	// because it can call itself for the include files, and that would put
	// various EOF tokens all over the code
	*tokens = append(*tokens, token.Token{token.EOF, len(*tokens), 0, *fInput, "EOF"})

	// Part of the debugging information is a list of tokens
	if *fDebug {
		fmt.Println("=== List of tokens after initial lexing ===\n")
		lexer.Tokenlister(tokens)
	}

	verbose("Lexer done.")

	// ***** PARSER *****

	// The parser takes a slice of tokens and returns an Abstract Syntax
	// Tree (AST) built of node.Node elements. This AST is used as the basis
	// for all other work.
	parser.Init(tokens)
	ast := parser.Parser()

	// Part of the debugging information is an indented list of nodes of the
	// AST
	if *fDebug {
		fmt.Println("=== AST after initial parsing: ===\n")
		parser.Nodelister(ast)
	}

	verbose("Parser done.")

	// ***** FORMATTER *****

	// The formatter produces a cleanly indented version of the source code,
	// much like the gofmt program included with Go. See the file itself for
	// more detail

	// TODO move from token to AST based formatting
	// TODO Since formatting is not related to the other parsing steps,
	//      we should be able to do this concurrently

	/*
		if *fFormat {
			formatter.Formatter(ast)
			verbose("Formatter run.")
		}
	*/

	// *** CONSTRUCT THE MACHINE ***

	// We are now at the point where we can construct a machine to hold the
	// greater values
	// TODO see if we need this
	machine := data.Machine{MPU: *mpu, AST: ast}

	// *** ANALYZER ***

	// The analyzer doesn't modify the AST, because that is used by the
	// formatter and other parts. Instead, it creates a second syntax tree,
	// the "BST" (because it comes after the AST) and modifies that step by
	// step. PURGE handles the initial creation of the BST. As the name says, it
	// deletes whitespace, EOL nodes, and does some easy processing of other
	// nodes
	bst := analyzer.Purge(*mpu, ast)

	if *fDebug {
		fmt.Println("=== BST after purge step: ===\n")
		parser.Nodelister(bst)
	}

	// The analyzer examens the AST provided by the parser and runs various
	// processes on it to convert numbers, etc.
	// TODO see about passing out symbol table(s)
	analyzer.Analyzer(&machine)

	if *fDebug {
		fmt.Println("=== Completed nodes after analyzer ===")
		fmt.Println()
		analyzer.Worklister(machine.AST)
		fmt.Println()
	}

	if *fDebug {
		fmt.Println("=== AST after analyzer ===")
		fmt.Println()
		parser.Nodelister(machine.AST)
		fmt.Println()
	}

	// *** GENERATOR ***

	// The generator takes the assembler instructions and other information
	// and produces the actual bytes that will be saved in the final file.
	// TODO

	generator.Generator(machine.AST)

	// *** LISTER ***

	// The lister produces a detailed listing of the code with useful
	// information such as the actual byte stored for each instruction and
	// the modes the 65816 was in during each instruction
	// TODO Since this is based on the AST, we should be able to do this
	//      concurrently

	if *fListing {
		lister.Lister(machine.AST)
		verbose("Lister run.")
	}

	// *** HEXDUMP ***

	// TODO Since this is based on the AST, we should be able to do this
	//      concurrently
	if *fHexdump {
		verbose("(Hexdump not installed yet)")
	}

}
