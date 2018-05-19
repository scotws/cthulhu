// Parser of the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 18. May 2018

// The Cthulhu parser has one job: To create an Abstract Syntax Tree (AST) out
// of the list of tokens. All further processing is handled in later steps,
// because the initial AST is the basis of other tools such as the formatter.
// This is why we keep otherwise useless information like the empty lines and
// comments.

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
)

// Init sets up the parser for the init run, haven been given the list of tokens
func Init(ts *[]token.Token, tr bool) {

	if len(*ts) == 0 {
		log.Fatal("LEXER FATAL: Received empty token list.")
	}

	tokens = ts
	ast = node.Node{Token: token.Token{Type: token.START, Text: "Cthulhu"}}
	lookahead = (*tokens)[0]
	p = -1 // current will catch up with first consume()
}

// Parser is the actual parsing function. It takes a list of token.Tokens and
// returns the root node.Node to the whole program.
func Parser() *node.Node {

	for {
		ast.Kids = append(ast.Kids, walk())

		// This is how we end the whole parser
		if current.Type == token.EOF {
			break
		}
	}
	return &ast
}

// walk is the top-level function that starts parsing
func walk() *node.Node {

	var n node.Node

	switch lookahead.Type {

	case token.EOL, token.EOF, token.EMPTY, token.START:
		n = node.Create(lookahead)

	case token.COMMENT_LINE, token.COMMENT:
		// We keep comments and structural whitespace for the formatted output.
		// the lexer has already done all the heavy lifting for strings
		n = node.Create(lookahead)

	case token.DIREC_PARA:
		n = parseDirectPara()

	case token.DIREC, token.STRING, token.SYMBOL:
		// The lexer has already done some of the work swith strings
		n = node.Create(lookahead)

	case token.LABEL, token.ANON_LABEL, token.LOCAL_LABEL:
		n = node.Create(lookahead)

	case token.OPC_0:
		n = node.Create(lookahead)

	case token.OPC_1:
		// The first one is the opcode
		n = node.Create(lookahead)

		consume() // current is opcode, lookahead is operand

		// Now comes the operand. We can also load single character
		// values this way, so we have to test for strings as well
		o := parseOperand()
		n.Kids = append(n.Kids, o)
	}

	consume()

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

// backtrack is the reverse of consume: It move the current token back by one
// and pretends the last step never happened. We currently only use this for
// ranges
func backtrack() {
	p--
	current = (*tokens)[p]
	lookahead = (*tokens)[p+1]
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

	} else if lookahead.Type != tt {
		// It looks like we got ourselves a literal, which is nice
		wrongToken(tt, lookahead)
	}
}

// *** PARSING ROUTINES ***

// Parsing works by calling functions that return a node that may have
// subnodes. They all work by examining the loohahead token.

// parseNumber examines the lookahead token and throws an error if it is not one
// of the three literals  binary number, decimal number, or hex number. If the
// token is in fact a number, a new node is generated and a link to it returned
func parseNumber() *node.Node {
	t := lookahead.Type

	if t != token.HEX_NUM && t != token.DEC_NUM && t != token.BIN_NUM {
		log.Fatalf("PARSER FATAL (%d, %d): Expected number, got '%s' \n",
			lookahead.Line, lookahead.Index, token.Name[lookahead.Type])
	}

	n := node.Create(lookahead)
	return &n
}

// parseElement examines the lookahead token and throws an error if it is not a
// string or an expression. Otherwise, it returns a node of one of those two
// types
func parseElement() *node.Node {
	var n node.Node

	if lookahead.Type == token.STRING {
		n = node.Create(lookahead)
		consume()
	} else {
		n = *parseExpr()
	}

	return &n
}

// parseValue examines the lookahead token and throws and error if it is not
// a symbol, a number, a RPN term, or the ".here" directive that signals the
// current PC counter. If it is one of those, a link to a new node is returned
// that contains those structures
func parseValue() *node.Node {
	var n node.Node

	switch lookahead.Type {

	case token.L_CURLY:
		n = *parseRPN()
	case token.SYMBOL:
		n = node.Create(lookahead)
	case token.DIREC:
		if lookahead.Text != ".here" {
			log.Fatalf("PARSER FATAL (%d, %d): Directive '%s' is not a value",
				lookahead.Line, lookahead.Index, lookahead.Text)
		}

		n = node.Create(lookahead)
	default:
		n = *parseNumber()
	}
	return &n
}

// parseOperand checks the operand of an opcode that takes a single operand and
// returns a node if successful. This is a bit tricky because depending on the
// opcode we accept strings (eg for lda.# "a") and minus and plus (eg bra -).
// Returns a node if successful, otherwise throws an error. The opcode is
// expected to be in the lookahead token
func parseOperand() *node.Node {
	var n node.Node

	// Catch branching to local labels
	if lookahead.Type == token.MINUS || lookahead.Type == token.PLUS {
		n = node.Create(lookahead)
	} else {
		n = *parseElement()
	}
	return &n
}

