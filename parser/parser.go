// Parser of the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 14. May 2018

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

// Rule of thumb:

// - parse routines deal with individual elements of a grammar rule, such as
// expressions. They take no parameters and return a node. They work on the
// lookahead token, not the current token

// - match() checks for a fixed literal (for instance, "{") and silently consue
// it, only raising an error if something goes wrong

// walk
func walk() *node.Node {

	var n node.Node

	consume()

	switch current.Type {

	case token.EOL, token.EOF, token.EMPTY, token.START:
		n = node.Create(current)
	case token.COMMENT_LINE, token.COMMENT:
		n = node.Create(current)
	case token.DIREC_PARA:
		n = parseDirectPara()
	}

	return &n
}

// Parse directive nodes with parameters
func parseDirectPara() node.Node {

	var n node.Node

	switch current.Text {

	case ".equ":
		// Save DIREC mother node
		n = node.Create(current)

		// Next token must be a symbol
		if lookahead.Type != token.SYMBOL {
			wrongToken(token.SYMBOL, lookahead)
		}
		consume()
		n.Adopt(&n, &current)

		// Next token must be an expression
		consume()
		kt := parseNumber() // TODO Testing, replace by parseExpr()
		n.Adopt(&n, &kt)

	case ".mpu":
		n = node.Create(current)

		if lookahead.Type != token.STRING {
			wrongToken(token.STRING, lookahead)
		}

		n.Adopt(&n, &lookahead)
		consume()

	}
	return n

}

func parseNumber() token.Token {

	t := current.Type

	if t != token.DEC_NUM &&
		t != token.HEX_NUM &&
		t != token.BIN_NUM {
		wrongToken(token.NUMBER, current)
	}
	return current
}

// wrongToken takes the token we want, the token we got, and complains by
// crashing that they are not the same
func wrongToken(want int, got token.Token) {
	log.Fatalf("PARSER FATAL (%d, %d): Expected token type '%s', got '%s'\n",
		got.Line, got.Index, token.Name[want], token.Name[got.Type])
}

/*

// match takes a token type and silently confirms that the current type is what
// we want it to be -- otherwise, it fails, currently with a fatal log message.
// For literal types, this can be quick, for composite types, we need different
// rules
// TODO this needs to be defined recursively
func (p *Parser) match(want int) {

	found := false
	p.consume()
	t := p.tokens[p.cur] // our new current token

	// If this is a composite type, we have to walk through all the literal
	// subtypes
	if t.IsComposite(want) {

		for _, got := range token.Subtypes(want) {

			if got == t.Type {
				found = true
				break
			}
		}

	} else {
		// This is already a literal
		if t.Type == want {
			found = true
		}
	}

	if !found {
		log.Fatalf("PARSER FATAL (%d, %d): Expected token type '%s', got '%s'\n",
			t.Line, t.Index, token.Name[want], token.Name[t.Type])
	}
}



// ***** LITERAL FUNCTIONS*****

func (p *Parser) parseString() {

	t := p.lookahead

	if t.Type != token.STRING {
		wrongToken(t.Line, t.Index, token.STRING, t.Type)
	}

	p.match(token.STRING)

	t := p.tokens[p.cur] // our new current token

	if t.Type != token.STRING {

	p.ast.Adopt(&n, &p.lookahead)
	p.consume()
}

// Directives with parameters
func (p *Parser) parseDirecPara() {

	t := p.tokens[p.cur]

	// We know that we must have a legal directive because the lexer checked
	// for us. We store that first
	n := node.Create(t)
	p.ast.Add(&n)

	switch t.Text {

	/*
		case ".byte":

			parseExpr()

				kt := match(token.DEC_NUM)
				adopt(&n, &kt)

				var nt token.Token

				for {
					nt = nextToken()

					// TODO This needs to be a lot more clever
					if nt.Type == token.DEC_NUM {
						adopt(&n, &nt)
					}

					p.match(token.COMMA)

					if nt.Type == token.EOL {
						adopt(&n, &nt) // need final EOL

	case ".include":
		p.ast.Adopt(&n, parseString())

		// p.match(token.STRING)
		// p.ast.Adopt(&n, &p.tokens[p.cur])

	case ".origin":
		p.match(token.NUMBER)
		p.ast.Adopt(&n, &p.tokens[p.cur])

	case ".ram":
		p.match(token.ADDRESS)
		p.ast.Adopt(&n, &p.tokens[p.cur])
		p.match(token.ADDRESS)
		p.ast.Adopt(&n, &p.tokens[p.cur])

	case ".rom":
		p.match(token.ADDRESS)
		p.ast.Adopt(&n, &p.tokens[p.cur])
		p.match(token.ADDRESS)
		p.ast.Adopt(&n, &p.tokens[p.cur])
	}
}

*/
