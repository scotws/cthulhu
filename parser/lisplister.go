// Print a Lisp-like listing of the AST for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 08. May 2018
// This version: 12. May 2018

package parser

import (
	"fmt"

	"cthulhu/data"
	"cthulhu/node"
	"cthulhu/token"
)

// Lisplister takes an AST from the parser and prints out a list of the tree
// elements in a Lisp-inspired S-format (ie, lots of braces). It is used for
// debugging.
func Lisplister(AST *node.Node) {

	switch AST.Type {

	// Special case
	case token.EOL:
		fmt.Print("<EOL> )\n")
	case token.EOF:
		fmt.Print("( <EOF> )\n")
	case token.EMPTY:
		fmt.Print("( <EMPTY> )\n")
	case token.START:
		fmt.Print(" ( ", AST.Text, " )", "\n")
	case token.OPC_0, token.OPC_1:
		fmt.Print("( ", AST.Text)

	// Some of the directors are actually operators that don't start a
	// a new line
	case token.DIREC, token.DIREC_PARA:
		_, ok := data.Operators[AST.Text]

		if ok {
			fmt.Print(AST.Text)
		} else {
			fmt.Print("( ", AST.Text)
		}

	// Comments come in two forms, at the beginning of a line or at the end
	// of a line.
	case token.COMMENT_LINE:
		fmt.Print("( ", AST.Text)
	case token.COMMENT:
		fmt.Print(") ( ", AST.Text)
	case token.HEX_NUM:
		fmt.Print("$", AST.Text)
	case token.BIN_NUM:
		fmt.Print("%", AST.Text)
	case token.STRING:

		if AST.Done {
			fmt.Print(node.FormatByteSlice(AST.Code))
		} else {

			fmt.Print("\"", AST.Text, "\"")
		}

	// TODO missing closing parens if label is not alone in the line
	case token.LABEL, token.LOCAL_LABEL:
		fmt.Print("( ", AST.Text, ":")

	case token.ANON_LABEL:
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
