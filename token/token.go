// Token structure for Cthulhu Assembler
// Scot W. Stevenson
// First version: 02. May 2018
// This version: 11. May 2018

// So here's a funny thing. The Go specs instist that you should always use
// camel case, and not all caps, even for constants. However, if you take a look
// at the Go tokens at https://golang.org/src/go/token/token.go you find ...
// yes, the tokens are in all caps. This would seem to be a classic case of
// "code as a say, not as I do", but we're not having any of it.

package token

const (
	EOF          int = iota // end of file
	START                   // start of file
	EOL                     // end of line
	EMPTY                   // marks empty line for formatting
	COMMENT                 // always go to end of the line
	DIREC                   // simple directive, no parameters
	DIREC_PARA              // directive with one or more parameters
	OPC_0                   // SAN opcode without any operands ("nop")
	OPC_1                   // SAN opcode with exactly one operand
	OPC_2                   // SAN opcode with two operands (mvn, mvp)
	BIN_NUM                 // binary number
	HEX_NUM                 // hexadecimal number
	DEC_NUM                 // decimal number
	LABEL                   // absolute label (ends with ":")
	LOCAL_LABEL             // scoped label (starts with "_")
	ANON_LABEL              // anonymous label (starts with "@")
	LEFT_SQUARE             // [
	RIGHT_SQUARE            // ]
	LEFT_PARENS             // (
	RIGHT_PARENS            // )
	LEFT_CURLY              // {
	RIGHT_CURLY             // }
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
)

var Name = map[int](string){
	EOF:          "EOF",
	START:        "START",
	EOL:          "EOL",
	EMPTY:        "EMPTY",
	COMMENT:      "COMMENT",
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
	LEFT_SQUARE:  "LEFT_SQUARE",
	RIGHT_SQUARE: "RIGHT_SQUARE",
	LEFT_PARENS:  "LEFT_PARENS",
	RIGHT_PARENS: "RIGHT_PARENS",
	LEFT_CURLY:   "LEFT_CURLY",
	RIGHT_CURLY:  "RIGHT_CURLY",
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
}

type Token struct {
	Type  int
	Text  string // raw text
	Line  int    // starts with 1
	Index int    // starts with 1
}

/*
// TODO see if we need these after testing, should probably be moved out to the
// specialized tools

// Print just prints the token name without a final line feed. It is mainly
// used for testing
func (t *Token) Print() {
	fmt.Printf("<%s>", Name[t.Type])
}

// PrintLine displays longform information from the token. It is used mainly
// for testing
func (t *Token) PrintLine() {
	ts := fmt.Sprintf("<%s>", Name[t.Type])
	fmt.Printf("%15s (%02d,%02d): \t%s\n", ts, t.Line, t.Index, t.Text)
}
*/
