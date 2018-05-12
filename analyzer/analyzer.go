// Analyzer package for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 12. May 2018
// This version: 12. May 2018

package analyzer

import "fmt"

// The analyzer walks the Abstract Syntax Tree (AST) created by the parser and
// modifies it in various ways
func Analyzer(ast *node.Node) *node.Node {
	fmt.Println("(Analyzer not working yet)")
}

// Walk is the main internal routine that visits every node and does something
// depending on type
func walk(tree *node.Node) {
	fmt.Println("(Analyzer first pass)")
	fmt.Println("(Analyzer second pass)")
}
