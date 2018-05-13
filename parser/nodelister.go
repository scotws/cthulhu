// Print a Lisp-like listing of the ast for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 13. May 2018
// This version: 13. May 2018

package parser

import (
	"fmt"

	"cthulhu/node"
	"cthulhu/token"
)

// Nodelister takes an ast from the parser and prints out a list of the tree
// elements in a Lisp-inspired S-format (ie, lots of braces). It is used for
// debugging.
func Nodelister(ast *node.Node) {

	fmt.Print(token.Name[ast.Type], " ")

	if ast.Type == token.EOL ||
		ast.Type == token.EMPTY ||
		ast.Type == token.START {
		fmt.Print("\n")
	}

	// If we don't have kids, we're done
	if len(ast.Kids) == 0 {
		return
	}

	fmt.Print("[ ")

	for _, k := range ast.Kids {
		Nodelister(k)
	}

	fmt.Print("]\n")
}
