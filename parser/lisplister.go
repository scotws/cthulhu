// Print a Lisp-like listing of the ast for the Cthulhu Assembler
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

// Lisplister takes an ast from the parser and prints out a list of the tree
// elements in a Lisp-inspired S-format (ie, lots of braces). It is used for
// debugging.
func Lisplister(ast *node.Node) {

	switch ast.Type {

	// Special case
	case token.EOL:
		fmt.Print("<EOL> )\n")
	case token.EOF:
		fmt.Print("( <EOF> )\n")
	case token.EMPTY:
		fmt.Print("( <EMPTY> )\n")
	case token.START:
		fmt.Print(" ( ", ast.Text, " )", "\n")
	case token.OPC_0, token.OPC_1:
		fmt.Print("( ", ast.Text)

	// Some of the directors are actually operators that don't start a
	// a new line
	case token.DIREC, token.DIREC_PARA:
		_, ok := data.Operators[ast.Text]

		if ok {
			fmt.Print(ast.Text)
		} else {
			fmt.Print("( ", ast.Text)
		}

	// Comments come in two forms, at the beginning of a line or at the end
	// of a line.
	case token.COMMENT_LINE:
		fmt.Print("( ", ast.Text)
	case token.COMMENT:
		fmt.Print(") ( ", ast.Text)

	// HEX_NUM are converted to DEC_NUM by the analyzer, so if we have a
	// HEX_NUM, we haven't converted it yet. Same for BIN_NUM
	case token.HEX_NUM:
		fmt.Print("$", ast.Text)
	case token.BIN_NUM:
		fmt.Print("%", ast.Text)

	case token.DEC_NUM:

		if ast.Done {
			fmt.Print(node.FormatByteSlice(ast.Code))
		} else {
			fmt.Print(ast.Text)
		}

	case token.STRING:

		if ast.Done {
			fmt.Print(node.FormatByteSlice(ast.Code))
		} else {

			fmt.Print("\"", ast.Text, "\"")
		}

	// TODO missing closing parens if label is not alone in the line
	case token.LABEL, token.LOCAL_LABEL:
		fmt.Print("( ", ast.Text, ":")

	case token.ANON_LABEL:
		fmt.Print("( ", ast.Text)
	default:
		fmt.Print(ast.Text)
	}

	// If we don't have kids, we're done
	if len(ast.Kids) == 0 {
		return
	}

	for _, k := range ast.Kids {
		fmt.Print(" ")
		Lisplister(k)
	}
}
