// Token structure for Cthulhu Assembler
// Scot W. Stevenson
// First version: 02. May 2018
// This version: 18. May 2018

// So here's a funny thing: The Go specs instist that you should always use
// camel case, and not all caps, even for constants. However, if you take a look
// at the Go tokens at https://golang.org/src/go/token/token.go you find ...
// gee, the tokens are in all caps. This would seem to be a classic case of
// "code as a say, not as I do", but we're not having any of it.

package token

type Token struct {
	Type  int
	Line  int    // starts with 1
	Index int    // starts with 1
	File  string // name of the file the token comes from
	Text  string // raw text
}

const (
	// We start off with literals, which are easy to work with
	lit_begin    int = iota // start of all literals
	EOF                     // end of file
	START                   // start of file
	EOL                     // end of line
	EMPTY                   // marks empty line for formatting
	COMMENT                 // in-line comments
	COMMENT_LINE            // whole-line comments
	DIREC                   // simple directive, no parameters
	DIREC_PARA              // directive with one or more parameters
	ELLIPSIS                // directive "..." used for various things
	OPC_0                   // SAN opcode without any operands ("nop")
	OPC_1                   // SAN opcode with exactly one operand
	OPC_2                   // SAN opcode with two operands (mvn, mvp)
	BIN_NUM                 // binary number
	HEX_NUM                 // hexadecimal number
	DEC_NUM                 // decimal number
	LABEL                   // absolute label (ends with ":")
	LOCAL_LABEL             // scoped label (starts with "_")
	ANON_LABEL              // anonymous label (starts with "@")
	L_SQUARE                // [
	R_SQUARE                // ]
	L_PARENS                // (
	R_PARENS                // )
	L_CURLY                 // {
	R_CURLY                 // }
	GREATER                 // >
	LESS                    // <
	COMMA                   // ,
	MINUS                   // -
	PLUS                    // +
	SLASH                   // /
	STAR                    // *
	HASH                    // #
	STRING                  // not a list of runes
	SYMBOL                  // any symbol
	EQUAL                   // =
	AMPERSAND               // &
	PIPE                    // |
	PERCENT                 // %
	DOLLAR                  // $
	PERIOD                  // .
	TILDE                   // ~
	CARET                   // ^
	lit_end

	// After the literals come the composits which need rules to check them
	composite_begin
	ADDRESS
	NUMBER
	composite_end

	// Then we have the complex types
	complex_begin
	EXPR  // Expressions
	RPN   // Reverse Polish Notation (RPN) math terms
	RANGE // two numbers connected by an ELLIPSIS
	complex_end
)

var Name = map[int](string){
	EOF:          "EOF",
	START:        "START",
	EOL:          "EOL",
	EMPTY:        "EMPTY",
	ELLIPSIS:     "ELLIPSIS",
	COMMENT:      "COMMENT",
	COMMENT_LINE: "COMMENT_LINE",
	DIREC:        "DIREC",
	DIREC_PARA:   "DIREC_PARA",
	OPC_0:        "OPC_0",
	OPC_1:        "OPC_1",
	OPC_2:        "OPC_2",
	BIN_NUM:      "BIN_NUM",
	HEX_NUM:      "HEX_NUM",
	DEC_NUM:      "DEC_NUM",
	LABEL:        "LABEL",
	LOCAL_LABEL:  "LOCAL_LABEL",
	ANON_LABEL:   "ANON_LABEL",
	L_SQUARE:     "LEFT_SQUARE",
	R_SQUARE:     "RIGHT_SQUARE",
	L_PARENS:     "LEFT_PARENS",
	R_PARENS:     "RIGHT_PARENS",
	L_CURLY:      "LEFT_CURLY",
	R_CURLY:      "RIGHT_CURLY",
	GREATER:      "GREATER",
	LESS:         "LESS",
	COMMA:        "COMMA",
	MINUS:        "MINUS",
	PLUS:         "PLUS",
	SLASH:        "SLASH",
	STAR:         "STAR",
	HASH:         "HASH",
	STRING:       "STRING",
	SYMBOL:       "SYMBOL",
	EQUAL:        "EQUAL",
	AMPERSAND:    "AMPERSAND",
	PIPE:         "PIPE",
	PERCENT:      "PERCENT",
	DOLLAR:       "DOLLAR",
	PERIOD:       "PERIOD",
	TILDE:        "TILDE",
	CARET:        "CARET",

	// Composite types
	ADDRESS: "ADDRESS",
	NUMBER:  "NUMBER", // either BIN_NUM, DEC_NUM, or HEX_NUM

	// Complex types: Made up out of totally different types of tokens
	RANGE: "RANGE", // includes ELLIPSIS "..."
	EXPR:  "EXPR",
	RPN:   "RPN", // includes complex math terms encased in "{" and "}"
}

// compositeTypes is a map that contains the literal subtypes that composite
// types contain
var compositeTokens = map[int][]int{
	ADDRESS: []int{SYMBOL, HEX_NUM, BIN_NUM},  // TODO missing math
	NUMBER:  []int{HEX_NUM, BIN_NUM, DEC_NUM}, // TODO missing math
}

// IsLiteral checks to see if the given token is a literal (say, HEX_NUM) or a
// composite value (say, NUMBER) that needs further testing
func IsLiteral(tt int) bool {
	f := false

	if tt > lit_begin && tt < lit_end {
		f = true
	}
	return f
}

// IsComplex checks to see if the given token is a complex type
// value (say, RPN)
func IsComplex(tt int) bool {
	f := false

	if tt > complex_begin && tt < complex_end {
		f = true
	}
	return f
}

// IsComposite checks to see if the given token is a composite
// value (say, NUMBER)
func IsComposite(tt int) bool {
	f := false

	if tt > composite_begin && tt < composite_end {
		f = true
	}
	return f
}

// Subtypes is given a composite type and returns a list of literal types to
// check against, such as NUMBER to HEX_NUM, BIN_NUM and DEC_NUM
func SubTypes(ct int) []int {
	at := compositeTokens[ct]
	return at
}
