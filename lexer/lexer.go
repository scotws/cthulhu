// Lexer package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 02. May 2018

package lexer

import (
	"fmt"

	"goasm65816/token"
)

// Lexer takes a list of raw code lines and returns a list of tokens and a flag
// indicating if the conversion was successful or not. Error are handled by the
// main function.
func Lexer(ls []string) ([]token.Token, bool) {

	var tokens []token.Token
	var ok bool = true

	for n, l := range ls {
		fmt.Println(n+1, ":", l)
	}

	t1 := token.Token{1, "test1", 0, 0, "dummy"}
	t2 := token.Token{2, "test2", 0, 0, "dummy"}

	tokens = append(tokens, t1, t2)

	return tokens, ok
}
