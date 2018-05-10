// Print the list of tokens produced by the lexer for the
// Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 10. May 2018
// This version: 10. May 2018

package lexer

import (
	"fmt"

	"cthulhu/token"
)

// Tokenlister takes the list of tokens produced by the lexer and prints them
// out

func Tokenlister(ts *[]token.Token) {

	for _, t := range *ts {

		if t.Type == token.COMMENT || t.Type == token.EMPTY ||
			t.Type == token.EOL || t.Type == token.EOF ||
			t.Type == token.COMMA || t.Type == token.MINUS ||
			t.Type == token.HASH || t.Type == token.LEFT_CURLY ||
			t.Type == token.RIGHT_CURLY || t.Type == token.PLUS ||
			t.Type == token.LEFT_PARENS ||
			t.Type == token.RIGHT_PARENS {
			fmt.Printf("<%s>", token.Name[t.Type])
		} else {
			fmt.Printf("<%s [%s]>", token.Name[t.Type], t.Text)
		}

		if t.Type == token.EOL ||
			t.Type == token.EMPTY ||
			t.Type == token.EOF {
			fmt.Printf("\n")
		}

	}
	fmt.Println()
}
