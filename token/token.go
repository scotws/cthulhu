// Token structure for goasm65816
// Scot W. Stevenson
// First version: 02. May 2018
// This version: 07. May 2018
package token

import (
	"fmt"
)

const (
	T_eof         int = iota // end of file
	T_comment                // starts with ;
	T_directive              // starts with .
	T_eol                    // added for formatting
	T_opcodeWDC              // opcode for WDC
	T_opcode0                // SAN opcode without an operand
	T_opcode1                // SAN opcode without one operand
	T_opcode2                // SAN opcode without two operands (mvn, mvp)
	T_whitespace             // tabs and spaces
	T_binary                 // starts with %
	T_hex                    // starts with $
	T_decimal                // starts with number 0-9
	T_label                  // starts with :
	T_localLabel             // starts with _
	T_anonLabel              // starts with @
	T_leftSquare             // [
	T_rightSquare            // ]
	T_leftParens             // (
	T_rightParens            // )
	T_leftCurly              // { (UNUSED)
	T_rightCurly             // } (UNUSED)
	T_greater                // >
	T_less                   // <
	T_comma                  // ,
	T_minus                  // -
	T_plus                   // +
	T_slash                  // /
	T_star                   // *
	T_hash                   // #
	T_string                 // not a list of runes
	T_symbol                 // any symbol
	T_start                  // start of file
)

var name = map[int](string){
	T_eof:         "EOF",
	T_comment:     "COMMENT",
	T_directive:   "DIRECTIVE",
	T_eol:         "EOL",
	T_opcodeWDC:   "OPC_WDC",
	T_opcode0:     "OPC_SAN_0",
	T_opcode1:     "OPC_SAN_1",
	T_opcode2:     "OPC_SAN_2",
	T_whitespace:  "WHITESPACE",
	T_binary:      "BINARY",
	T_hex:         "HEX",
	T_decimal:     "DECIMAL",
	T_label:       "LABEL",
	T_localLabel:  "LOCAL_LABEL",
	T_anonLabel:   "ANON_LABEL",
	T_leftSquare:  "LEFT_SQUARE",
	T_rightSquare: "RIGHT_SQUARE",
	T_leftParens:  "LEFT_PARENS",
	T_rightParens: "RIGHT_PARENS",
	T_leftCurly:   "LEFT_CURLY",  // UNUSED
	T_rightCurly:  "RIGHT_CURLY", // UNUSED
	T_greater:     "GREATER",
	T_less:        "LESS",
	T_comma:       "COMMA",
	T_minus:       "MINUS",
	T_plus:        "PLUS",
	T_slash:       "SLASH",
	T_star:        "STAR",
	T_hash:        "HASH",
	T_string:      "STRING",
	T_symbol:      "SYMBOL",
	T_start:       "START",
}

type Token struct {
	Type  int
	Text  string // raw text
	Line  int    // starts with 1
	Index int    // starts with 0
}

// Print just prints the token name without a final line feed. It is mainly
// used for testing
func (t *Token) Print() {
	fmt.Printf("<%s>", name[t.Type])
}

// PrintLine displays longform information from the token. It is used mainly
// for testing
func (t *Token) PrintLine() {
	ts := fmt.Sprintf("<%s>", name[t.Type])
	fmt.Printf("%15s (%02d,%02d): \t%s\n", ts, t.Line, t.Index, t.Text)
}