// parseRPN parses the Reverse Polish Notation (RPN) math forms. We come here
// with the left curly brace as the lookahead token. Grammar rule is
//	rpn = "{" value { value | prn_operator } "}"
// Returns the RPN sequences as a node
func parseRPN() *node.Node {

	// create a new node of type RPN
	rt := token.Token{
		Type:  token.RPN,
		Text:  "RPN",
		Line:  lookahead.Line,
		Index: lookahead.Index,
	}

	rn := node.Create(rt)

	consume() // current is now the left curly brace, lookahead must be value

	// We need to have at least one value -- a number, a symbol, another RPN
	// term, or the ".here" directive -- in the lookahead
	t := parseValue()
	rn.Kids = append(rn.Kids, t)

	consume() // current is now the value, lookahead is unknown

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
		} else {
			vn := parseValue()
			rn.Kids = append(rn.Kids, vn)
		}

		consume()
	}
	return &rn
}

// parseRange creates a range node.
// We arrive here with the ellipsis token as
// lookahead and the first value as the current token.
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

	// The current token is the first expression. This is bad, because we
	// can't call parseExpr() with the current expression. So we backtrack
	backtrack()

	// Get our expression
	e1 := parseElement()
	rn.Kids = append(rn.Kids, e1)

	consume() // Current token is the ellipsis, the lookahead must be a value

	// Get our expression
	e2 := parseElement()
	rn.Kids = append(rn.Kids, e2)

	// After the range, we end up with current on the last value of the
	// range and whatever is after the range in lookahead

	return &rn
}

// parseExpr checks to see if the lookahead token is an expression or a simple
// math term, which can be either take a unary or binary operator. If not, it
// throws an error. If yes, it returns a pointer to a node of the type
// expression. The grammar specification is
// 	expr =  value | unary_operator value | value binary_operator value
// We arrive here with the lookahead token as the first part of the expression
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
			bn := node.Create(lookahead)
			en.Kids = append(en.Kids, &bn)

			consume() // current now operator, lookahead must be value

			// The next one must be a value or we're in trouble
			vn := parseValue()
			en.Kids = append(en.Kids, vn)
		}
	}
	return &en
}

// parseDirectPara handles directive nodes that have parameters. The lexer has
// taken care of making sure that we only have legal directives at this point.
// We arrive here with the directive as the lookahead token
func parseDirectPara() node.Node {

	// Save DIREC mother node
	n := node.Create(lookahead)

	consume() // current is now the directive, lookahead is unknown

	switch current.Text {

	case ".byte", ".word", ".long":
		// There is a lot of code duplication with .rom and .ram here,
		// though the difference is the string, which may not appear in
		// a range
		for lookahead.Type != token.EOL && lookahead.Type != token.COMMENT {

			switch lookahead.Type {

			case token.ELLIPSIS:
				// If we landed here, we have to backtrack
				// because we've already saved the first expr
				// from the range directive as a normal address
				n.Kids = n.Kids[0 : len(n.Kids)-1] // shorten list

				rn := parseRange() // parseRange handles backtracking
				n.Kids = append(n.Kids, rn)

			case token.COMMA:
				consume() // current now comma, lookahead must be expr

			case token.STRING:
				s := node.Create(lookahead)
				n.Kids = append(n.Kids, &s)
				consume()

			default:
				// Even if this turns out to be a range, the next token
				// must be an expression
				e := parseElement()
				n.Kids = append(n.Kids, e)
			}
		}

	case ".equ":
		// First token must be a symbol
		match(token.SYMBOL)
		n.Adopt(&n, &lookahead)

		consume() // current is symbol, lookahead must be expression
		e := parseExpr()
		n.Kids = append(n.Kids, e)

	case ".include":
		match(token.STRING)
		n.Adopt(&n, &lookahead)

	case ".mpu", ".assert":
		match(token.STRING)
		n.Adopt(&n, &lookahead)

	case ".origin", ".advance", ".skip":
		// Next token must be an expression
		e := parseExpr()
		n.Kids = append(n.Kids, e)

	case ".ram", ".rom":
		// This has a lot of overlap with .byte and friends, but we
		// can't use the same code because .ram and .rom don't accept
		// strings. We arrive here with a directive as the current token
		// and know that the lookahead must be an expression

		// We basically have a series of expressions separated by either
		// an comma, an ellipsis, or (terminally) by an EOL token. We
		// stop when we hit an EOL. The next line means that we will
		// accept an empty .ram or .rom statement at first
		for lookahead.Type != token.EOL {

			switch lookahead.Type {

			case token.ELLIPSIS:
				// If we landed here, we have to backtrack
				// because we've already saved the first expr
				// from the range directive as a normal address
				n.Kids = n.Kids[0 : len(n.Kids)-1] // shorten list

				rn := parseRange() // parseRange handles backtracking
				n.Kids = append(n.Kids, rn)

			case token.COMMA:
				consume() // current now comma, lookahead must be expr

			default:
				// Even if this turns out to be a range, the next token
				// must be an expression
				e := parseExpr()
				n.Kids = append(n.Kids, e)
			}
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
