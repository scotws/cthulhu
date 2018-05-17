// Parser of the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 16. May 2018

// The Cthulhu parser has one job: To create an Abstract Syntax Tree (AST) out
// of the list of tokens. All further processing is handled in later steps,
// because the initial AST is the basis of other tools such as the formatter,
// which is why we keep otherwise useless information like the empty lines and
// comments.

package parser

import (
	"log"

	"cthulhu/node"
	"cthulhu/token"
)

var (
	tokens    *[]token.Token // list of tokens to parse (from lexer)
	p         int            // index to the current token we're looking at
	current   token.Token    // current token we're examining
	lookahead token.Token    // one token lookahead
	ast       node.Node      // root of the Abstract Syntax Tree (AST)
	mpu       string         // MPU type requested by the user
	trace     bool           // User requests lots and lots of info
)

// Init sets up the parser for the init run, haven been given the list of tokens
func Init(ts *[]token.Token, tr bool) {

	if len(*ts) == 0 {
		log.Fatal("LEXER FATAL: Received empty token list.")
	}

	tokens = ts
	p = -1 // bumped up to 0 by initial call to consume
	ast = node.Node{Token: token.Token{Type: token.START, Text: "Cthulhu"}}
	trace = tr
}

// Parser is the actual parsing function. It takes a list of token.Tokens and
// returns the root node.Node to the whole program. Errors
// are handled here, currently mostly by fatal logging (for now)
func Parser() *node.Node {

	for {
		ast.Kids = append(ast.Kids, walk())

		// This is what we actually use to end
		if current.Type == token.EOF {
			break
		}
	}

	return &ast
}

// consume moves the pointer to the current token up by one and retrieves the next
// token from the token list. This could also be called next
func consume() {

	if p+1 < len(*tokens) {
		p++
		current = (*tokens)[p]

		if p+1 < len(*tokens) {
			lookahead = (*tokens)[p+1]
		}
	} else {
		lookahead = token.Token{Type: token.EOF}
	}
}

// match takes a token type and checks it against the next (lookahead) token. If
// they are a literal match -- say, a STRING and a STRING -- it consumes the
// current token, making the lookahead the current one. If the token type is a
// composite -- say, token.NUMBER, it checks to see if it is legal such as
// token.BIN_NUM and returns that sub_type.
func match(tt int) {

	found := false

	// Walk through the composite types to see if this is a legal subtype
	if token.IsComposite(tt) {

		sts := token.SubTypes(tt)

		for _, t := range sts {

			if t == lookahead.Type {
				found = true
				break
			}
		}

		if !found {
			wrongToken(tt, lookahead)
		}

		// TODO deal with complex tokens such as RPN
	} else if token.IsComplex(tt) {

		log.Fatalf("PARSER FATAL: Can't deal with complex tokens yet")

		// TODO Deal with RPN
		// TODO Deal with EXPR
		// TODO Deal with RANGE

		// It looks like we got ourselves a literal, which is nice
	} else if lookahead.Type != tt {
		wrongToken(tt, lookahead)
	}

	consume()
}

// walk
func walk() *node.Node {

	var n node.Node

	consume()

	switch current.Type {

	case token.EOL, token.EOF, token.EMPTY, token.START:
		n = node.Create(current)

	// We keep comments for the formatted output; the lexer has already done
	// all the heavy lifting for strings
	case token.COMMENT_LINE, token.COMMENT, token.STRING:
		n = node.Create(current)

	case token.DIREC_PARA:
		n = parseDirectPara()

	case token.DIREC:
		n = node.Create(current)

	case token.LABEL, token.ANON_LABEL, token.LOCAL_LABEL:
		n = node.Create(current)
	}

	return &n
}

// Parse directive nodes with parameters. The lexer has taken care of making
// sure that we only have legal directives at this point
func parseDirectPara() node.Node {

	// Save DIREC mother node
	n := node.Create(current)

	switch current.Text {

	case ".equ":
		// First token must be a symbol
		match(token.SYMBOL)
		n.Adopt(&n, &current)

		// Next token must be an expression
		match(token.NUMBER) // TODO Testing, replace by parseExpr()
		n.Adopt(&n, &current)

	// TODO see if we even should have include files here anymore
	case ".include":
		match(token.STRING)
		n.Adopt(&n, &current)

	case ".mpu":
		// Next token must be a string
		match(token.STRING)
		n.Adopt(&n, &current)

	case ".origin":
		// Next token must be a NUMBER
		match(token.NUMBER)
		n.Adopt(&n, &current)

	case ".ram", ".rom":

		// First token must be an expression
		match(token.NUMBER) // TODO Testing, replace by EXPR

		// If we survived that, the current token is a NUMBER. We don't
		// save it quite yet, though, because we need to see if we're
		// dealing with a range

		switch lookahead.Type {

		case token.ELLIPSIS:

			// We insert a RANGE token as the first child of current
			// token and add the other children to it
			rt := token.Token{
				Type:  token.RANGE,
				Line:  current.Line,
				Index: current.Type,
				Text:  "RANGE",
			}

			// Create a range node
			rn := node.Create(rt)

			// At this point, the current token is still the first
			// number and the lookahead is the ellipsis. Next step
			// is to add the range to the RAM/ROM node
			n.Kids = append(n.Kids, &rn)

			// Now we add the current token as the first element of
			// the range
			rn.Adopt(&rn, &current)

			// get rid of RAM/ROM token, make RANGE current,
			// lookahead is second number
			consume()

			// Next token must be a number
			match(token.NUMBER) // TODO Testing, replace by EXPR
			rn.Adopt(&rn, &current)

		case token.COMMA:
			// TODO must be processed as "LIST"
			n.Adopt(&n, &current) // save number
			match(token.COMMA)
			match(token.NUMBER) // TODO Testing, replace by EXPR
			n.Adopt(&n, &current)

		case token.EOL:
			n.Adopt(&n, &current)
			consume()

		default:
			log.Fatalf("PARSER FATAL (%d, %d): Expected \",\" or \"...\", got '%d' \n",
				current.Line, current.Index, lookahead.Type)
		}

	}
	return n

}

// wrongToken takes the token we want, the token we got, and complains by
// crashing that they are not the same
func wrongToken(want int, got token.Token) {
	log.Fatalf("PARSER FATAL (%d, %d): Expected token type '%s', got '%s'\n",
		got.Line, got.Index, token.Name[want], token.Name[got.Type])
}
