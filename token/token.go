// Token structure for goasm65816
// Scot W. Stevenson
// First version: 02. May 2018
// This version: 04. May 2018 ("May the Force be with you")
package token

import (
	"fmt"
)

const (
	T_eof        int = iota // end of file
	T_comment               // starts with ;
	T_directive             // starts with .
	T_eol                   // added for formatting
	T_opcode0               // opcode without an operand
	T_opcode1               // opcode without one operand
	T_opcode2               // opcode without two operands (mvn, mvp)
	T_whitespace            // tabs and spaces
	T_binary                // starts with %
	T_hex                   // starts with $
	T_decimal               // starts with number 0-9
	T_label                 // starts with :
	T_localLabel            // starts with _
	T_anonLabel             // starts with @
	T_leftStack             // [
	T_rightStack            // ]
	T_comma                 // ,
	T_minus                 // -
	T_plus                  // +
	T_slash                 // /
	T_star                  // *
	T_string                // not a list of runes
)

var name = map[int](string){
	T_eof:        "EOF",
	T_comment:    "COMMENT",
	T_directive:  "DIRECTIVE",
	T_eol:        "EOL",
	T_opcode0:    "OPCODE0",
	T_opcode1:    "OPCODE1",
	T_opcode2:    "OPCODE2",
	T_whitespace: "WHITESPACE",
	T_binary:     "BINARY",
	T_hex:        "HEX",
	T_decimal:    "DECIMAL",
	T_label:      "LABEL",
	T_localLabel: "LOCAL_LABEL",
	T_anonLabel:  "ANON_LABEL",
	T_leftStack:  "LEFT_STACK",
	T_rightStack: "RIGHT_STACK",
	T_comma:      "COMMA",
	T_minus:      "MINUS",
	T_plus:       "PLUS",
	T_slash:      "SLASH",
	T_star:       "STAR",
	T_string:     "STRING",
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
	fmt.Print("<", name[t.Type], ">")
}

// PrintLine displays longform information from the token. It is used mainly
// for testing
func (t *Token) PrintLine() {
	fmt.Print("<", name[t.Type], "> (", t.Line, ",", t.Index, "): \t", t.Text, "\n")
}
