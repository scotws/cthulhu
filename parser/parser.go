// Parser of the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 18. May 2018

// The Cthulhu parser has one job: To create an Abstract Syntax Tree (AST) out
// of the list of tokens. All further processing is handled in later steps,
// because the initial AST is the basis of other tools such as the formatter.
// This is why we keep otherwise useless information like the empty lines and
// comments.

// TODO convert to all working on lookahead, not current

package parser

import (
	"log"

	"cthulhu/data"
	"cthulhu/node"
	"cthulhu/token"
)

var (
	tokens    *[]token.Token // list of tokens to parse (from lexer)
	p         int            // index to the current token we're looking at
	current   token.Token    // current token we're examining
	lookahead token.Token    // one token lookahead
	ast       node.Node      // root of the Abstract Syntax Tree (AST)
	mpu       string         // MPU type requested by the user TODO see if needed
	trace     bool           // User requests lots and lots of info TODO see if needed
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
// returns the root node.Node to the whole program.
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

// walk
// TODO see if we should be working with lookahead token, not current token
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

// consume moves the pointer to the current token up by one and retrieves the next
// token from the token list. This could also be called "next"
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

// match takes a token type and checks it against the lookahead token. If they
// are a literal match -- say, a STRING and a STRING -- it consumes the current
// token, making the lookahead the current one. If the token type is a composite
// -- say, token.NUMBER -- it checks to see if it is legal such as token.BIN_NUM
// and returns that sub_type.
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

// *** PARSING ROUTINES ***

// Parsing works by calling functions that return a node that might have
// subnodes. They work by examining the current token.

// parseNumber examines the current token and throws an error if it is not one
// of the three literals  binary number, decimal number, or hex number. If the
// token is a number, a new node is generated and a link to it returned
func parseNumber() *node.Node {
	t := current.Type

	if t != token.HEX_NUM && t != token.DEC_NUM && t != token.BIN_NUM {
		log.Fatalf("PARSER FATAL (%d, %d): Expected \",\" or \"...\", got '%s' \n",
			current.Line, current.Index, token.Name[lookahead.Type])
	}

	n := node.Create(current)
	return &n
}

// parseValue examines the current token and throws and error if it is not
// a symbol, a number, a RPN term, or the ".here" directive that signals the
// current PC counter. If it is one of those, a link to a new node is returned
// that contains those structures
func parseValue() *node.Node {
	var n node.Node

	switch current.Type {

	case token.L_CURLY:
		n = *parseRPN()
	case token.SYMBOL:
		n = node.Create(current)
	case token.DIREC:
		if current.Text != ".here" {
			log.Fatalf("PARSER FATAL (%d, %d): Directive '%s' is not a value",
				current.Line, current.Index, current.Text)
		}

		n = node.Create(current)
	default:
		n = *parseNumber()
	}

	return &n
}

// parseRPN parses the Reverse Polish Notation (RPN) math forms. We come here
// with the left curly brace as the current token. Grammar rule is
//	rpn = "{" value { value | prn_operator } "}"
// Returns the RPN sequences as a node
func parseRPN() *node.Node {

	// create a new node of type RPN
	rt := token.Token{
		Type:  token.RPN,
		Text:  "RPN",
		Line:  current.Line,
		Index: current.Index,
	}

	rn := node.Create(rt)

	consume() // current is now the first element, lookahead is unknown

	// We need to have at least one value -- a number, a symbol, another RPN
	// term, or the ".here" directive -- to return
	t := parseValue()
	rn.Kids = append(rn.Kids, t)

	// While we've not been told to stop, add stuff to the RPN term's
	// children
	for lookahead.Type != token.R_CURLY {

		// We shouldn't hit an end of line without a closing curly brace
		if lookahead.Type == token.EOL {
			log.Fatalf("PARSER FATAL (%d, %d): RPN term missing closing brace",
				lookahead.Line, lookahead.Index)
		}

		// After the initial value, we can either have another value or
		// an operator that is legal for the RPN
		_, ok := data.OperatorsRPN[lookahead.Text]
		if ok {
			rn.Adopt(&rn, &lookahead)
			consume()
			continue
		}

		// This has to be a value then or else were in trouble
		consume()
		vn := parseValue()
		rn.Kids = append(rn.Kids, vn)
	}
	return &rn
}

// parseRange creates a range node and returns it with the second expression in
// place and the first one empty. We arrive here with the ellipsis token as
// current.
// TODO see if we should rewrite this with lookahead, starting on the first
// expression
func parseRange() *node.Node {

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

	// Now we add a dummy expression as the first element
	// of the range
	rn.Kids = append(rn.Kids, &node.Node{})

	consume() // current is now an expression, lookahead unknown

	// Get our expression
	e2 := parseExpr()
	rn.Kids = append(rn.Kids, e2)

	return &rn

}

// parseExpr checks to see if the current token is an expression or a simple
// math term. If not, it throws an error. If yes, it returns a pointer to a node
// of the type expression
// specification is
// 	expr =  value | uni_operator value | value bin_operator value
func parseExpr() *node.Node {

	// create a new node of type EXPR
	et := token.Token{
		Type:  token.EXPR,
		Text:  "EXPR",
		Line:  current.Line,
		Index: current.Index,
	}

	en := node.Create(et)

	// See if we have a unary (single) operator
	_, ok := data.OperatorsUnary[lookahead.Text]
	if ok {

		// Add the unary operator to slice
		un := node.Create(lookahead)
		en.Kids = append(en.Kids, &un)

		consume() // uniary now current, lookahead must be value
		vn := parseValue()
		en.Kids = append(en.Kids, vn)

	} else {
		// One way or another, the lookahead must be a value
		vn := parseValue()
		en.Kids = append(en.Kids, vn)
		consume() // value now current, lookahead unknown

		// We either are done or we have a binary operator
		_, ok := data.OperatorsBinary[lookahead.Text]
		if ok {
			// This is a binary operation. Add the binary operator
			// to the slice
			consume() // current now operator, lookahead must be value
			bn := node.Create(lookahead)
			en.Kids = append(en.Kids, &bn)

			// The next one must be a value or we're in trouble
			vn := parseValue()
			en.Kids = append(en.Kids, vn)
		}
	}
	return &en
}

// parseDirectPara handles directive nodes that have parameters. The lexer has
// taken care of making sure that we only have legal directives at this point.
// We arrive here with the directive as the current token
func parseDirectPara() node.Node {

	// Save DIREC mother node
	n := node.Create(current)

	switch current.Text {

	case ".equ":
		// First token must be a symbol
		match(token.SYMBOL)
		n.Adopt(&n, &current)
		consume() // current now expression, lookahead unknown

		e := parseExpr()
		n.Kids = append(n.Kids, e)

	case ".include":
		match(token.STRING)
		n.Adopt(&n, &current)

	case ".mpu":
		// Next token must be a string
		match(token.STRING)
		n.Adopt(&n, &current)

	case ".origin":
		// Next token must be an expression
		consume()
		e := parseExpr()
		n.Kids = append(n.Kids, e)

	case ".ram", ".rom":
		// This has a lot of overlap with .byte and friends, but we
		// can't use the same code because .ram and .rom don't accept
		// strings

		consume() // current must now be expression, lookahead unknown
		e1 := parseExpr()

		// Don't save expression yet, need to see if this is a range or
		// a list

		switch current.Type {

		case token.ELLIPSIS:

			rn := parseRange()

			// Add the range to the RAM/ROM node
			n.Kids = append(n.Kids, rn)

			// Now we add the current token as the first element of
			// the range
			rn.Kids[0] = e1

		case token.COMMA:
			// If this is a list, we first store the entry we picked
			// up

			for current.Type != token.EOL {
				consume() // current now expression, lookhead number

				if current.Type == token.ELLIPSIS {
					e2 := parseRange()
					e2.Kids[0] = e1
					n.Kids = append(n.Kids, e2)
				} else {
					// This is a normal entry
					e2 := parseExpr()
					n.Kids = append(n.Kids, e1)
					n.Kids = append(n.Kids, e2)
				}
			}

		case token.EOL:
			// This is a single entry line
			n.Kids = append(n.Kids, e1)

		default:
			log.Fatalf("PARSER FATAL (%d, %d): Expected \",\" or \"...\", got '%s' \n",
				current.Line, current.Index, token.Name[current.Type])
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
