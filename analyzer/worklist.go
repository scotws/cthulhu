// Print the nodes modified by the analyzer
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 12. May 2018
// First version: 12. May 2018

// Worklister will walk the tree and display useful information

package analyzer

import (
	"fmt"

	"cthulhu/node"
	"cthulhu/token"
)

func Worklister(aAst *node.Node) {

	if aAst.Done {
		fmt.Printf("Node (%d, %d) [%s, \"%s\"] complete: %s\n",
			aAst.Line, aAst.Index, token.Name[aAst.Type], aAst.Text, node.FormatByteSlice(aAst.Code))
	}

	// If we don't have kids, we're done
	if len(aAst.Kids) == 0 {
		return
	}

	for _, k := range aAst.Kids {
		Worklister(k)
	}
}
