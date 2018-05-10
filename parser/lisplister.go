// Print a Lisp-like listing of the AST for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 08. May 2018
// First version: 10. May 2018

package parser

import (
	"fmt"

	"cthulhu/node"
	"cthulhu/token"
)

// lisplister takes an AST from the parser and prints out a list of the tree
// elements in a Lisp-inspired S-format (ie, lots of braces). It is used for
// debugging.

func Lisplister(AST *node.Node) {

	switch AST.Type {

	// Special case
	case token.EOL:
		fmt.Print(")\n")
	case token.EMPTY:
		fmt.Print("\n")
	case token.START:
		fmt.Print(AST.Text, "\n")
	case token.WDC, token.SAN_0, token.SAN_1:
		fmt.Print("( ", AST.Text)
	case token.DIREC, token.DIREC_PARA, token.COMMENT:
		fmt.Print("( ", AST.Text)
	case token.HEX_NUM:
		fmt.Print("$", AST.Text)
	case token.BIN_NUM:
		fmt.Print("%", AST.Text)
	case token.STRING:
		fmt.Print("\"", AST.Text, "\"")
	case token.LABEL, token.LOCAL_LABEL, token.ANON_LABEL:
		fmt.Print("( ", AST.Text)
	default:
		fmt.Print(AST.Text)
	}

	// If we don't have kids, we're done
	if len(AST.Kids) == 0 {
		return
	}

	for _, k := range AST.Kids {
		fmt.Print(" ")
		Lisplister(k)
	}
}
