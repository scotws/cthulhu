// Print the nodes modified by the analyzer
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 12. May 2018
// First version: 12. May 2018

package analyzer

import (
	"fmt"

	"cthulhu/node"
	"cthulhu/token"
)

// Worklister takes an AST from the parser and prints out the list of nodes that
// were modified
func Worklister(aAst *node.Node) {

	if aAst.Modified == true {

		switch aAst.Type {

		case token.DEC_NUM, token.HEX_NUM, token.BIN_NUM:
			fmt.Printf("Node (%d,%d) was '%s', now is %s with Value %d\n",
				aAst.Line, aAst.Index, aAst.Text, token.Name[aAst.Type], aAst.Value)

		default:
			fmt.Printf("Node (%d,%d) otherwise modified (removed subnode?), now is '%s'\n",
				aAst.Line, aAst.Index, aAst.Text)
		}
	}

	// If we don't have kids, we're done
	if len(aAst.Kids) == 0 {
		return
	}

	for _, k := range aAst.Kids {
		Worklister(k)
	}
}
