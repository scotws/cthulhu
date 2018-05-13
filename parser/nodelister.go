// Display structure of the AST
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 13. May 2018
// This version: 13. May 2018

package parser

import (
	"fmt"
	"strings"

	"cthulhu/node"
	"cthulhu/token"
)

var indent int

// Nodelister takes an ast from the parser and prints out an indented display
// of the AST. It is used for debugging.
func Nodelister(ast *node.Node) {

	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print(token.Name[ast.Type])

	if ast.Type != token.EOL &&
		ast.Type != token.EMPTY {

		// TODO clean this up with a switch once we have everything
		// working

		if ast.Type == token.DEC_NUM {
			fmt.Printf(" (%d) ", ast.Value)
		} else {
			fmt.Print(" (", ast.Text, ")")
		}

		if ast.Done {
			fmt.Print(" -> ", node.FormatByteSlice(ast.Code))
		}
	}

	fmt.Print("\n")

	// If we don't have kids, we're done
	if len(ast.Kids) == 0 {
		return
	}

	indent++

	for _, k := range ast.Kids {
		Nodelister(k)
	}

	indent--
}
