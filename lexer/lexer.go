// Lexer package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 02. May 2018

package lexer

import (
	"goasm65816/token"
)

func Lexer(code []string) ([]token.Token, bool) {

	var tokens []token.Token
	var ok bool = true

	t1 := token.Token{1, "test1", 2, 3, "file.txt"}
	t2 := token.Token{1, "test2", 2, 3, "file.txt"}

	tokens = append(tokens, t1, t2)

	return tokens, ok
}
