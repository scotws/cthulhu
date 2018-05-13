// Parser of the Cthulhu assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 13. May 2018

// The Cthulhu parser has one job: To create an Abstract Syntax Tree (AST) out
// of the list of tokens. All further processing is handled in later steps. Note
// this means that we also include information such as comments that would
// usually be thrown out -- we use the AST to output clean formatting, however.

package parser

import (
	"log"

	"cthulhu/node"
	"cthulhu/token"
)

var (
	tokens *[]token.Token // list of tokens to parse (from lexer)
)

// The Go parser (see https://golang.org/src/go/parser/parser.go) keeps
// everything in a struct, so that can't be a bad way to do it.
type Parser struct {
	tokens []token.Token // list of tokens to be modified
	cur    int           // index to the current token we're looking at
	tok    token.Token   // one token lookahead
	ast    node.Node     // root of the Abstract Syntax Tree (AST)
	mpu    string        // MPU type requested by the user
	trace  bool          // User requests lots and lots of info
}

// Init sets up the parser for the init run, haven been given the list of tokens
func (p *Parser) Init(ts *[]token.Token, wantMPU string, trace bool) {

	if len(*ts) == 0 {
		log.Fatal("LEXER FATAL: Received empty token list.")
	}

	p.tokens = *ts
	p.ast = node.Node{Token: token.Token{Type: token.START, Text: "ROOT"}}
	p.cur = -1 // pointer to current will be increased during first run
	p.mpu = wantMPU
	p.trace = trace
}

// next moves the pointer to the current token up by one and retrieves the next
// token from the token list. For things to work, we definitely need a EOF (end
// of file) token
func (p *Parser) next() {

	// Move to next token
	p.cur++

	// Get lookahead token
	if p.cur+1 <= len(p.tokens) {
		p.tok = p.tokens[p.cur+1]
	} else {
		p.tok = token.Token{Type: token.EOF} // paranoid
	}
}

// match takes a token type and silently confirms that the current type is what
// we want it to be -- otherwise, it fails, currently with a fatal log message.
// For literal types, this can be quick, for composite types, we need different
// rules
// TODO this needs to be defined recursively
func (p *Parser) match(want int) {

	found := false
	p.next()
	t := p.tokens[p.cur] // our new current token

	// If this is a composite type, we have to walk through all the literal
	// subtypes
	if !t.IsLiteral(want) {

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

// Parser is the actual parsing function. It takes a list of token.Tokens and
// returns the root node.Node to the whole program. Errors
// are handled here, currently mostly by fatal logging (for now)
func (p *Parser) Parse() *node.Node {

	for {
		// Get next token and lookahead
		p.next()

		// Continue until we're done
		if p.tok.Type == token.EOF {
			final := node.Create(p.tok) // We need EOF node
			p.ast.Add(&final)
			break
		}

		// Main switch statement to create new nodes

		var n node.Node
		t := p.tokens[p.cur] // just too lazy to type

		switch t.Type {

		case token.DIREC_PARA:
			p.parseDirecPara()

		default:
			n = node.Create(t)
			p.ast.Add(&n)
		}
	}

	return &p.ast
}

// ***** INDIVIDUAL FUNCTIONS *****

// Directives with parameters
func (p *Parser) parseDirecPara() {

	t := p.tokens[p.cur]

	// We know that we must have a legal directive because the lexer checked
	// for us. We store that first
	n := node.Create(t)
	p.ast.Add(&n)

	switch t.Text {

	case ".equ":
		p.match(token.SYMBOL)
		p.ast.Adopt(&n, &p.tokens[p.cur])
		p.match(token.NUMBER) // TODO can be another symbol or RPR etc
		p.ast.Adopt(&n, &p.tokens[p.cur])

	case ".include":
		p.match(token.STRING)
		p.ast.Adopt(&n, &p.tokens[p.cur])

	// In theory, we could already check here if the requested MPU and the
	// MPU from the source file match. However, the parser doesn't modify
	// information because it's the base for the nicely formatted output
	case ".mpu":
		p.match(token.STRING)
		p.ast.Adopt(&n, &p.tokens[p.cur])

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

/*
	// We only accept decimal numbers for now
	// TODO accept other stuff
	// TODO should work for .word and .long as well
	case ".byte":

		kt := match(token.DEC_NUM)
		adopt(&n, &kt)

		var nt token.Token

		for {
			nt = nextToken()

			// TODO This needs to be a lot more clever
			if nt.Type == token.DEC_NUM {
				adopt(&n, &nt)
			}

			if nt.Type == token.COMMA {
				continue
			}
			if nt.Type == token.EOL {
				adopt(&n, &nt) // need final EOL
				break
			}
		}
	}
}
*/
