// Print the list of tokens produced by the lexer for the
// Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 10. May 2018
// This version: 18. May 2018

package lexer

import (
	"fmt"

	"cthulhu/token"
)

// Tokenlister takes the list of tokens produced by the lexer and prints them
// out in a nice way for debugging purposes
func Tokenlister(ts *[]token.Token) {

	for _, t := range *ts {

		switch t.Type {

		case token.COMMENT, token.EMPTY, token.EOL, token.EOF, token.COMMA,
			token.MINUS, token.HASH, token.L_CURLY, token.R_CURLY,
			token.PLUS, token.L_PARENS, token.R_PARENS,
			token.COMMENT_LINE:
			fmt.Printf("<%s>", token.Name[t.Type])

		case token.LABEL, token.LOCAL_LABEL:
			fmt.Printf("<%s [%s]>", token.Name[t.Type], t.Text)

		default:
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
