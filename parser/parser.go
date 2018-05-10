// Token Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 09. May 2018

package parser

import (
	"log"

	"goasm65816/node"
	"goasm65816/token"
)

var (
	tokens *[]token.Token // list of tokens to parse (from lexer)
	AST    node.Node      // root of the Abstract Syntax Tree (AST)
	cur    int            // index of the current character we are looking at
)

// There are two ways to set up a parser like this: Either you can use
// nextToken, consume and match to move along without apparent indexing, or you
// use a normal for loop with the index. We'll go with the index because it
// gives greater flexibility in looking ahead.

// Parser is the actual parsing function. It takes a list of token.Tokens and
// returns a node that contains a list of pointers to the whole program. Errors
// are handled here, currently mostly by fatal logging (for now)
func Parser(ts *[]token.Token) node.Node {

	if len(*ts) == 0 {
		log.Fatal("PARSER FATAL: Did not receive any tokens from lexer")
	}

	AST = node.Node{token.Token{token.T_start, "ROOT", 1, 0}, nil}
	tokens = ts

	for cur = 0; cur < len(*tokens); cur++ {

		t := (*tokens)[cur]

		// We end if we get the EOF token regardless of where we are in
		// the token list
		if t.Type == token.T_eof {
			break
		}

		// We need to define this before the switch statement, but after
		// the beginning of the loop or the definition will not work
		var n node.Node

		// This is where the magic happens: Call various routines based
		// on what type we have
		switch t.Type {

		case token.T_directivePara:
			directivePara(t)
			continue

		// The simple stuff just gets packed in a node and returned
		default:
			n = node.Create(t)
		}
		AST.Add(&n)
	}

	return AST
}

// nextToken moves to the next token and returns it. If the next token is an
// end-of-file, we crash because the assumption is that we call nextToken while
// getting parameters
func nextToken() token.Token {

	cur++
	n := (*tokens)[cur]

	if n.Type == token.T_eof {
		log.Fatalf("PARSER FATAL: Ran out of tokens at '%s' (%d,%d)\n",
			n.Text, n.Line, n.Index)
	}
	return n
}

// match gets the next token, making sure that it is of the type we requested
func match(want int) token.Token {
	t := nextToken()
	if t.Type != want {
		log.Fatalf("PARSER FATAL (%d,%d): Wanted token type %s, got %s",
			t.Line, t.Index, token.Name[want], token.Name[t.Type])
	}
	return t
}

func adopt(mom *node.Node, kid *token.Token) {
	kidnode := node.Create(*kid)
	mom.Add(&kidnode)
}

// ***** INDIVIDUAL FUNCTIONS *****

// Directives with parameters
func directivePara(t token.Token) {

	// We know that we must have a legal directive because the lexer checked
	// that for us
	n := node.Create(t)
	AST.Add(&n)

	switch t.Text {

	case ".mpu":
		kt := match(token.T_string)

		if kt.Text != "65816" && kt.Text != "65c02" && kt.Text != "6502" {
			log.Fatalf("PARSER FATAL (%d,%d): MPU type '%s' not supported",
				kt.Line, kt.Index, kt.Text)
		}

		adopt(&n, &kt)

	case ".notation":
		kt := match(token.T_string)

		if kt.Text != "san" && kt.Text != "wdc" {
			log.Fatalf("PARSER FATAL (%d,%d): notation '%s' not supported",
				kt.Line, kt.Index, kt.Text)
		}

		adopt(&n, &kt)

	case ".origin":
		kt := match(token.T_hex)
		adopt(&n, &kt)

	}
}
