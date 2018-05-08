// Print a Lisp-like listing of the AST
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 08. May 2018
// First version: 08. May 2018

package parser

import (
	"fmt"

	"goasm65816/node"
	"goasm65816/token"
)

// lisplister takes an AST from the parser and prints out a list of the tree
// elements in a Lisp-inspired S-format (ie, lots of braces). It is used for
// debugging.

const indent = "       "

func Lisplister(AST *node.Node) {

	// An EOL token means we're done with this element
	if AST.Type == token.T_eol {
		fmt.Print("\n")
		return
	}

	fmt.Print(" ")

	switch AST.Type {

	// Special case
	case token.T_start:
		fmt.Print(AST.Text, "\n")
	case token.T_opcodeWDC, token.T_opcode0, token.T_opcode1:
		fmt.Print("(", AST.Text)
	case token.T_directive, token.T_comment:
		fmt.Print("(", AST.Text)
	case token.T_hex:
		fmt.Print("$", AST.Text)
	case token.T_binary:
		fmt.Print("%", AST.Text)
	case token.T_label, token.T_localLabel, token.T_anonLabel:
		fmt.Print("(", AST.Text)
	default:
		fmt.Print(AST.Text)
	}

	// If we don't have kids, we're done
	if len(AST.Kids) == 0 {
		fmt.Print(")")
		return
	}

	for _, k := range AST.Kids {
		Lisplister(k)

	}
}