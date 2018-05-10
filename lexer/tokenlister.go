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
		fmt.Printf("<%s>", token.Name[t.Type])

		if t.Type == token.EOL ||
			t.Type == token.EMPTY ||
			t.Type == token.EOF {
			fmt.Printf("\n")
		}
	}
	fmt.Println()
}
