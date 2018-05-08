// Token Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 08. May 2018

package parser

import (
	"log"

	"goasm65816/node"
	"goasm65816/token"
)

// TODO rewrite with channels?

// Root node starts first line, index zero
var AST = node.Node{token.Token{token.T_start, "ROOT", 1, 0}, nil, nil}

func Parser(tl *[]token.Token) node.Node {

	if len(*tl) == 0 {
		log.Fatal("PARSER FATAL: Received empty token list from Lexer")
	}

	// *** CREATE AST ***

	n := &AST

	for i := 0; i < len(*tl); i++ {

		t := (*tl)[i]

		switch t.Type {

		case token.T_directive:

			switch t.Text {

			// Directives with guarantied one parameter
			case ".mpu", ".notation", ".origin":
				dir := node.Node{t, nil, nil}
				n.Add(&dir)
				op := (*tl)[i+1]
				dir.Add(&node.Node{op, nil, nil})
				i += 1

			// Directives with guarantied two parameters
			case ".equ":
				dir := node.Node{t, nil, nil}
				n.Add(&dir)
				op1 := (*tl)[i+1]
				op2 := (*tl)[i+2]
				dir.Add(&node.Node{op1, nil, nil})
				dir.Add(&node.Node{op2, nil, nil})
				i += 2

			default:
				n.Add(&node.Node{t, nil, nil})
			}

		case token.T_opcode1:
			nn := node.Node{t, nil, nil}
			next := (*tl)[i+1]
			n.Add(&nn)
			nn.Add(&node.Node{next, nil, nil})
			i += 1

		// Most of our tokens are really easy to handle because they
		// don't have a parameter at all
		default:
			n.Add(&node.Node{t, nil, nil})
		}
	}

	return AST
}
