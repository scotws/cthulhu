// Purge: Analysis step for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 21. May 2018
// This version: 21. May 2018

// Purge is the first step of the analysis phase. It takes the Abstract
// Syntax Tree (AST) created by the parser and removes the whitespace and other
// nodes that are used for output formatting, but are just deadweight for the
// further processng towards a binary file. Also, simple number conversions are
// handled.

package analyzer

import (
	"cthulhu/node"
)

const (
	errPrelude = "ANALYZER"
)

var (
	bst *node.Node
)

func Purge(mpu string, ast *node.Node) *node.Node {

	// Walk AST. If node is a comment, an empty line or an EOL node, ignore
	// it. Save the others to BST.

	// Testing
	bst = ast
	return bst
}
