// Lister Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 07. May 2018

package lister

import (
	"fmt"

	"goasm65816/node"
	"goasm65816/token"
)

const indent = "       "

func Lister(AST node.Node) {
	Walk(&AST)
}

// Generic walking function, depth first
func Walk(AST *node.Node) {

	switch AST.Type {

	case token.T_eol:
		fmt.Print("\n")
	case token.T_directive:
		fmt.Print(indent, AST.Text, " ")
	case token.T_hex:
		fmt.Print("$", AST.Text, " ")
	case token.T_binary:
		fmt.Print("%", AST.Text, " ")
	case token.T_opcode0:
		fmt.Print(indent, indent, AST.Text, " ")
	case token.T_opcode1:
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
