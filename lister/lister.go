// Lister Package for the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 11. May 2018

// The lister will produce a detailed listing of the final
// binary file with source code, comments and whatnot

package lister

import (
	"fmt"

	"cthulhu/node"
	"cthulhu/token"
)

const indent = "       "

func Lister(AST node.Node) {
	Walk(&AST)
}

// Generic walking function, depth first
func Walk(AST *node.Node) {

	switch AST.Type {

	case token.EOL:
		fmt.Print("\n")
	case token.DIREC:
		fmt.Print(indent, AST.Text, " ")
	case token.HEX_NUM:
		fmt.Print("$", AST.Text, " ")
	case token.BIN_NUM:
		fmt.Print("%", AST.Text, " ")
	case token.OPC_0:
		fmt.Print(indent, indent, AST.Text, " ")
	case token.OPC_1:
		fmt.Print(indent, indent, AST.Text, " ")
	default:
		fmt.Print(AST.Text, " ")
	}

	if len(AST.Kids) == 0 {
		return
	}

	for _, k := range AST.Kids {
		Walk(k)
	}
}
