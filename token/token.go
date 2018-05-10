// Token structure for goasm65816
// Scot W. Stevenson
// First version: 02. May 2018
// This version: 09. May 2018

// The important part for the parser is the number of parameters the token has,
// for instance zero for an instruction such as "nop" and potentially lots of
// them for a directive such as ".byte".

package token

import (
	"fmt"
)

const (
	T_eof           int = iota // end of file
	T_eol                      // added for formatting
	T_whitespace               // tabs and spaces TODO see if even used
	T_comment                  // starts with ;
	T_directive                // simple directive, starts with .
	T_directivePara            // directive with one or more paras
	T_opcWDC                   // simple WDC mnemonic without operands
	T_opcWDCNoPara             // WDC mnemonics with definitely no operands
	T_opcSAN0                  // SAN opcode without operands (nop)
	T_opcSAN1                  // SAN opcode with one operand (jsr <TARGET>)
	T_opcSAN2                  // SAN opcode with two operands (mvn, mvp)
	T_binary                   // number, starts with %
	T_hex                      // number, starts with $
	T_decimal                  // number, starts with number 0-9
	T_label                    // starts with :
	T_localLabel               // starts with _
	T_anonLabel                // starts with @
	T_leftSquare               // [
	T_rightSquare              // ]
	T_leftParens               // (
	T_rightParens              // )
	T_leftCurly                // { (UNUSED)
	T_rightCurly               // } (UNUSED)
	T_greater                  // >
	T_less                     // <
	T_comma                    // ,
	T_minus                    // -
	T_plus                     // +
	T_slash                    // /
	T_star                     // *
	T_hash                     // #
	T_string                   // not a list of runes
	T_symbol                   // any symbol
	T_start                    // start of file
)

var Name = map[int](string){
	T_eof:           "EOF",
	T_eol:           "EOL",
	T_whitespace:    "WS",
	T_comment:       "COMMENT",
	T_directive:     "DIR",
	T_directivePara: "DIR_PARA",
	T_opcWDC:        "OPC_WDC",
	T_opcWDCNoPara:  "OPC_WDC_NOPARA",
	T_opcSAN0:       "OPC_SAN_0",
	T_opcSAN1:       "OPC_SAN_1",
	T_opcSAN2:       "OPC_SAN_2",
	T_binary:        "BINARY",
	T_hex:           "HEX",
	T_decimal:       "DECIMAL",
	T_label:         "LABEL",
	T_localLabel:    "LOCAL_LABEL",
	T_anonLabel:     "ANON_LABEL",
	T_leftSquare:    "LEFT_SQUARE",
	T_rightSquare:   "RIGHT_SQUARE",
	T_leftParens:    "LEFT_PARENS",
	T_rightParens:   "RIGHT_PARENS",
	T_leftCurly:     "LEFT_CURLY",  // UNUSED
	T_rightCurly:    "RIGHT_CURLY", // UNUSED
	T_greater:       "GREATER",
	T_less:          "LESS",
	T_comma:         "COMMA",
	T_minus:         "MINUS",
	T_plus:          "PLUS",
	T_slash:         "SLASH",
	T_star:          "STAR",
	T_hash:          "HASH",
	T_string:        "STRING",
	T_symbol:        "SYMBOL",
	T_start:         "START",
}

type Token struct {
	Type  int
	Text  string // raw text
	Line  int    // starts with 1
	Index int    // starts with 0
}

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
