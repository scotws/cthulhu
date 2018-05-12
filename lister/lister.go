// Lister Package for the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 11. May 2018

// The lister will produce a detailed listing of the final
// binary file with source code, comments and whatnot
// Currently, this is more just a version of the formatter

package lister

import (
	"fmt"

	"cthulhu/node"
	"cthulhu/token"
)

const indent = "       "

func Lister(ast *node.Node) {
	Walk(ast)
}

// Generic walking function, depth first
func Walk(ast *node.Node) {

	switch ast.Type {

	case token.EOL:
		fmt.Print("\n")
	case token.EMPTY:
		fmt.Print(" \n")
	case token.DIREC, token.DIREC_PARA:
		fmt.Print(indent, ast.Text, " ")
	case token.LABEL, token.LOCAL_LABEL:
		fmt.Print(ast.Text, ":")
	case token.HEX_NUM:
		fmt.Print("$", ast.Text, " ")
	case token.BIN_NUM:
		fmt.Print("%", ast.Text, " ")
	case token.OPC_0, token.OPC_1:
		fmt.Print(indent, indent, ast.Text, " ")
	default:
		fmt.Print(ast.Text, " ")
	}

	if len(ast.Kids) == 0 {
		return
	}

	for _, k := range ast.Kids {
		Walk(k)
	}
}
